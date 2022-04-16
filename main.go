package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
)

func count(r io.Reader) int {
    // scanner is used to read text from a Reader 
    scanner := bufio.NewScanner(r)
    
    // Define scanner split type to words 
    scanner.Split(bufio.ScanWords)

    // a Counter
    wc := 0
    for scanner.Scan() {
        wc++
    }
    return wc
}

func main() {
    fmt.Println(count(os.Stdin))
}
