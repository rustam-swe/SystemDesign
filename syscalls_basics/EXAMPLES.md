# System Calls Examples Summary

This file provides a quick reference to all the examples in this directory, ordered from simplest to most complex.

## Quick Start Guide

1. **Start here:** `00_hello_world.go` - Just prints "Hello World" using syscalls
2. **Basic file operations:** `01_simple_create_write.go` - Creates and writes to a file
3. **Reading files:** `02_simple_read.go` - Reads from the file you created
4. **Complete workflow:** `03_complete_demo.go` - Does everything in one program
5. **See the difference:** `comparison.go` - Shows regular Go vs syscalls

## All Examples

### 00_hello_world.go
**What it does:** Prints "Hello, World from syscalls!" to the screen
**Syscalls used:** `SYS_WRITE`, `SYS_EXIT`
**Good for:** Understanding the absolute basics
**Run with:** `go run 00_hello_world.go`

### 01_simple_create_write.go
**What it does:** Creates a file called `test.txt` and writes "Hello from syscalls!" to it
**Syscalls used:** `SYS_OPEN`, `SYS_WRITE`, `SYS_CLOSE`, `SYS_EXIT`
**Good for:** Learning file creation and writing
**Run with:** `go run 01_simple_create_write.go`

### 02_simple_read.go
**What it does:** Reads the content of `test.txt` and displays it on screen
**Syscalls used:** `SYS_OPEN`, `SYS_READ`, `SYS_WRITE`, `SYS_CLOSE`, `SYS_EXIT`
**Good for:** Learning file reading
**Prerequisites:** Run `01_simple_create_write.go` first
**Run with:** `go run 02_simple_read.go`

### 03_complete_demo.go
**What it does:** Creates `demo.txt`, writes to it, then reads it back and displays everything with progress messages
**Syscalls used:** `SYS_OPEN`, `SYS_WRITE`, `SYS_READ`, `SYS_CLOSE`, `SYS_EXIT`
**Good for:** Seeing a complete file workflow in one program
**Run with:** `go run 03_complete_demo.go`

### comparison.go
**What it does:** Shows the same file operations done with regular Go functions vs direct syscalls
**Good for:** Understanding why we normally use Go's standard library
**Uses both:** Regular Go functions AND syscalls for comparison
**Run with:** `go run comparison.go`

## Learning Path

### For Complete Beginners
1. Read the main `README.md` file first
2. Run `00_hello_world.go` to see the simplest possible syscall
3. Run `comparison.go` to understand the difference
4. Try `01_simple_create_write.go` and `02_simple_read.go`
5. Finish with `03_complete_demo.go`

### For Intermediate Learners
1. Study the code in each example
2. Try modifying the messages or filenames
3. Add your own error handling
4. Experiment with different file permissions
5. Look up other syscalls you might want to try

## Using the Makefile

Instead of typing `go run filename.go`, you can use these shortcuts:

```bash
make run-hello      # Runs 00_hello_world.go
make run-create     # Runs 01_simple_create_write.go  
make run-read       # Runs 02_simple_read.go
make run-demo       # Runs 03_complete_demo.go
make run-comparison # Runs comparison.go
make test           # Runs all examples in order
make clean          # Removes generated files
make help           # Shows all available commands
```

## Common Questions

**Q: Why is the syscall code so much more complex than regular Go?**
A: Syscalls are the raw interface to the operating system. Go's standard library provides convenient wrappers that handle all the complexity for you.

**Q: Should I use syscalls in real programs?**
A: Generally no. Use Go's standard library (`os`, `io`, etc.) for real programs. Learn syscalls to understand what happens underneath.

**Q: What's with all the `unsafe.Pointer` stuff?**
A: Syscalls need raw memory addresses, but Go normally protects you from working with raw memory. The `unsafe` package lets us work with raw pointers when necessary.

**Q: Why do we need to add `\x00` to filenames?**
A: C-style strings (which syscalls expect) need a null terminator to know where the string ends.

**Q: What other syscalls exist?**
A: Hundreds! Some interesting ones include `fork` (create process), `socket` (network), `mmap` (memory mapping), and `pipe` (interprocess communication).

## Files Created by Examples

- `test.txt` - Created by `01_simple_create_write.go`, read by `02_simple_read.go`
- `demo.txt` - Created by `03_complete_demo.go`
- `regular_example.txt` and `syscall_example.txt` - Temporarily created by `comparison.go`

Use `make clean` to remove all generated files.

## Next Steps

After mastering these examples, you might want to explore:
- Network syscalls (socket programming)
- Process management syscalls
- Memory management syscalls  
- Signal handling
- File system operations beyond basic read/write

Happy learning!