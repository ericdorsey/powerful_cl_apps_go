package main

import (
    "bufio"
    "flag"
    "fmt"
    "io"
    "os"
)

func count(r io.Reader, countLines bool) (int, int) {
    // scanner is used to read text from a Reader 
    scanner := bufio.NewScanner(r)

    if !countLines { 
        // Define scanner split type to words 
        fmt.Println("countlines was not set")
        scanner.Split(bufio.ScanWords)
    }

    // Total Byte length
    bl := 0

    // Word count
    wc := 0
    for scanner.Scan() {
        wc++
        //bl += len(scanner.Bytes())
    }

    //length := len(scanner.Bytes())
    fmt.Printf("Inside function -- wc is: %d, bl is : %d\n", wc, bl) 
    return wc, bl
}

func main() {
    // Define a flags -l to count lines instead of words
    lines := flag.Bool("l", false, "Count lines")
    bytel := flag.Bool("b", false, "Count bytes")
    flag.Parse()
    fmt.Printf("Flags set -- lines: %v, bytel: %v\n", *lines, *bytel)
    
    // Output count of words or lines
    if *bytel {
        wc, bl := count(os.Stdin, *lines)
        fmt.Printf("wc: %d\nbl: %d\n", wc, bl)
    } else {
        wc, _ := count(os.Stdin, *lines)
        fmt.Printf("wc: %d\n", wc)
    }
}
