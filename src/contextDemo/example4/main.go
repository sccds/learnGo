package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	d := time.Now().Add(1 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	// 调用别人的方法，设置时间，控制关闭
	defer cancel()
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
