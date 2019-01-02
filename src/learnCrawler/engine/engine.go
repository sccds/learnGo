package engine

import (
	"learnCrawler/fetch"
	"log"
)

func Run(seed ...Request) {
	var requests []Request
	for _, r := range seed {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)

		// 把对应结果Items循环出来
		for _, item := range parseResult.Items {
			log.Printf("get item %v", item)
		}
	}
}

func worker(r Request) (ParseResult, error) {
	log.Printf("fetching %s", r.Url)
	// 调用fetcher去获取网页内容
	body, err := fetch.Fetch(r.Url)
	if err != nil {
		log.Printf("fetcher[%s] err: %v\n", r.Url, err)
		return ParseResult{}, err
	}
	// 把requests加入队列
	return r.ParserFunc(body), err
}
