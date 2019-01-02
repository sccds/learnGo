package parse

import (
	"learnCrawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9A-Za-z]+)" [^>]*>([^<]+)</a>`

func ParseCity(content []byte) engine.ParseResult {
	res := regexp.MustCompile(cityListRe)
	matcher := res.FindAllSubmatch(content, -1)
	result := engine.ParseResult{}

	for _, m := range matcher {
		//fmt.Printf("city: %v, url: %v\n", string(m[2]), string(m[1]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCityUser,
		})
		result.Items = append(result.Items, "city: "+string(m[2]))
	}
	return result
}
