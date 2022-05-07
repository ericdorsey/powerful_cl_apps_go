package main

import (
    "bufio"
    "fmt"
    "os"
    "io"
    "flag"
    "strings"
    "github.com/ericdorsey/powerful_cl_apps/pcag_todo"
)

//const todoFileName = ".todo.json"
// Default filename
var todoFileName = ".todo.json"

// getTask function decides where to get the description for a new
// task from: arguments or STDIN
func getTask(r io.Reader, args ...string) (string, error) {
    // Add task from arguments
    if len(args) > 0 {
        return strings.Join(args, " "), nil
    }

    // Add task from STDIN
    s := bufio.NewScanner(r)
    s.Scan()
    if err := s.Err(); err != nil {
        return "", err
    }

    if len(s.Text()) == 0 {
        return "", fmt.Errorf("Task cannot be blank")
    }

    return s.Text(), nil
}

func main() {
    // Parse command line flags
    add := flag.Bool("add", false, "Add task to the ToDo list")
    list := flag.Bool("list", false, "List all tasks")
    complete := flag.Int("complete", 0, "Item to be completed") 
    flag.Parse()

    // Check if user defined the TODO_FILENAME env var for custome filename
    if os.Getenv("TODO_FILENAME") != "" {
        todoFileName = os.Getenv("TODO_FILENAME")
    }

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
            fmt.Print(l)
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
        case *add:
            // When any arguments (excluding flags) are provided, they will
            // be used as the new task
            t, err := getTask(os.Stdin, flag.Args()...)
            if err != nil {
                fmt.Fprintln(os.Stderr, err)
                os.Exit(1)
            }
            l.Add(t)

            // Save the new list
            if err := l.Save(todoFileName); err != nil {
                fmt.Fprintln(os.Stderr, err)
                os.Exit(1)
            }


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
