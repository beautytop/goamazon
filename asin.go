package goamazon

import (
	"fmt"
	"github.com/hunterhug/marmot/expert"
	"github.com/hunterhug/marmot/miner"
	"github.com/hunterhug/marmot/util/goquery"
	"regexp"
	"strings"
)

func (c *Client) ExistASIN(asin string) (bool, error) {
	ip := c.GetIP()
	data, err := c.download(ip, fmt.Sprintf("https://www.amazon.com/product-reviews/%s?pageNumber=%d", asin, 0))
	if err != nil {
		WorkerPool.Delete(ip)
		miner.Log().Errorf("api ExistASIN err:%s\n", err.Error())
		return c.ExistASIN(asin)
	}

	robot := IsRobot(data)
	if robot == "robot" {
		WorkerPool.Delete(ip)
		return false, RobotError
	}

	over := Is404(data)
	c.PutIP(ip)
	return !over, nil
}

type AsinDetail struct {
	Asin       string
	Title      string
	BigName    string
	IsStock    bool
	IsFba      bool
	IsAwsSold  bool
	SoldBy     string
	SoldById   string
	Img        string
	IsPrime    bool
	Price      float64
	Reviews    int64
	Score      float64
	Describe   string
	BigRank    int64
	RankDetail string
}

func (c *Client) GetASIN(asin string) (*AsinDetail, error) {
	ip := c.GetIP()
	data, err := c.download(ip, fmt.Sprintf("https://www.amazon.com/dp/%s", asin))
	if err != nil {
		WorkerPool.Delete(ip)
		miner.Log().Errorf("api GetASIN err:%s\n", err.Error())
		return c.GetASIN(asin)
	}

	robot := IsRobot(data)
	if robot == "robot" {
		WorkerPool.Delete(ip)
		return nil, RobotError
	}

	over := Is404(data)
	if over {
		return nil, NotFound404Error
	}

	detail := c.parseAsinDetail(data)
	detail.Asin = asin

	return detail, nil
}

func (c *Client) parseAsinDetail(content []byte) (detail *AsinDetail) {
	detail = new(AsinDetail)
	doc, _ := expert.QueryBytes(content)

	titleTrip := "Amazon.com:"

	//detailBulletsWrapper_feature_div
	title := strings.Replace(doc.Find("title").Text(), titleTrip, "", -1)
	bigName := "null"
	temp := strings.Split(title, ":")
	tempL := len(temp)
	if tempL >= 2 {
		bigName = strings.TrimSpace(temp[tempL-1])
		title = strings.Join(temp[0:tempL-1], ":")
	} else {
		temp := strings.Split(title, " at ")
		tempL := len(temp)
		if tempL >= 2 {
			bigName = strings.TrimSpace(temp[tempL-1])
			title = strings.Join(temp[0:tempL-1], " at ")
		}
	}

	detail.Title = strings.TrimSpace(title)
	detail.BigName = bigName

	detail.Img, _ = doc.Find("#imgTagWrapperId img").Attr("data-old-hires")
	if detail.Img == "" {
		imgStr, _ := doc.Find("#imgTagWrapperId img").Attr("data-a-dynamic-image")
		imgArray := strings.Split(imgStr, "\":")
		if len(imgArray) >= 2 {
			imgTemp := strings.Replace(imgArray[0], "{\"", "", -1)
			if strings.HasPrefix(imgTemp, "http") {
				detail.Img = imgTemp
			}
		}
	}

	inStock := doc.Find("#availability span").Text()
	if strings.Contains(inStock, "Currently unavailable.") {

	} else {
		detail.IsStock = true
	}

	merchantInfo := strings.TrimSpace(doc.Find("#merchant-info").Text())
	if strings.Contains(merchantInfo, "Ships from and sold by Amazon.com.") {
		detail.IsFba = true
		detail.IsAwsSold = true
		detail.SoldBy = "Amazon.com"
	} else {
		if strings.Contains(merchantInfo, "Fulfilled by Amazon.") {
			detail.IsFba = true
		}

		detail.SoldById, _ = doc.Find("#merchant-info #seller-popover-information").Attr("data-merchant-id")
		detail.SoldBy = doc.Find("#merchant-info #sellerProfileTriggerId").Text()
	}

	detail.Describe = fmt.Sprintf("<p>%s</p>", strings.TrimSpace(doc.Find("#productDescription p").Text()))

	review := strings.TrimSpace(doc.Find("#prodDetails #acrCustomerReviewText").Text())
	detail.Reviews, _ = SInt64(strings.Replace(review, " ratings", "", -1))

	score := strings.TrimSpace(doc.Find("#prodDetails .a-icon-star").Text())
	detail.Score, _ = SFloat64(strings.Replace(score, " out of 5 stars", "", -1))
	// descriptionAndDetails
	//prodDetails

	rankStr := strings.TrimSpace(doc.Find("#prodDetails").Text())
	r, _ := regexp.Compile(`#([,\d]{1,10})[\s]{0,1}[A-Za-z0-9]{0,6} in ([^#;)(\n]{2,30})[\s\n]{0,1}[(]{0,1}`)
	god := r.FindAllStringSubmatch(rankStr, -1)
	if len(god) > 0 {
		if len(god[0]) >= 2 {
			detail.BigRank, _ = SInt64(strings.Replace(god[0][1], ",", "", -1))
		}
	}

	i := 0
	for _, v := range god {
		if i == 0 {
			detail.RankDetail = strings.Replace(strings.Replace(v[0], " (", "", -1), "\n", "", -1)
			i = i + 1
			continue
		}
		detail.RankDetail = detail.RankDetail + "\n" + strings.Replace(strings.Replace(v[0], " (", "", -1), "\n", "", -1)
	}

	price := doc.Find(".priceBlockBuyingPriceString").Text()
	if price != "" {
		priceArray := strings.Split(price, " - ")
		if len(priceArray) > 0 {
			detail.Price, _ = SFloat64(strings.Replace(priceArray[0], "$", "", -1))
		}
	}

	table := make([]string, 0)
	doc.Find("#prodDetails table tbody tr").Each(func(i int, selection *goquery.Selection) {
		th := strings.TrimSpace(selection.Find("th").Text())
		td := strings.TrimSpace(selection.Find("td").Text())

		if strings.Contains(th, "Customer Reviews") {
			return
		}
		table = append(table, fmt.Sprintf(`<tr><th>%s</th><td>%s</td></tr>`, th, td))
	})
	if len(table) > 0 {
		if detail.Describe == "" {
			detail.Describe = fmt.Sprintf("<table>%s</table>", strings.Join(table, ""))
		} else {
			detail.Describe = fmt.Sprintf("%s<br/><table>%s</table>", detail.Describe, strings.Join(table, ""))
		}
	}
	return
}
