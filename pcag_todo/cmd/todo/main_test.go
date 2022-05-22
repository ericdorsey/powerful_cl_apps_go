package main_test

import (
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
    "runtime"
    "io"
    "testing"
    "time"
)

var (
    binName = "todo"
    fileName = ".todo.json"
)

// From https://stackoverflow.com/questions/44651266/comparing-current-time-in-unit-test/44654689#44654689
var timeNow = time.Now

func init() {
    fmt.Println("Setting mock time for testing")
    timeNow = func() time.Time {
        t, _ := time.Parse("2006-01-02 15:04", "2022-05-21 10:00")
        return t
    }
}

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
    var timeNow = time.Now
    testTimeNow := timeNow()

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

    t.Run("ListTaksVerbose", func(t *testing.T) {
        cmd := exec.Command(cmdPath, "-list", "-verbose")
        out, err := cmd.CombinedOutput()
        if err != nil {
            t.Fatal(err)
        }
        expected := fmt.Sprintf("  1: %s -- created: %v\n  2: %s -- created: %v\n", task, testTimeNow.Format("2006-01-02 15:04"), task2, testTimeNow.Format("2006-01-02 15:04"))
        if expected != string(out) {
            t.Errorf("Expected %q, got %q instead\n", expected, string(out))
        }
    })

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
