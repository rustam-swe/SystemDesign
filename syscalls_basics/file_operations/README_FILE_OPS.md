# File Operations with System Calls

This directory contains examples of file I/O operations using only system calls in Go. These examples demonstrate the fundamental building blocks of all file operations without any standard library wrappers.

## What You'll Learn

- **Basic System Calls**: Understanding the raw interface to the operating system
- **File Descriptors**: How the OS manages open files and resources
- **File Permissions**: Unix file permission system and security
- **Memory Management**: Working with raw pointers and byte arrays
- **Error Handling**: Checking syscall return values and handling failures

## Examples Overview

### 1. `00_hello_world.go` - Hello World with Syscalls
**What it does:** Prints "Hello, World from syscalls!" to the screen using only system calls
**New concepts:**
- `SYS_WRITE` - Writing data to file descriptors
- `SYS_EXIT` - Exiting programs cleanly
- File descriptor 1 (stdout)
- Basic syscall structure

**Key learning:** This is the simplest possible syscall program - perfect for understanding the basic pattern.

### 2. `01_simple_create_write.go` - Creating and Writing Files
**What it does:** Creates a new file and writes text to it using only syscalls
**New concepts:**
- `SYS_OPEN` - Opening/creating files with flags
- File creation flags (`O_CREAT`, `O_WRONLY`, `O_TRUNC`)
- File permissions (0644)
- Resource cleanup with `SYS_CLOSE`

**Key learning:** File creation, writing, and the importance of closing file descriptors.

### 3. `02_simple_read.go` - Reading Files
**What it does:** Opens an existing file, reads its contents, and displays them
**New concepts:**
- `SYS_OPEN` with read-only flag (`O_RDONLY`)
- `SYS_READ` - Reading data into buffers
- Buffer management and null termination
- Reading from files vs writing to stdout

**Key learning:** File reading patterns and buffer handling.

### 4. `03_complete_demo.go` - Complete File Workflow
**What it does:** Demonstrates a complete file workflow: create → write → read → display
**New concepts:**
- Complete file lifecycle management
- Progress reporting and user feedback
- Error handling throughout the process
- Combining multiple operations safely

**Key learning:** How to build robust file processing applications.

### 5. `comparison.go` - Syscalls vs Standard Library
**What it does:** Shows the same operations done with regular Go vs direct syscalls
**New concepts:**
- Standard library convenience functions
- When to use syscalls vs high-level APIs
- Trade-offs between control and convenience
- Understanding what libraries do "under the hood"

**Key learning:** Appreciation for standard libraries and when raw syscalls are appropriate.

## How to Run the Examples

### Prerequisites
- Go 1.11+ installed
- Unix-like system (macOS, Linux, WSL)
- Terminal/command line access

### Running Individual Examples
```bash
# Start with the simplest example
go run 00_hello_world.go

# Create and write a file
go run 01_simple_create_write.go

# Read the file you just created
go run 02_simple_read.go

# See a complete workflow
go run 03_complete_demo.go

# Compare syscalls with standard Go
go run comparison.go
```

### Using the Makefile
```bash
# From the parent directory
make file-hello      # Run hello world
make file-create     # Create and write file
make file-read       # Read file
make file-demo       # Complete demo
make file-comparison # Comparison example
make file-test       # Run all file examples
```

## Key Concepts Explained

### File Descriptors
- **File descriptor**: A number that represents an open file
- **Standard descriptors**:
  - `0` = stdin (keyboard input)
  - `1` = stdout (screen output)
  - `2` = stderr (error output)
- **Custom descriptors**: `3+` for files you open
- **Important**: Always close file descriptors to prevent resource leaks

### File Permissions (Unix)
```
0644 means:
- Owner: read (4) + write (2) = 6
- Group: read (4) = 4
- Others: read (4) = 4

Common patterns:
- 0644: Regular files (rw-r--r--)
- 0755: Executable files (rwxr-xr-x)
- 0600: Private files (rw-------)
```

### File Flags
- `O_RDONLY`: Open for reading only
- `O_WRONLY`: Open for writing only
- `O_RDWR`: Open for reading and writing
- `O_CREAT`: Create file if it doesn't exist
- `O_TRUNC`: Truncate (empty) file if it exists
- `O_APPEND`: Write data at end of file

