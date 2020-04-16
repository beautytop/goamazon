package main

import (
	"fmt"
	"github.com/hunterhug/goamazon"
	"time"
)

func main() {
	// New Amazon API Client
	client := goamazon.New().SetWaitTime(500 * time.Microsecond)

	// ExistAsin
	asin := "B07DHRTXF6"
	exist, err := client.ExistASIN(asin)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if !exist {
		fmt.Println("not exist asin:", asin)
		return
	} else {
		fmt.Println("exist asin:", asin)
	}

	// GetASIN
	detail, err := client.GetASIN(asin)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%#v\n", detail)

	// ListReview
	// start from 0 page
	page := 0
	for {
		fmt.Println("get review page:", page)
		rr, err := client.ListReview(asin, page)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		if len(rr) == 0 {
			return
		}

		for _, r := range rr {
			fmt.Printf("[%s %s]:%s=%s\n", r.Date, r.UserName, r.Title, r.Content)
			fmt.Println("--------------------")
		}

		page = page + 1
	}
}
