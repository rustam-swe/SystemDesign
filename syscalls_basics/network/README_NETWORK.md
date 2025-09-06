# Network Programming with System Calls

This directory contains examples of network programming using only system calls in Go. These examples build on the file I/O syscalls you learned earlier and introduce network concepts.

## What You'll Learn

- **TCP Socket Programming**: How to create, connect, and communicate over TCP
- **HTTP Protocol Basics**: Understanding HTTP requests and responses at the syscall level
- **Server Programming**: Building servers that listen for and handle connections
- **Network Security**: Understanding DoS attacks and their countermeasures

## Prerequisites

Before starting these examples, make sure you've completed the basic syscall examples in the parent directory. You should understand:
- Basic syscalls (SYS_WRITE, SYS_READ, SYS_CLOSE)
- File descriptors
- Error handling with syscalls

## Examples Overview

### 1. `01_tcp_client.go` - Basic TCP Client
**What it does:** Creates a TCP connection to a server and exchanges simple messages
**New concepts:**
- `SYS_SOCKET` - Creating network sockets
- `SYS_CONNECT` - Connecting to remote servers
- Socket addresses (sockaddr_in structure)
- Network byte order (big endian)

**Key syscalls:**
- `SYS_SOCKET` - Create a TCP socket
- `SYS_CONNECT` - Connect to server
- `SYS_WRITE` - Send data over network
- `SYS_READ` - Receive data from network
- `SYS_CLOSE` - Close network connection

### 2. `02_http_client.go` - HTTP Client
**What it does:** Sends a proper HTTP GET request and parses the response
**New concepts:**
- HTTP/1.1 protocol format
- HTTP headers (Host, User-Agent, Connection)
- HTTP response parsing
- Request/response message structure

**Builds on:** TCP client concepts plus HTTP protocol understanding

### 3. `03_http_server.go` - HTTP Server
**What it does:** Creates a web server that responds to HTTP requests
**New concepts:**
- `SYS_BIND` - Binding sockets to addresses
- `SYS_LISTEN` - Listening for connections
- `SYS_ACCEPT` - Accepting client connections
- Server socket lifecycle
- HTTP response generation

**Key syscalls:**
- `SYS_SOCKET` - Create server socket
- `SYS_BIND` - Bind to address/port
- `SYS_LISTEN` - Listen for connections
- `SYS_ACCEPT` - Accept client connections
- `SYS_READ`/`SYS_WRITE` - Handle HTTP requests/responses
- `SYS_CLOSE` - Clean up connections

### 4. `04_slow_loris.go` - Educational DoS Attack
**What it does:** Demonstrates how connection exhaustion attacks work
**New concepts:**
- Resource exhaustion attacks
- Connection management
- Network security vulnerabilities
- Attack mitigation strategies

**⚠️ IMPORTANT:** Educational purposes only! Only use against your own servers.

## How to Run the Examples

### Step 1: Start with TCP Basics
```bash
# In one terminal, you'll need a simple server to connect to
# We'll use netcat as a simple test server
nc -l 8080

# In another terminal, run the TCP client
cd SystemDesign/syscalls_basics/network
go run 01_tcp_client.go
```

### Step 2: HTTP Client and Server
```bash
# First, start the HTTP server
go run 03_http_server.go

# In another terminal, test with the HTTP client
go run 02_http_client.go

# Or test with a web browser
# Go to http://127.0.0.1:8080 in your browser
```

### Step 3: Educational Attack (Advanced)
```bash
# Start the HTTP server first
go run 03_http_server.go

# In another terminal, run the slow loris demo
go run 04_slow_loris.go

# Observe how the server becomes less responsive
# Try accessing http://127.0.0.1:8080 in your browser during the attack
```

## Network Concepts Explained

### Sockets
A socket is an endpoint for network communication. Think of it as a "telephone" for programs:
- `socket()` creates the telephone
- `connect()` dials a number
- `read()/write()` talk and listen
- `close()` hangs up

### TCP vs UDP
- **TCP (SOCK_STREAM)**: Reliable, ordered, connection-based (like a phone call)
- **UDP (SOCK_DGRAM)**: Fast, unreliable, connectionless (like sending postcards)

Our examples use TCP because HTTP requires reliable communication.

