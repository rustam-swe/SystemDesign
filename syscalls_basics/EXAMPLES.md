# System Calls Examples Reference

This file provides a quick reference to all examples in the syscalls learning environment, organized by module and difficulty.

## 📁 File Operations Module

Located in `file_operations/` directory - Start here for beginners!

### 00_hello_world.go
**What it does:** Prints "Hello, World from syscalls!" to the screen
**Difficulty:** ⭐ Beginner
**Syscalls used:** `SYS_WRITE`, `SYS_EXIT`
**Key concepts:** Basic syscall structure, stdout file descriptor
**Run with:** `make file-hello` or `go run file_operations/00_hello_world.go`
**Perfect for:** Understanding the absolute basics of syscalls

### 01_simple_create_write.go
**What it does:** Creates a file called `test.txt` and writes "Hello from syscalls!" to it
**Difficulty:** ⭐⭐ Beginner
**Syscalls used:** `SYS_OPEN`, `SYS_WRITE`, `SYS_CLOSE`, `SYS_EXIT`
**Key concepts:** File creation, file permissions (0644), resource cleanup
**Run with:** `make file-create` or `go run file_operations/01_simple_create_write.go`
**Perfect for:** Learning file creation and writing basics

### 02_simple_read.go
**What it does:** Reads the content of `test.txt` and displays it on screen
**Difficulty:** ⭐⭐ Beginner
**Syscalls used:** `SYS_OPEN`, `SYS_READ`, `SYS_WRITE`, `SYS_CLOSE`, `SYS_EXIT`
**Key concepts:** File reading, buffer management, null termination
**Prerequisites:** Run `01_simple_create_write.go` first to create test.txt
**Run with:** `make file-read` or `go run file_operations/02_simple_read.go`
**Perfect for:** Understanding file reading and buffer handling

### 03_complete_demo.go
**What it does:** Creates `demo.txt`, writes to it, then reads it back with progress messages
**Difficulty:** ⭐⭐⭐ Intermediate
**Syscalls used:** `SYS_OPEN`, `SYS_WRITE`, `SYS_READ`, `SYS_CLOSE`, `SYS_EXIT`
**Key concepts:** Complete file workflow, error handling, user feedback
**Run with:** `make file-demo` or `go run file_operations/03_complete_demo.go`
**Perfect for:** Seeing a complete file processing application

### comparison.go
**What it does:** Shows the same file operations done with regular Go functions vs direct syscalls
**Difficulty:** ⭐⭐⭐ Intermediate
**Key concepts:** Standard library vs syscalls, when to use each approach
**Uses both:** Regular Go functions AND syscalls for educational comparison
**Run with:** `make file-comparison` or `go run file_operations/comparison.go`
**Perfect for:** Understanding why we normally use Go's standard library

## 📡 Network Programming Module

Located in `network/` directory - Advanced concepts, complete file operations first!

### 01_tcp_client.go
**What it does:** Creates a TCP connection and exchanges messages with a server
**Difficulty:** ⭐⭐⭐ Intermediate
**Syscalls used:** `SYS_SOCKET`, `SYS_CONNECT`, `SYS_WRITE`, `SYS_READ`, `SYS_CLOSE`
**Key concepts:** Socket creation, network addressing, connection management
**Prerequisites:** Understanding of file operations module
**Run with:** `make network-tcp` or `cd network && make tcp-client`
**Perfect for:** Learning basic network programming concepts

### 02_http_client.go
**What it does:** Sends a proper HTTP GET request and parses the response
**Difficulty:** ⭐⭐⭐⭐ Advanced
**Syscalls used:** `SYS_SOCKET`, `SYS_CONNECT`, `SYS_WRITE`, `SYS_READ`, `SYS_CLOSE`
**Key concepts:** HTTP protocol, request headers, response parsing
**Prerequisites:** TCP client concepts
**Run with:** `make network-http` or `cd network && make http-client`
**Perfect for:** Understanding HTTP protocol implementation

### 03_http_server.go
**What it does:** Creates a web server that responds to HTTP requests on port 8080
**Difficulty:** ⭐⭐⭐⭐⭐ Expert
**Syscalls used:** `SYS_SOCKET`, `SYS_BIND`, `SYS_LISTEN`, `SYS_ACCEPT`, `SYS_READ`, `SYS_WRITE`, `SYS_CLOSE`
**Key concepts:** Server sockets, binding, listening, accepting connections, HTTP responses
**Prerequisites:** HTTP client understanding
**Run with:** `make network-server` or `cd network && make server`
**Test with:** Open http://127.0.0.1:8080 in your browser
**Perfect for:** Building complete network applications

