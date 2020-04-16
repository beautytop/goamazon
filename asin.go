package goamazon

import (
	"fmt"
	"github.com/hunterhug/marmot/miner"
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
