package main

import (
	"crypto/sha256"
	"fmt"
)

func diffBitCount(s1 [32]byte, s2 [32]byte) int {
	cnt := 0
	for i := 0; i < 32; i++ {
		if s1[i] != s2[i] {
			cnt++
		}
	}
	return cnt
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%d\n", c1, c2, diffBitCount(c1, c2))
}
