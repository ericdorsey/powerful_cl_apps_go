package todo

import (
    "encoding/json"
    "errors"
    "fmt"
    "os"
    "time"
)

// Repr a to do item
type item struct {
    Task            string
    Done            bool
    CreatedAt       time.Time
    CompletedAt     time.Time
}

// Slice of todo items. List type is visibile outside the package (captialized)
type List []item

// Add creates a new todo item
func (l *List) Add(task string) {
    t := item {
        Task: task,
        Done: false,
        CreatedAt: time.Now(),
        CompletedAt: time.Time{},
    }
    // Dereference the pointer to the List type to access the underlying slice
    *l = append(*l, t)
}

// Complete marks an item complete
func (l *List) Complete(i int) error {
    ls := *l
    if i <= 0 || i > len(ls) {
        return fmt.Errorf("Item %d does not exist", i)
    }
    // Adjust index for 0 based index
    ls[i-1].Done = true
    ls[i-1].CompletedAt = time.Now()

    return nil
}

// Deletes a todo item from the list
func (l *List) Delete(i int) error {
    ls := *l
    if i <=0 || i > len(ls) {
        return fmt.Errorf("Item %d does not exist", i)
    }

    // Adjusting index for 0 based index
    *l = append(ls[:i-1], ls[i:]...)

    return nil
}

// Save encodes the list as JSON and saves it
func (l *List) Save(filename string) error {
    js, err := json.Marshal(l)
    if err != nil {
        return err
    }

    return os.WriteFile(filename, js, 0644)
}

// Get method opens filename, decodes JSON and parses into a List
func (l *List) Get(filename string) error {
    file, err := os.ReadFile(filename)
    if err != nil {
        // File doesn't exist
        if errors.Is(err, os.ErrNotExist) {
            return nil
        }
        return err
    }
    // Empty file
    if len(file) == 0 {
        fmt.Println("File was empty! in .Get()")
        return nil
    }

    return json.Unmarshal(file, l)
}

// String prints out a formatted list
// Implements the fmt.Stringer interface
func (l *List) String() string {
    formatted := ""

    for k, t := range *l {
        prefix := "  "
        if t.Done {
            prefix = "X "
        }

        // Adjust the item number k to print numbers starting from 1 instead of 0
        formatted += fmt.Sprintf("%s%d: %s\n", prefix, k+1, t.Task)
    }

    return formatted
}

func main() {
}