### Client vs Server
- **Client**: Initiates connections (`connect()`)
- **Server**: Waits for connections (`bind()`, `listen()`, `accept()`)

### Network Byte Order
Network protocols use "big endian" byte order:
- Most significant byte first
- Port 8080 becomes bytes [0x1F, 0x90]
- IP 127.0.0.1 becomes bytes [127, 0, 0, 1]

### Socket Addresses
The `sockaddr_in` structure tells the system where to connect:
```
struct sockaddr_in {
    family   (2 bytes)  - AF_INET for IPv4
    port     (2 bytes)  - Port number in network byte order
    address  (4 bytes)  - IP address in network byte order
    padding  (8 bytes)  - Zero padding
}
```

## HTTP Protocol Basics

### HTTP Request Format
```
METHOD /path HTTP/1.1\r\n
Header-Name: Header-Value\r\n
Another-Header: Another-Value\r\n
\r\n
[Optional request body]
```

### HTTP Response Format
```
HTTP/1.1 STATUS_CODE Reason Phrase\r\n
Header-Name: Header-Value\r\n
Content-Length: 123\r\n
\r\n
[Response body content]
```

### Key Points
- Lines end with `\r\n` (carriage return + line feed)
- Headers and body are separated by an empty line (`\r\n\r\n`)
- `Content-Length` tells how many bytes the body contains
- `Connection: close` means close after this request

## Security Concepts: Slow Loris Attack

### How It Works
1. **Connection Flood**: Open many TCP connections to the target server
2. **Partial Requests**: Send incomplete HTTP requests very slowly
3. **Resource Exhaustion**: Server keeps connections open, waiting for complete requests
4. **Denial of Service**: Server runs out of connection slots for legitimate users

### Why It's Effective
- Uses minimal bandwidth (hard to detect)
- Exploits servers' patience with slow clients
- Each connection consumes server memory/resources
- Can overwhelm servers with just a few dozen connections

### Defenses
- **Connection timeouts**: Close connections that are too slow
- **Rate limiting**: Limit connections per IP address
- **Request timeouts**: Close connections with incomplete requests
- **Load balancers**: Distribute load and filter malicious traffic
- **DDoS protection**: Commercial services that detect and block attacks

## Common Network Programming Patterns

### Error Handling
Always check syscall return values:
```go
result, _, err := syscall.Syscall(...)
if err != 0 {
    // Handle error
    syscall.Syscall(syscall.SYS_CLOSE, socketFD, 0, 0)
    return
}
```

### Resource Cleanup
Always close sockets when done:
```go
defer syscall.Syscall(syscall.SYS_CLOSE, socketFD, 0, 0)
```

### Server Main Loop
Typical server pattern:
1. Create socket
2. Bind to address
3. Listen for connections
4. Loop: Accept → Handle → Close client
5. Clean up server socket

## Troubleshooting

### "Connection refused"
- Server is not running
- Wrong IP address or port
- Firewall blocking the connection

### "Address already in use"
- Another program is using the port
- Previous server instance didn't shut down cleanly
- Try a different port or wait a few seconds

### "Permission denied"
- Trying to bind to a privileged port (< 1024)
- Run as administrator or use port > 1024

## Safety and Ethics

### Important Guidelines
- **Only test against your own servers**
- **Never attack systems you don't own**
- **Use knowledge responsibly**
- **Understand the legal implications**

### Educational Value
These examples teach you:
- How network attacks work (so you can defend against them)
- Why proper server configuration is important
- How to build resilient network applications
- The fundamentals of network security

## Next Steps

After mastering these examples:
- Learn about SSL/TLS encryption
- Explore UDP and other protocols  
- Study advanced HTTP features (POST, cookies, sessions)
- Investigate network monitoring and intrusion detection
- Learn about modern web technologies and their security implications

## Real-World Applications

Understanding these concepts helps with:
- **Web development**: Knowing what happens "under the hood"
- **System administration**: Configuring servers and firewalls
- **Security analysis**: Understanding attack vectors
- **Performance tuning**: Optimizing network applications
- **Debugging**: Troubleshooting network issues

Remember: These examples use only syscalls for educational purposes. In real applications, use Go's `net` and `net/http` packages for production code!