### Memory Management with Syscalls
```go
// Convert Go string to syscall format
message := "Hello, World!"
messageBytes := []byte(message)
messagePtr := uintptr(unsafe.Pointer(&messageBytes[0]))
```

**Why this is necessary:** Syscalls need raw memory addresses, but Go normally hides memory management for safety.

### Error Handling Pattern
```go
result, _, err := syscall.Syscall(...)
if err != 0 {
    // Handle error - always check!
    errorMsg := "Something went wrong\n"
    // ... report error and clean up
    syscall.Syscall(syscall.SYS_EXIT, 1, 0, 0)
}
```

## Common Syscalls Reference

### File Operations
- `SYS_OPEN` - Open or create files
- `SYS_READ` - Read data from file descriptors
- `SYS_WRITE` - Write data to file descriptors
- `SYS_CLOSE` - Close file descriptors
- `SYS_CREAT` - Create new files (legacy, use SYS_OPEN instead)

### Process Control
- `SYS_EXIT` - Terminate the program with exit code

### Syscall Pattern
```go
result, _, errno := syscall.Syscall(
    syscall.SYS_OPERATION,  // What operation to perform
    arg1,                   // First argument
    arg2,                   // Second argument
    arg3,                   // Third argument
)
```

## Troubleshooting

### Common Errors

**"Permission denied"**
- Check file permissions with `ls -la filename`
- Make sure you have write access to the directory
- Use appropriate permission flags (0644 for regular files)

**"No such file or directory"**
- Make sure the file exists before trying to read it
- Run the create example before the read example
- Check the current directory with `pwd`

**"Too many open files"**
- You're not closing file descriptors
- Always call `SYS_CLOSE` for every `SYS_OPEN`
- Use proper error handling and cleanup

### Build Errors

**"declared and not used"**
- Go is strict about unused variables
- Use `_` to discard unused return values: `_, _, err := syscall.Syscall(...)`

**"undefined: syscall.SYS_*"**
- Some syscalls are platform-specific
- Make sure you're on a Unix-like system (macOS/Linux)

## Files Created by Examples

- `test.txt` - Created by `01_simple_create_write.go`, read by `02_simple_read.go`
- `demo.txt` - Created by `03_complete_demo.go`
- `regular_example.txt` and `syscall_example.txt` - Created by `comparison.go`

All generated files are excluded from git via `.gitignore`.

## Learning Path

### For Complete Beginners
1. **Start simple**: `00_hello_world.go` - understand the basic syscall pattern
2. **File creation**: `01_simple_create_write.go` - learn file creation and writing
3. **File reading**: `02_simple_read.go` - understand reading and buffers
4. **Complete workflow**: `03_complete_demo.go` - see everything together
5. **Compare approaches**: `comparison.go` - understand when to use each approach

### For Intermediate Learners
1. Study the code carefully - understand every line
2. Experiment with different file permissions
3. Try error conditions (reading non-existent files, etc.)
4. Modify the examples to do different operations
5. Add your own error handling and features

## Why Learn This?

### Educational Value
- **Understand fundamentals**: See how file operations really work
- **Debugging skills**: Know what happens when things go wrong
- **Performance awareness**: Understand the cost of file operations
- **System understanding**: Learn how programs interact with the OS

### Professional Skills
- **Systems programming**: Foundation for low-level development
- **Performance optimization**: Understanding syscall overhead
- **Debugging**: Knowing what libraries do internally
- **Architecture**: Designing efficient file handling systems

## Connection to Real World

Every time you:
- Save a file in an editor
- Load a webpage
- Install software
- Run a database query

...you're using these same fundamental syscalls, just wrapped in convenient libraries.

Understanding these basics helps you:
- Write more efficient programs
- Debug mysterious file issues
- Understand performance bottlenecks
- Make informed architectural decisions

## Next Steps

After mastering file operations:
1. **Move to networking**: Try the `network/` examples
2. **Explore advanced topics**: Process management, memory mapping
3. **Build real applications**: Use this knowledge in practical projects
4. **Study source code**: Look at how Go's `os` package implements these operations

Remember: In real applications, use Go's standard library (`os`, `io`, `ioutil`) for production code. These examples are for understanding the fundamentals that make everything else possible!