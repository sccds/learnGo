package main

import (
	"learnCrawler/engine"
	"learnCrawler/scheduler"
	"learnCrawler/zhenai/parse"
)

func main() {

	/*
		// 单并发版本
		engine.Run(engine.Request{
			Url:        "http://www.zhenai.com/zhenghun",
			ParserFunc: parse.ParseCity,
		})
	*/

	e := engine.ConcurrentEngine{
		//Scheduler:   &scheduler.SimpleScheduler{},
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parse.ParseCity,
	})

}
