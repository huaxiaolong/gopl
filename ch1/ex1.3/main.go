package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now().Unix()
	str := ""
	for i := 0; i < 100000; i++ {
		str += strconv.Itoa(i)
	}
	end := time.Now().Unix()
	fmt.Printf("string + cost time:%vsecond\n", end-start)

	start = time.Now().Unix()
	var arr [100000]string
	str = strings.Join(arr[0:], "")
	end = time.Now().Unix()
	fmt.Printf("strings.Join cost time:%vsecond\n", end-start)
}
