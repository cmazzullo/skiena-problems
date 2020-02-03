// 4-37. Implement versions of several different sorting algorithms, such
// as selection sort, insertion sort, heapsort, mergesort, and
// quicksort. Conduct experiments to assess the relative performance of
// these algorithms in a simple application that reads a large text file
// and reports exactly one instance of each word that appears within
// it. This application can be efficiently implemented by sorting all the
// words that occur in the text and then passing through the sorted
// sequence to identify one instance of each distinct word. Write a brief
// report with your conclusions.

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
}
