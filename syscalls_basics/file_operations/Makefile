# Makefile for File Operations Examples
# Easy commands to run and test file syscall examples

.PHONY: all clean help run-hello run-create run-read run-demo run-comparison test

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
	@echo "=== All File Operations Examples Completed Successfully ==="

# Clean up generated files
clean:
	@echo "Cleaning up generated files..."
	@rm -f test.txt demo.txt hello.txt regular_example.txt syscall_example.txt
	@echo "✓ Cleanup complete"

# Show help
help:
	@echo "File Operations Examples with System Calls"
	@echo "=========================================="
	@echo
	@echo "📚 LEARNING PATH:"
	@echo "  1. Read README_FILE_OPS.md for comprehensive guide"
	@echo "  2. Start with basic examples, progress to advanced"
	@echo "  3. Understand each syscall before moving to next example"
	@echo
	@echo "🚀 INDIVIDUAL EXAMPLES:"
	@echo "  run-hello      - Hello world with syscalls (simplest)"
	@echo "  run-create     - Create and write files"
	@echo "  run-read       - Read files and display content"
	@echo "  run-demo       - Complete file workflow demonstration"
	@echo "  run-comparison - Compare syscalls vs standard Go"
	@echo
	@echo "🧪 TESTING:"
	@echo "  test          - Run all examples in sequence"
	@echo
	@echo "📋 UTILITIES:"
	@echo "  clean         - Remove all generated files"
	@echo "  help          - Show this help message"
	@echo
	@echo "🎯 QUICK START:"
	@echo "  make test     - Run everything"
	@echo "  make clean    - Clean up when done"
	@echo
	@echo "📖 For detailed explanations, see README_FILE_OPS.md"
