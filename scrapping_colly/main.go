package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Post struct {
	Title string
	Date  string
	Url   string
	Body  string
}

func main() {
	c := colly.NewCollector()
	c2 := colly.NewCollector()
	posts := make([]Post, 0)
	counter := 0

	// Find and visit all links
	c.OnHTML(".entry-title a", func(e *colly.HTMLElement) {
		post := Post{}
		post.Url = e.Attr("href")
		post.Title = e.Text
		// posts = append(posts, post)
		// fmt.Println(post)
		// e.Request.Visit(e.Attr("href"))
		c2.Visit(e.Attr("href"))
	})

	c2.OnHTML(".single-post", func(e *colly.HTMLElement) {
		post := Post{}
		postDate := e.DOM.Find(".post-date").Eq(0)
		postTitle := e.DOM.Find(".entry-title a").Eq(0).Text()
		// postContent := e.DOM.Find(".entry-content").Eq(0).Not(".nextup")
		postContent := e.DOM.Find(".entry-content").Eq(0).Children().Not(".nextup").Not("script")

		post.Url = e.Attr("href")
		post.Title = postTitle
		post.Date = postDate.Text()
		post.Body = postContent.Text()
		posts = append(posts, post)
		// fmt.Println(post)
		// e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c2.OnScraped(func(r *colly.Response) {
		// fmt.Println("finished:", counter)
		fmt.Printf("%v. %q \n", counter, posts[0].Title)
		fmt.Printf("%#v\n", posts[0].Body)
		fmt.Println("--------------")
		posts = posts[1:]
		counter++
		// fmt.Printf("%#v \n", posts[0].Body)
		// for _, v := range posts {
		// 	fmt.Printf("%#v \n", v)
		// 	fmt.Println("--------------------")
		// }

	})

	c.Visit("https://www.howtogeek.com/author/rothgar/")
}
