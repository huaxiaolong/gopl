package main

import (
	"fmt"
	"time"

	"github.com/huaxiaolong/gopl/ch2/ex2.4/popcount"
)

func main() {
	start := -time.Now().UnixNano()
	fmt.Println(popcount.PopCount(uint64(11111111111111111111)))
	fmt.Println(popcount.PopCount(uint64(10)))
	fmt.Println(popcount.PopCount(uint64(7)))
	costTime := start + time.Now().UnixNano()
	fmt.Printf("cost time:%d\n", costTime)
}
