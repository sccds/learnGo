package main

import (
	"time"
	"fmt"
	"os"
	"strconv"
	"net/http"
	"io/ioutil"
	"regexp"
)

// 定义 spider 类型
type Spider struct {
	url string
	header map[string]string
}

func (keyword Spider) get_html_header() string  {
	client := &http.Client{}
	println(keyword.url)
	req, err := http.NewRequest("GET", keyword.url, nil)
	if err != nil {
	}
	for key, val := range keyword.header {
		req.Header.Add(key, val)
	}
	resp, err := client.Do(req)
	if err != nil {
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	}
	return string(body)
}


func parse()  {
	header := map[string]string {
		"Host": "movie.douban.com",
		"Connection": "keep-alive",
		"Cache-Control": "max-age=0",
		"Upgrade-Insecure-Requests": "1",
		"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36",
		"Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8",
		"Referer": "https://movie.douban.com/top250",
	}

	// 创建 excel 文件
	f, err := os.Create("/Users/xliu/Documents/51CTO_Golang/learnGo/src/goproject/day8/project_crawler/result.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 写入标题
	f.WriteString("电影名称,评分,评价人数\t\r\n")

	// 循环每页解析并把结果写入excel
	for i:=0; i<4; i++ {
		fmt.Println("正在抓取第" + strconv.Itoa(i) + "页 ... ")
		url := "https://movie.douban.com/top250?start=" + strconv.Itoa(i*25) + "&filter="
		spider := &Spider{url, header}
		html := spider.get_html_header()

		// 获得 评价人数
		pattern2 := `<span>(.*?)评价</span>`
		rp2 := regexp.MustCompile(pattern2)
		find_txt2 := rp2.FindAllStringSubmatch(html, -1)

		// 获得 评分
		pattern3 := `property="v:average">(.*?)</span>`
		rp3 := regexp.MustCompile(pattern3)
		find_txt3 := rp3.FindAllStringSubmatch(html, -1)

		// 获得 电影名称
		pattern4 := `img width="100" alt="(.*?)" src=`
		rp4 := regexp.MustCompile(pattern4)
		find_txt4 := rp4.FindAllStringSubmatch(html, -1)

		// 写入UTF-8 BOM
		f.WriteString("\xEF\xBB\xBF")  // 如果不写入就会导致写入的汉字乱码
		for i:=0; i<len(find_txt2); i++  {
			fmt.Printf("%s %s %s\n", find_txt4[i][1], find_txt3[i][1], find_txt2[i][1], )
			f.WriteString(find_txt4[i][1] + "," + find_txt3[i][1] + "," + find_txt2[i][1] + "\t\r\n")
		}
	}
}




func main()  {
	t1 := time.Now() // get current time
	parse()
	elapsed := time.Since(t1)
	fmt.Println("爬虫结束，耗时：", elapsed)
}