package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main() {
	var words []string
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {  // for each line of input
		// split lines around whitespace:
		for _, v := range strings.Fields(input.Text()) {
			words = append(words, v)
		}
	}
	fmt.Printf("%v", words)
	var s1, s2 = "hello world", "hello " + "world"
	fmt.Printf("%v", s1 == s2)
}
