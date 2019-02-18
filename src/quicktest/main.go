package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Printf("%02d-%d\n", time.Now().Month(), time.Now().Year())

}
