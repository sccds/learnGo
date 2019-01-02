package parse

import (
	"learnCrawler/engine"
	"learnCrawler/model"
	"regexp"
	"strconv"
)

//var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
//var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span>([\d]+)CM</td>`)

// 城市，年龄，学历，婚况，身高，收入
var infoRe = regexp.MustCompile(`<div class="des f-cl" [^>]*>(.+) \| ([\d]+)岁 \| (.+) \| (.+) \| ([\d]+)cm \| (.+)元</div>`)

func ParserProfile(content []byte) engine.ParseResult {
	profile := model.Profile{}
	/*
		age, err := strconv.Atoi(extractString(content, ageRe))
		if err == nil {
			profile.Age = age
		}
		height, err := strconv.Atoi(extractString(content, heightRe))
		if err == nil {
			profile.Height = height
		}
	*/
	matcher := infoRe.FindAllSubmatch(content, -1)

	for _, m := range matcher {
		profile.City = string(m[1])
		profile.Age, _ = strconv.Atoi(string(m[2]))
		profile.Education = string(m[3])
		profile.Marriage = string(m[4])
		profile.Height, _ = strconv.Atoi(string(m[5]))
		profile.Income = string(m[6])
	}

	result := engine.ParseResult{Items: []interface{}{profile}}
	return result
}

// 正则表达式查找方法
func extractString(content []byte, re *regexp.Regexp) string {
	matcher := re.FindSubmatch(content)
	if len(matcher) >= 2 {
		return string(matcher[1])
	} else {
		return ""
	}
}
