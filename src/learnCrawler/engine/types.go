package engine

type ParseResult struct {
	Items    []interface{} // 返回对应的结果值
	Requests []Request     // 解析器返回一个url对应一个解析器的名称
}

type Request struct {
	Url        string                   // 下一次要爬虫的地址
	ParserFunc func([]byte) ParseResult // 地址内容对应的解析器
}

// 定义一个方法，返回无任何值的 ParseResult
func NilParser([]byte) ParseResult {
	return ParseResult{}
}
