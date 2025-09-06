# Makefile for System Calls Examples
# Makes it easy to run and test all the syscall examples

.PHONY: all clean run-hello run-create run-read run-demo run-comparison test help

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
	@echo "=== All Examples Completed Successfully ==="

# Clean up generated files
clean:
	@echo "Cleaning up generated files..."
	@rm -f test.txt demo.txt hello.txt
	@echo "✓ Cleanup complete"

# Show help
help:
	@echo "System Calls Learning Examples"
	@echo "=============================="
	@echo
	@echo "Available targets:"
	@echo "  run-hello   - Run the hello world example (simplest)"
	@echo "  run-create     - Run the create/write example (creates test.txt)"
	@echo "  run-read       - Run the read example (reads test.txt)"
	@echo "  run-demo       - Run the complete demo (creates demo.txt)"
	@echo "  run-comparison - Run the comparison between Go and syscalls"
	@echo "  test           - Run all examples in sequence"
	@echo "  clean          - Remove all generated files"
	@echo "  help           - Show this help message"
	@echo
	@echo "Quick start:"
	@echo "  make test   - Run everything"
	@echo "  make clean  - Clean up when done"
