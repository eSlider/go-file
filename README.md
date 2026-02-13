# go-file

Go library providing cross-platform file system utilities: path resolution, existence checks, directory creation, and writability testing.

## Installation

```bash
go get github.com/eslider/go-file
```

## Features

- Project root detection (walks up to find `etc/` or `data/` directories)
- File/directory existence checks
- File size queries
- Directory writability testing
- Recursive directory creation

## Quick Start

```go
package main

import (
    "fmt"

    file "github.com/eslider/go-file"
)

func main() {
    // Check if a file exists
    if file.Exists("/etc/hosts") {
        fmt.Println("File exists, size:", file.Size("/etc/hosts"), "bytes")
    }

    // Ensure a directory exists
    if err := file.PreCreateDirectory("/tmp/my-app/data"); err != nil {
        panic(err)
    }

    // Check if a directory is writable
    if file.IsWritable("/tmp/my-app/data") {
        fmt.Println("Directory is writable")
    }

    // Get project root path
    root, err := file.GetModRootPath()
    if err != nil {
        fmt.Println("Not inside a project:", err)
    } else {
        fmt.Println("Project root:", root)
    }
}
```

## Use Cases

### Test Data Loader

```go
func loadTestFixture(name string) []byte {
    root := file.GetRootPath()
    data, _ := os.ReadFile(filepath.Join(root, "testdata", name))
    return data
}
```

### Safe Directory Setup

```go
func initStorage(base string) error {
    dirs := []string{"uploads", "cache", "logs"}
    for _, d := range dirs {
        path := filepath.Join(base, d)
        if err := file.PreCreateDirectory(path); err != nil {
            return fmt.Errorf("failed to create %s: %w", d, err)
        }
        if !file.IsWritable(path) {
            return fmt.Errorf("%s is not writable", d)
        }
    }
    return nil
}
```

## API

| Function | Description |
|---|---|
| `Exists(path)` | Check if file or directory exists |
| `IsExist(path)` | Check if path exists and is a file (not directory) |
| `Size(path)` | Get file size in bytes (0 if not found) |
| `IsWritable(path)` | Test if directory is writable |
| `PreCreateDirectory(path)` | Create directory and parents if needed |
| `GetModRootPath()` | Find project root by walking up to `etc/` or `data/` |
| `GetRootPath()` | Cached project root (panics if not found) |
| `GetCmdRootPath()` | Current dir, stripping `/etc/` suffix |

## License

[MIT](LICENSE)
