package main

import (
	"fmt"

	"github.com/huaxiaolong/gopl/ch2/weightconv"
)

func main() {
	p := weightconv.Pound(10)
	fmt.Println(weightconv.PToK(p))
	k := weightconv.Kg(10)
	fmt.Println(weightconv.KToP(k))
}
