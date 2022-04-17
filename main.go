package main

import (
    "bufio"
    "flag"
    "fmt"
    "io"
    "os"
)

func count(r io.Reader, countLines bool) int {
    // scanner is used to read text from a Reader 
    scanner := bufio.NewScanner(r)
  
    if !countLines { 
        // Define scanner split type to words 
        scanner.Split(bufio.ScanWords)
    }

    // a Counter
    wc := 0
    for scanner.Scan() {
        wc++
    }
    return wc
}

func main() {
    // Define a flags -l to count lines instead of words
    lines := flag.Bool("l", false, "Count lines")
    flag.Parse()
    
    // Output count of words or lines
    fmt.Println(count(os.Stdin, *lines))
}
