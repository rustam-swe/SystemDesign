# Makefile for System Calls Learning Examples
# Master makefile for all syscall modules

.PHONY: all clean help file-examples network-examples file-help network-help test-all

# Default target
all: help

# File operations shortcuts
file-hello:
	@echo "=== Running File Operations: Hello World ==="
	@cd file_operations && make run-hello

file-create:
	@echo "=== Running File Operations: Create/Write ==="
	@cd file_operations && make run-create

file-read:
	@echo "=== Running File Operations: Read ==="
	@cd file_operations && make run-read

file-demo:
	@echo "=== Running File Operations: Complete Demo ==="
	@cd file_operations && make run-demo

file-comparison:
	@echo "=== Running File Operations: Comparison ==="
	@cd file_operations && make run-comparison

file-test:
	@echo "=== Running All File Operations Examples ==="
	@cd file_operations && make test

# Network operations shortcuts
network-tcp:
	@echo "=== Running Network: TCP Client ==="
	@cd network && make tcp-client

network-http:
	@echo "=== Running Network: HTTP Client ==="
	@cd network && make http-client

network-server:
	@echo "=== Running Network: HTTP Server ==="
	@cd network && make server

network-attack:
	@echo "=== Running Network: Educational Attack Demo ==="
	@cd network && make slow-loris

network-test:
	@echo "=== Running Network Test Sequence ==="
	@cd network && make test-http

# Module help commands
file-help:
	@echo "📁 File Operations Module"
	@echo "========================"
	@echo
	@cd file_operations && make help

network-help:
	@echo "📡 Network Programming Module"
	@echo "============================="
	@echo
	@cd network && make help

file-examples:
	@echo "Switching to file operations examples..."
	@cd file_operations && make help

network-examples:
	@echo "Switching to network programming examples..."
	@cd network && make help

# Run complete test suite
test-all: file-test
	@echo
	@echo "📁 File Operations: ✅ PASSED"
	@echo "📡 Network examples available - run 'make network-help' to explore"
	@echo
	@echo "🎉 All basic examples completed successfully!"

# Clean up all modules
clean:
	@echo "Cleaning up all modules..."
	@cd file_operations && make clean 2>/dev/null || true
	@cd network && make clean 2>/dev/null || true
	@echo "✓ All modules cleaned"

# Show main help
help:
	@echo "System Calls Learning Environment"
	@echo "================================="
	@echo
	@echo "🏗️ PROJECT STRUCTURE:"
	@echo "  📁 file_operations/  - Basic file I/O with syscalls"
	@echo "  📡 network/          - Network programming and security"
	@echo
	@echo "📚 LEARNING MODULES:"
	@echo
	@echo "📁 FILE OPERATIONS (Start Here):"
	@echo "  file-hello      - Hello world with syscalls (simplest)"
	@echo "  file-create     - Create and write files"
	@echo "  file-read       - Read files and display content"
	@echo "  file-demo       - Complete file workflow"
	@echo "  file-comparison - Syscalls vs standard Go"
	@echo "  file-test       - Run all file examples"
	@echo "  file-help       - Detailed file operations help"
	@echo
	@echo "📡 NETWORK PROGRAMMING (Advanced):"
	@echo "  network-tcp     - Basic TCP client"
	@echo "  network-http    - HTTP client"
	@echo "  network-server  - HTTP server"
	@echo "  network-attack  - Educational DoS demo (read safety guide first!)"
	@echo "  network-test    - Run network test sequence"
	@echo "  network-help    - Detailed network programming help"
	@echo
	@echo "🔧 UTILITIES:"
	@echo "  test-all        - Run all basic examples"
	@echo "  clean           - Clean up all generated files"
	@echo "  help            - Show this help message"
	@echo
	@echo "🎯 RECOMMENDED LEARNING PATH:"
	@echo "  1. Start here:    make file-test"
	@echo "  2. Then try:      make network-help"
	@echo "  3. Full journey:  make test-all"
	@echo
	@echo "📖 DOCUMENTATION:"
	@echo "  • README.md - Project overview and introduction"
	@echo "  • QUICK_START.md - 5-minute getting started guide"
	@echo "  • COMPLETE_GUIDE.md - Comprehensive learning guide"
	@echo "  • file_operations/README_FILE_OPS.md - File I/O deep dive"
	@echo "  • network/README_NETWORK.md - Network programming guide"
	@echo "  • network/SAFETY.md - Important safety and ethics guidelines"
	@echo
	@echo "⚡ QUICK START:"
	@echo "  make file-test  # Learn the basics"
	@echo "  make clean      # Clean up when done"
