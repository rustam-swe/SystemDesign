# Quick Start Guide - System Calls in Go

Welcome! This is your 5-minute introduction to system calls using Go.

## What You'll Learn

In just a few minutes, you'll understand:
- What system calls are and why they matter
- How to create, write, and read files using only syscalls
- The difference between regular Go and direct syscalls

## Prerequisites

- Go installed on your system
- Basic understanding of Go syntax
- Terminal/command line access

## Step 1: Start Simple (2 minutes)

Run the hello world example:

```bash
go run 00_hello_world.go
```

**What just happened?** You printed text to the screen using a direct system call instead of `fmt.Println()`. The program bypassed Go's standard library and talked directly to the operating system.

## Step 2: Create Your First File (2 minutes)

Create and write to a file:

```bash
go run 01_simple_create_write.go
```

Check that it worked:

```bash
cat test.txt
```

**What just happened?** You created a file and wrote to it using only syscalls. No `os.Create()` or `file.WriteString()` - just raw operating system calls.

## Step 3: Read the File Back (1 minute)

Read what you just wrote:

```bash
go run 02_simple_read.go
```

**What just happened?** You opened a file, read its contents, and displayed them using only syscalls.

## Step 4: See the Complete Picture (2 minutes)

Run the complete demo that does everything:

```bash
go run 03_complete_demo.go
```

**What just happened?** You saw a complete workflow: create file → write data → read it back → display results, all using direct syscalls.

## Step 5: Understand the Difference (3 minutes)

See how syscalls compare to regular Go:

```bash
go run comparison.go
```

**What just happened?** You saw the same operations done two ways:
1. The easy way (regular Go functions)
2. The hard way (direct syscalls)

Both do the same thing, but syscalls show you what's really happening underneath.

## Key Takeaways

✅ **System calls are the real interface to the OS** - everything else is just convenience wrappers

✅ **Regular Go is easier** - use `os.Create()`, `fmt.Println()`, etc. for real programs

✅ **Learning syscalls helps you understand** what happens "under the hood"

✅ **Go's standard library uses syscalls internally** - you just don't see them

## What's Next?

- Read `README.md` for detailed explanations
- Check `EXAMPLES.md` for a complete guide to all examples
- Try modifying the examples - change messages, filenames, permissions
- Use `make test` to run everything at once
- Use `make clean` to clean up generated files

## The Bottom Line

You now understand the foundation of how programs interact with operating systems. Every time you:
- Save a file
- Print to the screen  
- Connect to the internet
- Start a program

...you're using system calls (just hidden behind convenient libraries).

Congratulations! You've learned something that many programmers never see directly. 🎉