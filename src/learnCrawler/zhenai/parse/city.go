package parse

import (
	"learnCrawler/engine"
	"regexp"
)

const (
	cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`
	nextRe = `href="(http://www.zhenai.com/zhenghun/[^"]+)"`
)

func ParseCityUser(content []byte) engine.ParseResult {
	res := regexp.MustCompile(cityRe)
	matcher := res.FindAllSubmatch(content, -1)
	result := engine.ParseResult{}
	for _, m := range matcher {
		name := string(m[2])
		result.Items = append(result.Items, "user: "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParserProfile,
		})
	}
	// 解析后面的页
	res = regexp.MustCompile(nextRe)
	matcher = res.FindAllSubmatch(content, -1)
	for _, m := range matcher {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCityUser,
		})
	}

	return result
}
