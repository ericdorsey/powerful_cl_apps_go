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

// printList function prints a List, optionally verbose or only showing incomplete items
func printList(l *todo.List, verbose, incomplete bool) {
    formatted := ""
    for k, t := range *l {
        // If incomplete is true, we only want to see incomplete items
        formatted = ""
        if incomplete {
            if t.Done {
                // Skip printing this one, it's completed
                continue
            }
        }
        prefix := "  "
        if t.Done {
            prefix = "X "
        }
        if verbose {
            // Show completed and created time
            if t.Done {
                formatted += fmt.Sprintf("%s%d: %s -- created: %v, completed %v\n", prefix, k+1, t.Task, t.CreatedAt.Format("2006-01-02 15:04"), t.CompletedAt.Format("2006-01-02 15:04"))
            // Show created time
            } else {
                formatted += fmt.Sprintf("%s%d: %s -- created: %v\n", prefix, k+1, t.Task, t.CreatedAt.Format("2006-01-02 15:04"))
            }
        } else {
            formatted += fmt.Sprintf("%s%d: %s\n", prefix, k+1, t.Task)
        }
        fmt.Printf(formatted)
    }
}

func main() {
    // Parse command line flags
    add := flag.Bool("add", false, "Add task to the ToDo list")
    list := flag.Bool("list", false, "List all tasks")
    complete := flag.Int("complete", 0, "Item to be completed")
    del := flag.Int("del", 0, "Item to be deleted")
    verbose := flag.Bool("verbose", false, "Use in conjunction with -list to show verbose output")
    incomplete := flag.Bool("incomplete", false, "Use in conjunction with -list to only show incomplete items")
    // Parse all the flags
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
            // List the items
            printList(l, *verbose, *incomplete)
        case *incomplete:
            fmt.Printf("This flag must be used in conjunction with -list\n")
        case *verbose:
            fmt.Printf("This flag must be used in conjunction with -list\n")
        case *complete > 0:
            // Complete the given item
            if err := l.Complete(*complete); err != nil {
                fmt.Fprintln(os.Stderr, err)
            }
            // Save the new list
            if err := l.Save(todoFileName); err != nil {
                fmt.Fprintln(os.Stderr, err)
                os.Exit(1)
            }
        case *del > 0:
            // Delete the given item
            if err := l.Delete(*del); err != nil {
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
