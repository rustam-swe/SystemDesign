# Makefile for System Calls Examples
# Makes it easy to run and test all the syscall examples

.PHONY: all clean run-hello run-create run-read run-demo run-comparison test help network-help network-examples

# Default target
all: help

# Run the hello world example
run-hello:
	@echo "=== Running Hello World Example ==="
	go run 00_hello_world.go
	@echo "✓ Hello world completed"
	@echo

# Run the simple create/write example
run-create:
	@echo "=== Running Create/Write Example ==="
	go run 01_simple_create_write.go
	@echo "✓ Created test.txt"
	@echo

# Run the simple read example
run-read:
	@echo "=== Running Read Example ==="
	go run 02_simple_read.go
	@echo "✓ Read test.txt"
	@echo

# Run the complete demo
run-demo:
	@echo "=== Running Complete Demo ==="
	go run 03_complete_demo.go
	@echo "✓ Demo completed"
	@echo

# Run the comparison example
run-comparison:
	@echo "=== Running Comparison Example ==="
	go run comparison.go
	@echo "✓ Comparison completed"
	@echo

# Run all examples in sequence
test: run-hello run-create run-read run-demo run-comparison
	@echo "=== All Basic Examples Completed Successfully ==="
	@echo "📡 Ready for network examples? Try: make network-help"

# Clean up generated files
clean:
	@echo "Cleaning up generated files..."
	@rm -f test.txt demo.txt hello.txt
	@cd network && make clean 2>/dev/null || true
	@echo "✓ Cleanup complete"

# Network examples help
network-help:
	@echo "📡 Network Programming Examples Available!"
	@echo "========================================="
	@echo
	@echo "Location: ./network/"
	@echo
	@echo "To get started with network examples:"
	@echo "  cd network"
	@echo "  make help"
	@echo
	@echo "Or run directly:"
	@echo "  make network-examples"
	@echo
	@echo "⚠️  Important: Read network/SAFETY.md before using attack examples!"

# Run network examples (redirect to network directory)
network-examples:
	@echo "Switching to network examples..."
	@cd network && make help

# Show help
help:
	@echo "System Calls Learning Examples"
	@echo "=============================="
	@echo
	@echo "📁 BASIC FILE OPERATIONS:"
	@echo "  run-hello      - Run the hello world example (simplest)"
	@echo "  run-create     - Run the create/write example (creates test.txt)"
	@echo "  run-read       - Run the read example (reads test.txt)"
	@echo "  run-demo       - Run the complete demo (creates demo.txt)"
	@echo "  run-comparison - Run the comparison between Go and syscalls"
	@echo "  test           - Run all basic examples in sequence"
	@echo
	@echo "📡 NETWORK PROGRAMMING:"
	@echo "  network-help     - Show network examples information"
	@echo "  network-examples - Go to network examples"
	@echo
	@echo "🔧 UTILITIES:"
	@echo "  clean - Remove all generated files"
	@echo "  help  - Show this help message"
	@echo
	@echo "🚀 LEARNING PATH:"
	@echo "  1. Start here: make test"
	@echo "  2. Advanced:   make network-help"
	@echo
	@echo "📚 Documentation:"
	@echo "  • README.md - Comprehensive guide to basic syscalls"
	@echo "  • EXAMPLES.md - Quick reference for all examples"
	@echo "  • QUICK_START.md - 5-minute getting started guide"
	@echo "  • network/README_NETWORK.md - Network programming guide"
