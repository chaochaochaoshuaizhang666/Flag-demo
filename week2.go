package main

import (
	"flag"
	"fmt"
)

func main() {
	var t = flag.Int("type", 0, "number type")
	flag.Parse() //解析
	switch *t {
	case 1:
		fmt.Println("花落知多少")
	case 2:
		fmt.Println("西北望，射天狼！")
	default:
		fmt.Println(0)
	}
}
