package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		if (n-i)%3 == 0 && i > 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
	}
	return buf.String()
}

func main() {
	fmt.Println(comma("123"))
	fmt.Println(comma("123123"))
	fmt.Println(comma("12312312"))
}