### 04_slow_loris.go
**What it does:** Educational demonstration of Slow Loris DoS attack
**Difficulty:** ⭐⭐⭐⭐⭐ Expert
**Syscalls used:** `SYS_SOCKET`, `SYS_CONNECT`, `SYS_WRITE`, `SYS_CLOSE`
**Key concepts:** Connection exhaustion, resource management, attack vectors, defenses
**Prerequisites:** Complete understanding of all previous examples + read SAFETY.md
**⚠️ SAFETY:** Educational purposes only, targets localhost only
**Run with:** `make network-attack` or `cd network && make slow-loris`
**Perfect for:** Understanding security vulnerabilities and defenses

## 🗂️ Module Quick Reference

### File Operations Commands
```bash
# Individual examples
make file-hello        # Hello world
make file-create       # Create and write file
make file-read         # Read file
make file-demo         # Complete demo
make file-comparison   # Compare approaches

# All file examples
make file-test         # Run all file operations

# Help and documentation
make file-help         # Detailed file operations help
```

### Network Programming Commands
```bash
# Individual examples  
make network-tcp       # TCP client
make network-http      # HTTP client  
make network-server    # HTTP server
make network-attack    # Educational attack (read safety first!)

# Test sequences
make network-test      # Automated HTTP test

# Help and documentation
make network-help      # Detailed network programming help
cd network && make safety  # Display safety guidelines
```

## 📚 Learning Progression

### Phase 1: Foundations (30 minutes)
- `file_operations/00_hello_world.go` - Understand basic syscall structure
- `file_operations/comparison.go` - See syscalls vs standard library

### Phase 2: File Mastery (1-2 hours)
- `file_operations/01_simple_create_write.go` - File creation
- `file_operations/02_simple_read.go` - File reading  
- `file_operations/03_complete_demo.go` - Complete workflows

### Phase 3: Network Basics (2-3 hours)
- `network/01_tcp_client.go` - Basic networking
- `network/02_http_client.go` - HTTP protocol
- `network/03_http_server.go` - Server programming

### Phase 4: Advanced Security (2-3 hours)
- Read `network/SAFETY.md` thoroughly
- `network/04_slow_loris.go` - Security concepts
- Study defense mechanisms and ethical considerations

## 🎯 Quick Start Paths

### For Absolute Beginners
```bash
make file-hello        # Start here
make file-comparison   # Understand the difference
make file-test         # Practice everything
```

### For Intermediate Developers
```bash
make file-test         # Quick review of basics
make network-help      # Explore advanced concepts
cd network && make test-http  # Try network programming
```

### For Security-Focused Learning
```bash
make file-test         # Master fundamentals first
cd network && make safety     # Read safety guidelines
cd network && make demo-attack # Educational security demo
```

## 🔧 Files Generated by Examples

### File Operations Module
- `test.txt` - Created by `01_simple_create_write.go`, read by `02_simple_read.go`
- `demo.txt` - Created by `03_complete_demo.go`
- `regular_example.txt` and `syscall_example.txt` - Created by `comparison.go`

### Network Module
- No files generated (network communication only)

**Cleanup:** Use `make clean` to remove all generated files

## 📖 Related Documentation

- **README.md** - Project overview and getting started
- **QUICK_START.md** - 5-minute introduction
- **COMPLETE_GUIDE.md** - Comprehensive learning journey
- **file_operations/README_FILE_OPS.md** - File I/O deep dive
- **network/README_NETWORK.md** - Network programming guide
- **network/SAFETY.md** - Security ethics and guidelines

## 🚀 Getting Started

1. **Choose your starting point** based on your experience level
2. **Read the relevant documentation** for context
3. **Run the examples** using the provided Makefile commands
4. **Experiment and modify** the code to deepen understanding
5. **Progress through the phases** at your own pace

## ✅ Completion Checklist

### File Operations Module
- [ ] Successfully run hello world example
- [ ] Create and read files using syscalls
- [ ] Understand file descriptors and permissions
- [ ] Compare syscalls with standard library approaches
- [ ] Handle errors and manage resources properly

### Network Programming Module
- [ ] Establish TCP connections
- [ ] Send and receive HTTP requests/responses
- [ ] Build a working HTTP server
- [ ] Understand security vulnerabilities and defenses
- [ ] Follow ethical guidelines for security testing

### Overall Mastery
- [ ] Comfortable with direct syscall programming
- [ ] Understand when to use syscalls vs standard libraries
- [ ] Can debug low-level system issues
- [ ] Aware of security implications and best practices
- [ ] Ready to apply knowledge to real-world projects

---

**Happy learning!** Use this reference to quickly find the examples that match your current learning goals and skill level.