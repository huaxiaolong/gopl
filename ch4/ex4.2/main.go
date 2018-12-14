package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var (
	h bool
	s string
)

func init() {
	flag.BoolVar(&h, "h", false, "thsi help")
	flag.StringVar(&s, "s", "", "sha256, sha384, sha512")
	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stderr, `sha
Usage: [-s sha]

Options:
`)
	flag.PrintDefaults()
}

func main() {
	flag.Parse()

	if h {
		flag.Usage()
		return
	}
	if len(os.Args) == 1 {
		return
	}
	x := []byte(os.Args[1])
	if s == "sha384" {
		fmt.Printf("%x\n", sha512.Sum384(x))
	}
	if s == "sha512" {
		fmt.Printf("%x\n", sha512.Sum512(x))
	}
	fmt.Printf("%x\n", sha256.Sum256(x))
}
