package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/olivere/elastic"
)

type Employee struct {
	FirstName string   `json:"firstname"`
	LastName  string   `json:"lastname"`
	Age       int      `json:"age"`
	About     string   `json:"about"`
	Interests []string `json:"interests"`
}

var client *elastic.Client
var host = "http://127.0.0.1:9200"

func init() {
	errorlog := log.New(os.Stdout, "app", log.LstdFlags)
	var err error
	client, err = elastic.NewClient(elastic.SetErrorLog(errorlog), elastic.SetURL(host))
	if err != nil {
		panic(err)
	}
	info, code, err := client.Ping(host).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("es return with code: %d, version: %v \n", code, info)

	esversionCode, err := client.ElasticsearchVersion(host)
	if err != nil {
		panic(err)
	}
	fmt.Printf("es version: %s\n", esversionCode)
}

// 创建索引, 方式1
func create1() {
	// 1 使用结构体方式存入到es
	e1 := Employee{"Jane", "Smith", 32, "I like collect rock music", []string{"music"}}
	put, err := client.Index().Index("info").Type("employee").Id("3").BodyJson(e1).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("index %s to index %s, type %s \n", put.Id, put.Index, put.Type)
}

// 创建索引，方式2
func create2() {
	e1 := `{"firstname":"john", "lastname":"smith", "age":22, "about":"I like play computer", "interests":["computer", "music"]}`
	put, err := client.Index().Index("info").Type("employee").Id("4").BodyJson(e1).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("index %s to index %s, type %s \n", put.Id, put.Index, put.Type)
}

// get 方式查找
func get() {
	get, err := client.Get().Index("info").Type("employee").Id("1").Do(context.Background())
	if err != nil {
		panic(err)
	}
	if get.Found {
		fmt.Printf("get doc %s in version %d from index %s type %s\n", get.Id, get.Version, get.Index, get.Type)
	}
}

func update() {
	res, err := client.Update().Index("info").Type("employee").Id("1").Doc(map[string]interface{}{"age": 88}).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("update age %s\n", res.Result)
}

// 删除
func delete() {
	res, err := client.Delete().Index("info").Type("employee").Id("1").Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("delete result %s", res.Result)
}

func query1() {
	var res *elastic.SearchResult
	var err error
	res, err = client.Search("info").Type("employee").Do(context.Background())
	if err != nil {
		panic(err)
	}
	printEmployee(res, err)

}

// 条件查询
func query2() {
	q := elastic.NewQueryStringQuery("firstname: john")
	res, err := client.Search("info").Type("employee").Query(q).Do(context.Background())
	if err != nil {
		panic(err)
	}
	printEmployee(res, err)
	if res.Hits.TotalHits > 0 {
		fmt.Printf("found total of %d Employees\n", res.Hits.TotalHits)
		for _, hit := range res.Hits.Hits {
			var t Employee
			err := json.Unmarshal(*hit.Source, &t) // 另一种取出数据的方法
			if err != nil {
				fmt.Println("failed")
			}
			fmt.Printf("employee name %s:%s\n", t.FirstName, t.LastName)
		}
	} else {
		fmt.Printf("not found employees")
	}
}

func query3() {
	// 查询年龄大于30的
	boolq := elastic.NewBoolQuery()
	boolq.Must(elastic.NewMatchQuery("lastname", "smith"))
	boolq.Filter(elastic.NewRangeQuery("age").Gt(30))
	res, err := client.Search("info").Type("employee").Query(boolq).Do(context.Background())
	if err != nil {
		panic(err)
	}
	printEmployee(res, err)
}

// 查询包含like的
func query4() {
	matchPhrase := elastic.NewMatchPhraseQuery("about", "rock")
	res, err := client.Search("info").Type("employee").Query(matchPhrase).Do(context.Background())
	if err != nil {
		panic(err)
	}
	printEmployee(res, err)
}

// 分析
func aggs1() {
	aggs := elastic.NewTermsAggregation().Field("interests")
	res, err := client.Search().Index("info").Type("employee").Aggregation("interests", aggs).Do(context.Background())
	printEmployee(res, err)
}

// 打印查询结果
func printEmployee(res *elastic.SearchResult, err error) {
	if err != nil {
		print(err.Error())
		return
	}
	var typ Employee
	for _, item := range res.Each(reflect.TypeOf(typ)) {
		t := item.(Employee)
		fmt.Printf("%#v\n", t)
	}
}

// 分页
func list(size, page int) {
	if size < 0 || page < 1 {
		fmt.Printf("param error")
		return
	}
	res, err := client.Search("info").Type("employee").Size(size).From((page - 1) * size).Do(context.Background())
	if err != nil {
		print(err.Error())
		return
	}
	printEmployee(res, err)
}

func main() {
	//create1()
	//create2()
	//delete()
	//query1()
	//query2()
	//query3()
	//query4()
	aggs1()
	//list(2, 1)
}
