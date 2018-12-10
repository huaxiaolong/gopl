package main

import (
	"fmt"

	"github.com/huaxiaolong/gopl/ch2/tempconv"
)

func main() {
	var c tempconv.Celsius = 10
	fmt.Println(tempconv.CToF(c))
	fmt.Println(tempconv.CToK(c))
}
