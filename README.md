# Marmot Simple Example ToGet Amazon Goods Info

```
              ,
       __,.._; )
  ,--``' / ,";,\  
  |   __; `-'   ;
  |```          ;                          _
  '-""`!------'/                      _,-'`/
   "===`-'"|_|"                 ____,(__,-'
          (fxx`.________,,---``` ;__|
          | ,-"""""\-..._____,"""""-.
          |;;;'''':::````:::; ;'''': :
          (( .---.  ))     ( ( .---.) )
           : \    \ ;  ____ : /    / ;
            \ |````|',-"----`-|    |'
             (`----'          `----'
             /(____\          /____)
          ,-\ /   /          ,\    \
         (_ _/   /          (__\    \
          ,-\   /               ;-._ |
         (___)_/               (____\|  
```

With [https://github.com/hunterhug/marmot](https://github.com/hunterhug/marmot) Easy Get Public Info In The Internet~

## Usage

```go
go get -u -v github.com/hunterhug/goamazon
```

## API

1. ExistASIN: Check ASIN Goods is exist.
2. GetASIN: Get ASIN Detail Info.
3. ListReview: Get ASIN Goods Review List.

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

```

output:

```go
exist asin: B07DHRTXF6
&goamazon.AsinDetail{Asin:"B07DHRTXF6", Title:"Gildan Men's Crew T-Shirt Multipack | Amazon.com", BigName:"null", IsStock:true, IsFba:false, IsAwsSold:false, SoldBy:"", SoldById:"", Img:"https://images-na.ssl-images-amazon.com/images/I/710o0VupScL._AC_UL1500_.jpg", IsPrime:false, Price:13.95, Reviews:0, Score:0, Describe:"<p>Gildan is one of the world's largest vertically integrated manufacturers of apparel and socks. Gildan uses cotton grown in the USA, which represents the best combination of quality and value for Gildan cotton and cotton blended products. Since 2009, Gildan has proudly displayed the cotton USA mark, licensed by cotton council international, on consumer's product packaging and shipping materials. Gildan environmental program accomplishes two core objectives: reduce our environmental impact and preserve the natural Resources being used in our manufacturing process. At all operating levels, Gildan is aware of the fact that we operate as a part of a greater unit: the environment in which we live and work.</p>", BigRank:0, RankDetail:""}
get review page: 0
[Reviewed in the United States on July 30, 2018 kardude]:DO NOT WASTE YOUR MONEY=bad! bad! bad! Very bad quality under-shirts with extremely thin material. All shirts came with defects and loose threads everywhere. Never again any product from Gildan!
--------------------
[Reviewed in the United States on July 16, 2018 Ray]:Sleeves and Collars not stitched=Few items to note here. First the cost of these shirts are super cheap so I expected them to not be the greatest quality or to last a super long time either. Figured I would buy these for an upcoming vacation and if they lasted after that, it would be a bonus. Two of the four shirts were not even stitched at the sleeve and one had some threads that if pulled would pull apart the collar stitches. There is a difference being of cheaper quality and not even assembling your product properly to sell. With only one good shirt in the bunch, the deal for 15.00 for four shirts really is 15 for one wearable shirt. Beware as you will get what you pay for here.
--------------------
[Reviewed in the United States on April 17, 2018 Jon Brooks]:Comparable to Hanes=I'm used to Hanes shirts, and these are very comparable in quality and fit. The large fits my 6'1" frame correctly.My minor complaint is that the arm holes and girth around the belly is ever so slightly larger than the Hanes, which are also too large. My arms are not small, and there is so much material in the arms that they flare out and minimize the appearance of my physique. The length could be just a tad longer too (perhaps half an inch). I realize these are undershirts and wear them as such, but I would prefer that they were more form-fitting. I realize they are necessarily made larger to accommodate heavier guys (I'm 190lbs), but they just aren't quite what I was hoping to get.For the price, these are a fantastic buy.
--------------------
[Reviewed in the United States on August 14, 2018 Sean]:Product arrived with holes in it.=This set of shirts was pretty affordable and had different colors, making it a must buy for me since I wanted to replace a good portion of my undershirts. Arrived today and I put one one. Fit nice, material was nice, but straight out of the package, one shirt has a very noticeable hole in the shoulder seam. Not entirely sure what happened but I hope the seller contacts me and resolves this.
--------------------
[Reviewed in the United States on June 1, 2018 Brrrrrrad]:Cheap, good for wearing under a work shirt, shrinks=You are looking at these because you want some white t-shirts but want the cheapest set you can find. These shirts are "good enough" for an undershirtz provided that you don't care if they shrink after one wash. If you just wear them under your work shirt and take them off when you get home, these are perfect. If you want to wear these as a shirt by itself, I'm not so sure this is the best option. Maybe I just shouldn't have put this in the dryer but then again other shirts I've done this with haven't shrunk so much.
--------------------
[Reviewed in the United States on September 15, 2018 Blind Faith 99]:Not wearable in public=So the price was right. Hard to complain when it's $2 a shirt. As it turns out, that's about what they're worth. First thing I noticed was that they're short and wide. They're not tapered as suggested in the description and pictures. It's a one-size-fits-all design. That can really be seen in the neck opening. For the Large size that I got, it's a whopping 11.5 inches in diameter. That means for normal men, it will sit low, causing a lot of neck and back to be exposed, and also, the neckline will be really loose. For comparison, my Jockey t-shirts are 8.5 to 10 inches, and a couple that I got at Target (Merona) are 9 inches. To make matters worse, after one wash, the material along the neckline curled badly. An oversized neck opening with a curled rim is NOT a good look. The shirts are simply unwearable in public. So all 6 immediately go into the rag drawer and I'll use them for painting  or other messy jobs where it's okay if the shirt gets ruined. No idea why these shirts are getting good reviews.
--------------------
[Reviewed in the United States on July 22, 2018 D.K]:Plenty good enough for the price=Good price for a 6pack of undershirts. Im 6'3, on the lean side, with a basketball in front.A top notch high end Tshirt, I can get by with a true American large.I ordered an XL on these, expecting a good bit of shrink, like I do other bulk pack T's.The XL fits me absolutely perfect out of pack. Not skin tight, but not a moo moo either.Length out of pack is to my fingertips, arms held down,  like a good sport coat length.Hands up, length covers my belt buckle out of pack.Cheap XL = high end L.After one wash/dry, these drew up about 2 inches, shrank about 1/2-1" over all girth.The fit actually got better, neck is great, but the length is kinda pushing it for moving around, and bending over etc.It will come un tucked under my work shirt. But so does every other cheap bulk pack non-tall shirt.I really probably should've went with a tall size, but these will do fine for an undershirt.Nice summer weight.
--------------------
[Reviewed in the United States on March 8, 2018 Michael A.]:Decent shirt, but shoddy QC.=The shirts are comfortable, but the quality control is bad.  One of the shirts had a little white sticker with a red arrow that pointed to a hole in the shirt. Clearly the QC department marked that as a defective shirt, but it was still included in the package even though it was marked as defective.A couple weeks ago I had bought a pack of the same Gildan shirts at Target and one of the shirts had a 1" diameter hole in it, so apparently this is an ongoing problem with quality control.
--------------------
[Reviewed in the United States on May 27, 2018 G.Ruggiero]:Perfect for the business man=Gildan Men's Crew T-Shirts are very comfortable very well made and I would highly recommend them I wear a lot of white shirts and comfort bility and the style and fabric of the shirt is very important to me I  Used to always  Hans but they've got to be very expensive and I don't mind the price but the quality has also been reduced but Gildan Men's Crew T-Shirts  Suppress my expectations and if I could give them a 10 star I would I would highly recommend
--------------------
[Reviewed in the United States on June 12, 2019 MD]:Exactly as described!=I'm going to be COMPLETELY honest, i was extremely hesitant to make this purchase because on the reviews people were receiving red in their pack which wasn't the the color on the advertisement. Despite my better judgement I ordered and figured if it was wrong I could send them back for a refund. To my pleasant surprise, they came exactly as pictured and they're really nice, thick quality at a great price! My son is extremely picky and will only wear shirts that fit loosely but not baggy and are not bright colors (hence the reason I bought these) and he loves them! They are true to size as well. I washed then and he put one on immediately after the dryer was done so I'd say he's a happy customer as am i. I definitely recommend!!!
--------------------
get review page: 1
[Reviewed in the United States on July 30, 2018 kardude]:DO NOT WASTE YOUR MONEY=bad! bad! bad! Very bad quality under-shirts with extremely thin material. All shirts came with defects and loose threads everywhere. Never again any product from Gildan!
--------------------
[Reviewed in the United States on July 16, 2018 Ray]:Sleeves and Collars not stitched=Few items to note here. First the cost of these shirts are super cheap so I expected them to not be the greatest quality or to last a super long time either. Figured I would buy these for an upcoming vacation and if they lasted after that, it would be a bonus. Two of the four shirts were not even stitched at the sleeve and one had some threads that if pulled would pull apart the collar stitches. There is a difference being of cheaper quality and not even assembling your product properly to sell. With only one good shirt in the bunch, the deal for 15.00 for four shirts really is 15 for one wearable shirt. Beware as you will get what you pay for here.
--------------------
[Reviewed in the United States on April 17, 2018 Jon Brooks]:Comparable to Hanes=I'm used to Hanes shirts, and these are very comparable in quality and fit. The large fits my 6'1" frame correctly.My minor complaint is that the arm holes and girth around the belly is ever so slightly larger than the Hanes, which are also too large. My arms are not small, and there is so much material in the arms that they flare out and minimize the appearance of my physique. The length could be just a tad longer too (perhaps half an inch). I realize these are undershirts and wear them as such, but I would prefer that they were more form-fitting. I realize they are necessarily made larger to accommodate heavier guys (I'm 190lbs), but they just aren't quite what I was hoping to get.For the price, these are a fantastic buy.
--------------------
[Reviewed in the United States on August 14, 2018 Sean]:Product arrived with holes in it.=This set of shirts was pretty affordable and had different colors, making it a must buy for me since I wanted to replace a good portion of my undershirts. Arrived today and I put one one. Fit nice, material was nice, but straight out of the package, one shirt has a very noticeable hole in the shoulder seam. Not entirely sure what happened but I hope the seller contacts me and resolves this.
--------------------
[Reviewed in the United States on June 1, 2018 Brrrrrrad]:Cheap, good for wearing under a work shirt, shrinks=You are looking at these because you want some white t-shirts but want the cheapest set you can find. These shirts are "good enough" for an undershirtz provided that you don't care if they shrink after one wash. If you just wear them under your work shirt and take them off when you get home, these are perfect. If you want to wear these as a shirt by itself, I'm not so sure this is the best option. Maybe I just shouldn't have put this in the dryer but then again other shirts I've done this with haven't shrunk so much.
--------------------
[Reviewed in the United States on September 15, 2018 Blind Faith 99]:Not wearable in public=So the price was right. Hard to complain when it's $2 a shirt. As it turns out, that's about what they're worth. First thing I noticed was that they're short and wide. They're not tapered as suggested in the description and pictures. It's a one-size-fits-all design. That can really be seen in the neck opening. For the Large size that I got, it's a whopping 11.5 inches in diameter. That means for normal men, it will sit low, causing a lot of neck and back to be exposed, and also, the neckline will be really loose. For comparison, my Jockey t-shirts are 8.5 to 10 inches, and a couple that I got at Target (Merona) are 9 inches. To make matters worse, after one wash, the material along the neckline curled badly. An oversized neck opening with a curled rim is NOT a good look. The shirts are simply unwearable in public. So all 6 immediately go into the rag drawer and I'll use them for painting  or other messy jobs where it's okay if the shirt gets ruined. No idea why these shirts are getting good reviews.
--------------------

```
