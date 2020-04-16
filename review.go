package goamazon
//
//import (
//	"fmt"
//	"github.com/hunterhug/marmot/miner"
//	"github.com/hunterhug/marmot/util"
//	"sync"
//)
//
//func(c *Client)ReviewList(asin string,maxPage int,thread int){
//	m := new(sync.Map)
//
//	chs := make(chan int64, thread)
//
//	page := 0
//	for {
//		if page >= maxPage {
//			miner.Log().Debugf("爬取结束了，总共进行了%d次爬取。\n", page)
//			break
//		}
//		for i := 0; i <= maxPage; i++ {
//			page = page + 1
//			if page >= maxPage {
//				chs <- -1
//			} else {
//				go review(i, loopIP(), asin, asinNum, m, chs)
//			}
//		}
//
//		over := false
//		for i := 0; i <= thread; i++ {
//			v := <-chs
//			if v == -1 {
//				over = true
//			}
//		}
//
//		if over {
//			fmt.Printf("爬取结束了，总共进行了%d次爬取。\n", page)
//			break
//		}
//	}
//
//	fmt.Println("处理结果中...")
//	gg := make([]ReviewRecord, 0)
//	m.Range(func(key, value interface{}) bool {
//		rs := value.([]ReviewRecord)
//		gg = append(gg, rs...)
//		return true
//	})
//}
//
//type ReviewRecord struct {
//	UserName string
//	Star     string
//	Title    string
//	Date     string
//	Content  string
//	Color    string
//	Help     string
//	Other    string
//}
//
//
//func(c *Client)review(thread int, ip string, asin string, num int, m *sync.Map, chs chan int64) {
//	fmt.Printf("线程-%d正在努力中，页数:%d\n", thread, num)
//	data, err := spider.Download(ip, fmt.Sprintf("https://www.amazon.com/product-reviews/%s?pageNumber=%d", asin, num))
//	if err != nil {
//		fmt.Printf("爬取失败:%s\n", err.Error())
//		review(thread, loopIP(), asin, num, m, chs)
//		return
//	}
//
//	over := spider.Is404(data)
//	if over {
//		spider.Spiders.Delete(ip)
//		chs <- -1
//		return
//	}
//
//	robot := spider.IsRobot(data)
//	if robot == "robot" {
//		spider.Spiders.Delete(ip)
//		fmt.Printf("爬取失败:%s\n", "机器人攻击，换代理！")
//		review(thread, loopIP(), asin, num, m, chs)
//		return
//	}
//
//	//util.SaveToFile(fmt.Sprintf("./review_%d.html", num), data)
//
//	rs := Parse(data)
//	if len(rs) == 0 {
//		spider.Spiders.Delete(ip)
//		chs <- -1
//		return
//	}
//
//	for _, v := range rs {
//		fmt.Printf("Review %s:%s\n", v.UserName, v.Title)
//	}
//	m.Store(num, rs)
//	chs <- 0
//	return
//}