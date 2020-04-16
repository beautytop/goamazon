package goamazon

import (
	"fmt"
	"github.com/hunterhug/marmot/expert"
	"github.com/hunterhug/marmot/miner"
	"github.com/hunterhug/marmot/util/goquery"
	"strings"
)

type ReviewRecord struct {
	UserName string
	Star     string
	Title    string
	Date     string
	Content  string
	Color    string
	Help     string
	Other    string
}

func (c *Client) ListReview(asin string, page int) ([]ReviewRecord, error) {
	ip := c.GetIP()
	data, err := c.download(ip, fmt.Sprintf("https://www.amazon.com/product-reviews/%s?pageNumber=%d", asin, page))
	if err != nil {
		WorkerPool.Delete(ip)
		miner.Log().Errorf("api ReviewList err:%s\n", err.Error())
		return c.ListReview(asin, page)
	}

	robot := IsRobot(data)
	if robot == "robot" {
		WorkerPool.Delete(ip)
		return nil, RobotError
	}

	if Is404(data) {
		return nil, NotFound404Error
	}
	c.PutIP(ip)
	rr := c.parseReview(data)
	return rr, nil
}

func (c *Client) parseReview(data []byte) (rs []ReviewRecord) {
	rs = make([]ReviewRecord, 0)
	d, err := expert.QueryBytes(data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	d.Find(".aok-relative").Each(func(i int, node *goquery.Selection) {
		name := strings.TrimSpace(node.Find(".a-profile-name").Text())
		star := strings.TrimSpace(node.Find(".a-icon-alt").Text())
		title := strings.TrimSpace(node.Find(".review-title").Text())
		if title == "" {
			return
		}
		date := strings.TrimSpace(node.Find(".review-date").Text())
		content := strings.TrimSpace(node.Find(".review-text-content").Text())
		color := strings.TrimSpace(node.Find(".a-size-mini.a-color-state.a-text-bold").Text())
		other := strings.TrimSpace(node.Find(".a-size-mini.a-link-normal.a-color-secondary").Text())
		help := strings.TrimSpace(node.Find(".a-size-base.a-color-tertiary.cr-vote-text").Text())
		rs = append(rs, ReviewRecord{
			UserName: name,
			Star:     star,
			Title:    title,
			Date:     date,
			Content:  content,
			Color:    color,
			Help:     help,
			Other:    other,
		})
	})
	return
}
