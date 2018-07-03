package main

import (
	"io/ioutil"
	"fmt"
)

func main()  {
	const filename = "day1/example9/a.txt"


/*	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", content)
	}*/

	if content, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", content)
	}

}
