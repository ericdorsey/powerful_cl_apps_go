package main

import (
    "fmt"
    "os"
    "flag"
    "github.com/ericdorsey/powerful_cl_apps/pcag_todo"
)

const todoFileName = ".todo.json"

func main() {
    // Parse command line flags
    task := flag.String("task", "", "Task to be included in the Todo list")
    list := flag.Bool("list", false, "List all tasks")
    complete := flag.Int("complete", 0, "Item to be completed") 
    flag.Parse()

    // Define an items List as a pointer to the type todo.List
    l := &todo.List{}

    // Use the Get command to read to do items from file
    if err := l.Get(todoFileName); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }

    // Decide what to do based on number of arguments provided
    switch {
        case *list:
            // List current todo items
            for _, item := range *l {
                if !item.Done { 
                    fmt.Println(item.Task) 
                }
            } 
        case *complete > 0:
            // Complete the given item
            if err := l. Complete(*complete); err != nil {
                fmt.Fprintln(os.Stderr, err)
            }

            // Save the new list
            if err := l.Save(todoFileName); err != nil {
                fmt.Fprintln(os.Stderr, err)
                os.Exit(1)
            }
        case *task != "":
            // Add the task
            l.Add(*task)

            // Save the new list
            if err := l.Save(todoFileName); err != nil {
                fmt.Fprintln(os.Stderr, err)
                os.Exit(1)
            }
        default: 
            // Invalid flag provided
            fmt.Fprintln(os.Stderr, "Invalid option")
            os.Exit(1) 
    }
}
