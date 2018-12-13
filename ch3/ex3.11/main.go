package main

import (
	"bytes"
	"fmt"
	"strings"
)

func comma(s string) string {
	n := len(s)
	m := n
	dotIdx := strings.Index(s, ".")
	if dotIdx != -1 {
		m = len(s[:dotIdx])
	}
	var buf bytes.Buffer
	for i, j := 0, 0; i < n; i++ {
		if s[i] == '-' || s[i] == '+' {
			buf.WriteByte(s[i])
			m--
			continue
		}
		if s[i] == '.' {
			buf.WriteString(s[i:])
			break
		}
		if (m-j)%3 == 0 && j > 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
		j++
	}
	return buf.String()
}

func main() {
	fmt.Println(comma("123"))
	fmt.Println(comma("123123"))
	fmt.Println(comma("12312312"))
	fmt.Println(comma("123.123"))
	fmt.Println(comma("123123.123"))
	fmt.Println(comma("12312312.123"))
	fmt.Println(comma("-123.123"))
	fmt.Println(comma("-123123.123"))
	fmt.Println(comma("-12312312.123"))
}
