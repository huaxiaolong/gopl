package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			input := bufio.NewScanner(f)
			for input.Scan() {
				if counts[input.Text()] == nil {
					counts[input.Text()] = make(map[string]int)
				}
				counts[input.Text()][arg]++
			}
			f.Close()
		}
	}
	for line, temp := range counts {
		for filename, n := range temp {
			if len(temp) == 1 && n == 1 {
				break
			}
			fmt.Printf("%s\t%s\n", line, filename)
		}
	}
}
