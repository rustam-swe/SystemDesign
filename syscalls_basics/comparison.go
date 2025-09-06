/*
 * comparison.go
 *
 * This file shows the difference between using regular Go functions
 * and using direct syscalls for the same operations.
 *
 * This is for educational purposes - normally we use the regular Go way!
 */

package main

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

func regularGoWay() {
	fmt.Println("=== The Regular Go Way ===")

	// Create and write file the normal way
	file, err := os.Create("regular_example.txt")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}

	message := "Hello from regular Go!\n"
	bytesWritten, err := file.WriteString(message)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		file.Close()
		return
	}

	fmt.Printf("Wrote %d bytes using regular Go functions\n", bytesWritten)
	file.Close()

	// Read file the normal way
	content, err := os.ReadFile("regular_example.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	fmt.Printf("Read back: %s", string(content))
	fmt.Println()
}

func syscallWay() {
	fmt.Println("=== The Direct Syscall Way ===")

	// Create and write file using syscalls
	filename := "syscall_example.txt"
	message := "Hello from direct syscalls!\n"

	// Convert to syscall format
	filenameBytes := []byte(filename + "\x00")
	filenamePtr := uintptr(unsafe.Pointer(&filenameBytes[0]))

	// Create/open file
	fd, _, err := syscall.Syscall(
		syscall.SYS_OPEN,
		filenamePtr,
		syscall.O_CREAT|syscall.O_WRONLY|syscall.O_TRUNC,
		0644,
	)

	if err != 0 {
		// Can't use fmt here in pure syscall way, so we'll use syscall to write error
		errorMsg := "Error creating file with syscalls\n"
		errorBytes := []byte(errorMsg)
		errorPtr := uintptr(unsafe.Pointer(&errorBytes[0]))
		syscall.Syscall(syscall.SYS_WRITE, 2, errorPtr, uintptr(len(errorMsg)))
		return
	}

	// Write to file
	messageBytes := []byte(message)
	messagePtr := uintptr(unsafe.Pointer(&messageBytes[0]))

	bytesWritten, _, err := syscall.Syscall(
		syscall.SYS_WRITE,
		fd,
		messagePtr,
		uintptr(len(message)),
	)

	if err != 0 {
		errorMsg := "Error writing to file with syscalls\n"
		errorBytes := []byte(errorMsg)
		errorPtr := uintptr(unsafe.Pointer(&errorBytes[0]))
		syscall.Syscall(syscall.SYS_WRITE, 2, errorPtr, uintptr(len(errorMsg)))
		syscall.Syscall(syscall.SYS_CLOSE, fd, 0, 0)
		return
	}

	// Close file
	syscall.Syscall(syscall.SYS_CLOSE, fd, 0, 0)

	// Print success message (using regular Go for display purposes)
	fmt.Printf("Wrote %d bytes using direct syscalls\n", bytesWritten)

	// Read file back using syscalls
	fd, _, err = syscall.Syscall(
		syscall.SYS_OPEN,
		filenamePtr,
		syscall.O_RDONLY,
		0,
	)

	if err != 0 {
		fmt.Println("Error opening file for reading with syscalls")
		return
	}

	buffer := make([]byte, 100)
	bufferPtr := uintptr(unsafe.Pointer(&buffer[0]))

	bytesRead, _, err := syscall.Syscall(
		syscall.SYS_READ,
		fd,
		bufferPtr,
		uintptr(len(buffer)),
	)

	syscall.Syscall(syscall.SYS_CLOSE, fd, 0, 0)

	if err != 0 {
		fmt.Println("Error reading file with syscalls")
		return
	}

	fmt.Printf("Read back: %s", string(buffer[:bytesRead]))
	fmt.Println()
}

func showComparison() {
	fmt.Println("=== Code Comparison ===")
	fmt.Println()

	fmt.Println("Regular Go (simple and safe):")
	fmt.Println("  file, err := os.Create(\"filename.txt\")")
	fmt.Println("  file.WriteString(\"Hello World\")")
	fmt.Println("  file.Close()")
	fmt.Println()

	fmt.Println("Direct Syscalls (complex but educational):")
	fmt.Println("  filenamePtr := uintptr(unsafe.Pointer(&filename[0]))")
	fmt.Println("  fd, _, err := syscall.Syscall(SYS_OPEN, filenamePtr, flags, perms)")
	fmt.Println("  syscall.Syscall(SYS_WRITE, fd, messagePtr, length)")
	fmt.Println("  syscall.Syscall(SYS_CLOSE, fd, 0, 0)")
	fmt.Println()

	fmt.Println("Key Differences:")
	fmt.Println("• Regular Go: Easy to read, handles errors automatically, memory safe")
	fmt.Println("• Syscalls: More code, manual error checking, but you see exactly what happens")
	fmt.Println("• Regular Go functions are actually using syscalls under the hood!")
	fmt.Println("• Syscalls are what the operating system actually understands")
	fmt.Println()
}

func main() {
	showComparison()
	regularGoWay()
	fmt.Println()
	syscallWay()

	// Clean up
	os.Remove("regular_example.txt")
	os.Remove("syscall_example.txt")

	fmt.Println("=== Conclusion ===")
	fmt.Println("Both methods do the same thing, but:")
	fmt.Println("• Use regular Go for real programs (easier, safer)")
	fmt.Println("• Learn syscalls to understand what's happening underneath")
	fmt.Println("• Syscalls help you appreciate what Go does for you automatically!")
}
