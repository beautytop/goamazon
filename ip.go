package goamazon

import (
	"github.com/hunterhug/marmot/miner"
	"github.com/hunterhug/marmot/util"
	"strings"
	"sync"
	"time"
)

var ipPool = new(sync.Pool)

func (c *Client) GetIP() string {
	if !c.Proxy {
		time.Sleep(time.Duration(500) * time.Millisecond)
		return ""
	}
	ip := ipPool.Get()
	if ip != nil {
		return ip.(string)
	}

	miner.Log().Debugf("no ip, sleep 1s to get proxy ip")
	util.Sleep(1)
	return c.GetIP()
}

func (c *Client) PutIP(ip string) {
	ipPool.Put(ip)
}

func (c *Client) LoopPutMiIP(account string) {
	for {
		if !c.Proxy {
			break
		}

		miner.Log().Debugf("mi proxy get...")
		ips, err := MiProxy(account, 100)
		if err != nil {
			miner.Log().Errorf("mi proxy ip get err:%s\n", err.Error())
			if strings.Contains(err.Error(), "您提取太过频繁") {
				miner.Log().Errorf("60秒后尝试再次获取，如一直失败，请关闭窗口")
				util.Sleep(60)
			} else {
				miner.Log().Errorf("5秒后尝试再次获取，如一直失败，请关闭窗口")
				util.Sleep(5)
			}
			continue
		}
		if len(ips) == 0 {
			continue
		}
		for _, v := range ips {
			miner.Log().Debugf("mi proxy get %s", v)
			ipPool.Put(v)
		}
		miner.Log().Debugf("mi proxy wait 50s...")
		break
	}

	for {
		select {
		case <-time.After(50 * time.Second):
			if !c.Proxy {
				break
			}
			miner.Log().Debugf("mi proxy get...")
			for {
				ips, err := MiProxy(account, 100)
				if err != nil {
					miner.Log().Errorf("mi proxy ip get err:%s\n", err.Error())
					if strings.Contains(err.Error(), "您提取太过频繁") {
						miner.Log().Errorf("60秒后尝试再次获取，如一直失败，请关闭窗口")
						util.Sleep(60)
					} else {
						miner.Log().Errorf("5秒后尝试再次获取，如一直失败，请关闭窗口")
						util.Sleep(5)
					}
					continue
				}
				if len(ips) == 0 {
					continue
				}
				for _, v := range ips {
					miner.Log().Debugf("mi proxy get %s", v)
					ipPool.Put(v)
				}
				miner.Log().Debugf("mi proxy wait 50s...")
				break
			}
		}
	}

}
