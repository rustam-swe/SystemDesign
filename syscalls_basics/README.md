# System Calls Basics in Go

This directory contains simple examples to learn the fundamentals of system calls using Go. All examples use only direct syscalls without any standard library functions for file operations.

## What are System Calls?

System calls are the interface between your program and the operating system kernel. They allow your program to:
- Create, read, and write files
- Allocate memory
- Create processes
- Communicate over networks
- And much more!

## Examples in this Directory

### 0. `00_hello_world.go` - Hello World with Syscalls
The absolute simplest example that shows:
- How to write text to the screen using only syscalls
- The basic structure of a syscall program
- No file operations - just pure syscall basics

**Key syscalls used:**
- `SYS_WRITE` - Write data to stdout (screen)
- `SYS_EXIT` - Exit the program

### 1. `01_simple_create_write.go` - Creating and Writing Files
This is the simplest example that shows how to:
- Create a new file
- Write data to it
- Close the file

**Key syscalls used:**
- `SYS_OPEN` - Create/open a file
- `SYS_WRITE` - Write data to the file
- `SYS_CLOSE` - Close the file descriptor
- `SYS_EXIT` - Exit the program

### 2. `02_simple_read.go` - Reading Files
Shows how to:
- Open an existing file for reading
- Read data from it into a buffer
- Display the contents on screen

**Key syscalls used:**
- `SYS_OPEN` - Open file for reading
- `SYS_READ` - Read data from file
- `SYS_WRITE` - Write to stdout (screen)
- `SYS_CLOSE` - Close file descriptor
- `SYS_EXIT` - Exit the program

### 3. `03_complete_demo.go` - Complete Example
A comprehensive example that combines all operations:
- Creates a file and writes to it
- Reopens the file and reads from it
- Displays everything with helpful messages

## How to Run the Examples

1. **Start with the hello world example:**
   ```bash
   cd SystemDesign/syscalls_basics
   go run 00_hello_world.go
   ```
   This just prints a message using syscalls - perfect for beginners!

2. **Then, run the create/write example:**
   ```bash
   go run 01_simple_create_write.go
   ```
   This creates a file called `test.txt` with some content.

3. **Then, run the read example:**
   ```bash
   go run 02_simple_read.go
   ```
   This reads and displays the content of `test.txt`.

4. **Or run the complete demo:**
   ```bash
   go run 03_complete_demo.go
   ```
   This does everything in one program and creates `demo.txt`.

## Key Concepts Explained

### File Descriptors
- A file descriptor is just a number that represents an open file
- `0` = stdin (keyboard input)
- `1` = stdout (screen output)
- `2` = stderr (error output)
- `3+` = files you open

### File Flags
- `O_RDONLY` - Open for reading only
- `O_WRONLY` - Open for writing only
- `O_CREAT` - Create file if it doesn't exist
- `O_TRUNC` - Truncate (empty) file if it exists

### File Permissions
- `0644` means:
  - Owner can read and write (6 = 4+2)
  - Group can read (4)
  - Others can read (4)

### The `unsafe.Pointer` Mystery
Go syscalls need pointers to memory, but Go normally doesn't let you work with raw memory addresses. The `unsafe` package lets us convert Go data to the format syscalls expect. Don't worry about the details - just know it's necessary for syscalls.

## What Makes This "Raw"?

These examples are special because they:
- Don't use `os.Open()`, `os.Write()`, etc. (those are wrappers around syscalls)
- Don't use `fmt.Println()` or any printing functions
- Don't use `ioutil` or any file utilities
- Call the operating system directly using `syscall.Syscall()`

This is as close as you can get to the operating system without writing assembly language!

## Next Steps

After understanding these basics, you might want to explore:
- Error handling in more detail
- Different file modes and permissions
- Network syscalls (socket, bind, listen)
- Process management syscalls (fork, exec)
- Memory management syscalls (mmap, brk)

## Common Errors

1. **File not found**: Make sure to run the write example before the read example
2. **Permission denied**: Check file permissions and directory access
3. **Compilation errors**: Make sure you're using Go 1.11+ and have proper module support

## Why Learn This?

Understanding syscalls helps you:
- Write more efficient programs
- Debug system-level issues
- Understand how libraries really work "under the hood"
- Build low-level tools and systems software
- Appreciate what higher-level languages do for you!
