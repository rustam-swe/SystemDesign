# Makefile for Network Programming Examples
# Easy commands to run and test network syscall examples

.PHONY: all clean help test-tcp test-http test-server demo-attack safety server client http-client tcp-client slow-loris

# Default target
all: help

# Individual example runners
tcp-client:
	@echo "=== Running TCP Client ==="
	@echo "Note: You need a server listening on localhost:8080"
	@echo "Try running 'make server' in another terminal first"
	@echo
	go run 01_tcp_client.go

http-client:
	@echo "=== Running HTTP Client ==="
	@echo "Note: You need an HTTP server running on localhost:8080"
	@echo "Try running 'make server' in another terminal first"
	@echo
	go run 02_http_client.go

server:
	@echo "=== Starting HTTP Server ==="
	@echo "Server will start on http://127.0.0.1:8080"
	@echo "Press Ctrl+C to stop the server"
	@echo "You can test it by:"
	@echo "  1. Opening http://127.0.0.1:8080 in your browser"
	@echo "  2. Running 'make http-client' in another terminal"
	@echo "  3. Running 'make slow-loris' for educational DoS demo"
	@echo
	go run 03_http_server.go

slow-loris:
	@echo "=== Educational Slow Loris Demonstration ==="
	@echo "⚠️  SAFETY WARNING: This is for educational purposes only!"
	@echo "⚠️  Only targets localhost:8080 (your own server)"
	@echo "⚠️  Make sure you've read SAFETY.md before proceeding"
	@echo
	@echo "Make sure the HTTP server is running first (make server)"
	@echo "Press Enter to continue or Ctrl+C to cancel..."
	@read dummy
	go run 04_slow_loris.go

# Test sequences
test-tcp:
	@echo "=== TCP Client/Server Test ==="
	@echo "This will test basic TCP communication"
	@echo "You'll need to run a simple server manually:"
	@echo "  1. Open another terminal"
	@echo "  2. Run: nc -l 8080"
	@echo "  3. Then run the client"
	@echo
	@echo "Ready? Press Enter to run TCP client..."
	@read dummy
	go run 01_tcp_client.go

test-http:
	@echo "=== HTTP Client/Server Test ==="
	@echo "This will start the server and then test with the client"
	@echo "Note: You'll need to stop the server manually (Ctrl+C)"
	@echo
	@echo "Starting server in background..."
	@go run 03_http_server.go &
	@sleep 2
	@echo "Testing with HTTP client..."
	@go run 02_http_client.go
	@echo
	@echo "Test complete! Don't forget to stop the server (Ctrl+C in server terminal)"

demo-attack:
	@echo "=== Complete Attack Demonstration ==="
	@echo "This demonstrates the full Slow Loris attack scenario"
	@echo
	@echo "⚠️  EDUCATIONAL PURPOSE ONLY ⚠️"
	@echo "⚠️  READ SAFETY.md FIRST ⚠️"
	@echo
	@echo "This will:"
	@echo "  1. Start the HTTP server"
	@echo "  2. Show normal operation"
	@echo "  3. Launch the Slow Loris attack"
	@echo "  4. Demonstrate the impact"
	@echo
	@echo "Press Enter to continue or Ctrl+C to cancel..."
	@read dummy
	@echo
	@echo "Step 1: Starting HTTP server in background..."
	@go run 03_http_server.go &
	@sleep 2
	@echo
	@echo "Step 2: Testing normal operation..."
	@go run 02_http_client.go
	@echo
	@echo "Step 3: Launching educational DoS attack..."
	@go run 04_slow_loris.go
	@echo
	@echo "Demo complete! Server may still be running - use Ctrl+C to stop it"

safety:
	@echo "=== SAFETY AND ETHICS INFORMATION ==="
	@cat SAFETY.md

# Clean up any temporary files
clean:
	@echo "Cleaning up network examples..."
	@echo "✓ No temporary files to clean"

# Help message
help:
	@echo "Network Programming Examples with System Calls"
	@echo "=============================================="
	@echo
	@echo "📚 LEARNING PATH:"
	@echo "  1. Read README_NETWORK.md for comprehensive guide"
	@echo "  2. Read SAFETY.md for important safety information"
	@echo "  3. Start with basic examples, progress to advanced"
	@echo
	@echo "🚀 INDIVIDUAL EXAMPLES:"
	@echo "  tcp-client    - Basic TCP client (needs server on port 8080)"
	@echo "  http-client   - HTTP client (needs HTTP server running)"
	@echo "  server        - Start HTTP server on localhost:8080"
	@echo "  slow-loris    - Educational DoS attack (LOCALHOST ONLY)"
	@echo
	@echo "🧪 TESTING SEQUENCES:"
	@echo "  test-tcp      - Guided TCP client test"
	@echo "  test-http     - HTTP client/server test"
	@echo "  demo-attack   - Complete attack demonstration"
	@echo
	@echo "📋 UTILITIES:"
	@echo "  safety        - Display safety and ethics guide"
	@echo "  clean         - Clean up temporary files"
	@echo "  help          - Show this help message"
	@echo
	@echo "⚠️  SAFETY REMINDERS:"
	@echo "  • All examples target localhost only"
	@echo "  • Educational purposes only"
	@echo "  • Read SAFETY.md before using attack examples"
	@echo "  • Never target systems you don't own"
	@echo
	@echo "🎯 QUICK START:"
	@echo "  make server     # Start server"
	@echo "  make http-client # Test in another terminal"
	@echo
	@echo "📖 For detailed explanations, see README_NETWORK.md"
