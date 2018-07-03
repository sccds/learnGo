package example1

import (
	"fmt"
	"time"
)

func newTask()  {
	for {
		fmt.Println("this is a newTask")
		time.Sleep(time.Second)
	}
}


func main()  {
	go newTask()  // 新建立一个协程
	for  {
		fmt.Println("this is a main goroutine")
		time.Sleep(time.Second)
	}

}
