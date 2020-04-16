# Amazon non official API lib By Golang


## Usage

```go
go get -u -v github.com/hunterhug/goamazon
```

## API

1. ExistASIN: Check ASIN Goods is exist.
2. ListReview: Get ASIN Goods Review List.

## Example

```go
package main

import (
	"fmt"
	"github.com/hunterhug/goamazon"
	"time"
)

func main() {
	// New Amazon API Client
	client := goamazon.New().SetWaitTime(500 * time.Microsecond)

	//client.SetDebug()

	// ExistAsin
	asin := "B07PBJB3R4"
	exist, err := client.ExistASIN(asin)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if !exist {
		fmt.Println("not exist asin:", asin)
		return
	}

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
```

output:

```go
exist asin: B07PBJB3R4
get review page: 0
[Reviewed in the United States on October 20, 2018 J]:True to life=I appreciate this film for depicting the realities of how psychiatric medications and other treatments are forced onto people without their full consent. People in emotional distress still have the ability and the human right to make decisions about their care, and I'm glad this film is helping show that point of view.  This movie is not perfect, it has some generic tropes (like "hero who works so much that and it damages a romantic relationship" or "smarmy lawyer on the opposing side speaks condescendingly to hero"), but I would still recommend this film. I hope it reaches a wider audience because the situations depicted in this movie are still happening in psych wards.
--------------------
[Reviewed in the United States on January 13, 2019 Lisa C. Nelson]:Amazing and touching without being preachy=I loved the moving reality of this.  It showed the difficulty of mental illness.  And it wasn't anti-medication.  So ignore the Scientology reviews.  This movie focused on the issue of consent.  You can have a mental illness yet still have the mental capacity to make your own health decisions.  People do it all the time.  I happened upon this movie, but found myself riveted and moved by it, both by the performances, the story, and the legal precedent it set.  It also opened my eyes a bit wider to the plight of the mentally ill who seek evidence based treatment then get stuck in a detrimental mental health care loop.
--------------------
[Reviewed in the United States on December 15, 2018 Erin]:Tragic story with an important message.=The psychiatric industry is not a source of "help". Unfortunately, many looking for help meet their demise by seeking psychiatrists/institutes, and doctors for psychological/emotional help. Medications can never treat someone's emotions, they are toxic, and have serious life threatening side-effects.
--------------------
```