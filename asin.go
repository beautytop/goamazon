package goamazon

import (
	"fmt"
	"github.com/hunterhug/marmot/miner"
)

func (c *Client) ExistASIN(asin string) bool {
	ip := c.GetIP()
	data, err := Download(ip, fmt.Sprintf("https://www.amazon.com/product-reviews/%s?pageNumber=%d", asin, 0))
	if err != nil {
		miner.Log().Errorf("api ExistASIN err:%s\n", err.Error())
		return c.ExistASIN(asin)
	}

	over := Is404(data)
	if over {
		WorkerPool.Delete(ip)
		return false
	}

	robot := IsRobot(data)
	if robot == "robot" {
		WorkerPool.Delete(ip)
		miner.Log().Errorf("api ExistASIN robot")
		return c.ExistASIN(asin)

	}
	return true
}
