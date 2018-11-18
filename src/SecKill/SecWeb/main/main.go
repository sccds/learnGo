package main

import (
	_ "SecKill/SecWeb/router"
	"fmt"

	"github.com/astaxie/beego"
)

func main() {
	err := initAll()
	if err != nil {
		panic(fmt.Sprintf("init databases failed, err: %v", err))
	}
	beego.Run()
}
