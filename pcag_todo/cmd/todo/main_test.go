package main_test

import (
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
    "runtime"
    "io"
    "testing"
)

var (
    binName = "todo"
    fileName = ".todo.json"
)

func TestMain(m *testing.M) {
    fmt.Println("Building tool...")

    if runtime.GOOS == "windows" {
        binName += ".exe"
    }

    build := exec.Command("go", "build", "-o", binName)

    if err := build.Run(); err != nil {
        fmt.Fprintf(os.Stderr, "Cannot build tool %s: %s", binName, err)
        os.Exit(1)
    }

    fmt.Println("Running tests....")
    result := m.Run()

    fmt.Println("Cleaning up....")
    os.Remove(binName)
    os.Remove(fileName)

    os.Exit(result)
}

func TestTodoCLI(t *testing.T) {
    task := "test task number 1"
    dir, err := os.Getwd()
    if err != nil {
        t.Fatal(err)
    }
    cmdPath := filepath.Join(dir, binName)

    t.Run("AddNewTaskFromArguments", func(t *testing.T) {
        cmd := exec.Command(cmdPath, "-add", task)

        if err := cmd.Run(); err != nil {
            t.Fatal(err)
        }
    })
    task2 := "test task number 2"
    t.Run("AddNewTaskFromSTDIN", func(t *testing.T) {
        cmd := exec.Command(cmdPath, "-add")
        cmdStdIn, err := cmd.StdinPipe()
        if err != nil {
            t.Fatal(err)
        }
        io.WriteString(cmdStdIn, task2)
        cmdStdIn.Close()
        if err := cmd.Run(); err != nil {
            t.Fatal(err)
        }
    })

    t.Run("ListTasks", func(t *testing.T) {
        cmd := exec.Command(cmdPath, "-list")
        out, err := cmd.CombinedOutput()
        if err != nil {
            t.Fatal(err)
        }
        expected := fmt.Sprintf("  1: %s\n  2: %s\n", task, task2)
        if expected != string(out) {
            t.Errorf("Expected %q, got %q instead\n", expected, string(out))
        }
    })

    //t.Run("ListTaksVerbose", 

    t.Run("CompleteTask", func(t *testing.T) {
        cmd := exec.Command(cmdPath, "-complete", "1")
        if err := cmd.Run(); err != nil {
            t.Fatal(err)
        }
    })

    t.Run("DeleteTask", func(t *testing.T) {
        cmd := exec.Command(cmdPath, "-del", "1")
        if err := cmd.Run(); err != nil {
            t.Fatal(err)
        }
    })
}
