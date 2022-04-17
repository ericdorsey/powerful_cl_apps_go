package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func count(r io.Reader, countLines bool, countBytes bool) (int, int) {
	// scanner is used to read text from a Reader
	scanner := bufio.NewScanner(r)

	if !countLines {
		// Define scanner split type to words
		scanner.Split(bufio.ScanWords)
	}

	// Word count
	wc := 0

	// Total Byte length
	bl := 0

	if !countBytes {
		for scanner.Scan() {
			wc++
		}
	} else {
		for scanner.Scan() {
			wc++
			bl += len(scanner.Bytes())
		}
	}

	//fmt.Printf("Inside function --> wc is: %d, bl is : %d\n", wc, bl)
	return wc, bl
}

func main() {
	// Define a flags -l to count lines instead of words
	lines := flag.Bool("l", false, "Count lines")
	// Define a flags -b to count bytes
	bytel := flag.Bool("b", false, "Count bytes")
	flag.Parse()
	//fmt.Printf("Flags set --> lines: %v, bytel: %v\n", *lines, *bytel)

	if *bytel {
		wc, bl := count(os.Stdin, *lines, *bytel)
		if *lines {
			fmt.Printf("%d lines, %d bytes\n", wc, bl)
		} else {
			if wc == 1 {
				fmt.Printf("%d word, %d bytes\n", wc, bl)
			} else {
				fmt.Printf("%d words, %d bytes\n", wc, bl)
			}
		}
	} else {
		wc, _ := count(os.Stdin, *lines, *bytel)
		if *lines {
			fmt.Printf("%d lines\n", wc)
		} else {
			if wc == 1 {
				fmt.Printf("%d word\n", wc)
			} else {
				fmt.Printf("%d words\n", wc)
			}
		}
	}
}
