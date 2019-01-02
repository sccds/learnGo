package engine

import "log"

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	// 传入request方法
	Submit(Request)
	WorkerChan() chan Request
	Run()
	ReadyNotifier
}

type ReadyNotifier interface {
	// 传入协程对应的接收request的chan
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seed ...Request) {

	out := make(chan ParseResult)
	e.Scheduler.Run()

	// 启动 worker
	for i := 0; i < e.WorkerCount; i++ {
		createWork(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	// 把 request 放入到 channel 里面
	for _, req := range seed {
		if isDuplicate(req.Url) {
			continue
		}
		e.Scheduler.Submit(req)
	}
	itemCount := 0

	// 循环遍历结果
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Get item %d: %v\n", itemCount, item)
			itemCount++
		}
		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true
	return false
}

func createWork(in chan Request, out chan ParseResult, s ReadyNotifier) {
	go func() {
		for {
			s.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
