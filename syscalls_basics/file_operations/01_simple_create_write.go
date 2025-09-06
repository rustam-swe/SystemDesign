package main

import (
	"syscall"
	"unsafe"
)

func main() {
	// What we want to write to the file
	message := "Hello from syscalls!\n"
	filename := "test.txt"

	// Step 1: Create/Open a file for writing
	// We need to convert our filename string to a format syscalls understand
	filenameBytes := []byte(filename + "\x00") // Add null terminator
	filenamePtr := uintptr(unsafe.Pointer(&filenameBytes[0]))

	// Make the syscall to open/create the file
	// Parameters: SYS_OPEN, filename, flags, permissions
	fd, _, err := syscall.Syscall(
		syscall.SYS_OPEN,                                    // The syscall we want
		filenamePtr,                                         // Pointer to filename
		syscall.O_CREAT|syscall.O_WRONLY|syscall.O_TRUNC,  // Create if not exists, write-only, truncate
		0644,                                                // File permissions (owner: read/write, others: read)
	)

	// Check if opening the file failed
	if err != 0 {
		// If error, just exit the program
		syscall.Syscall(syscall.SYS_EXIT, 1, 0, 0)
	}

	// Step 2: Write our message to the file
	messageBytes := []byte(message)
	messagePtr := uintptr(unsafe.Pointer(&messageBytes[0]))

	// Make the syscall to write data
	// Parameters: SYS_WRITE, file_descriptor, data_pointer, data_length
	syscall.Syscall(
		syscall.SYS_WRITE,        // The syscall we want
		fd,                       // File descriptor from step 1
		messagePtr,               // Pointer to our message
		uintptr(len(message)),    // How many bytes to write
	)

	// Step 3: Close the file
	syscall.Syscall(syscall.SYS_CLOSE, fd, 0, 0)

	// Step 4: Exit the program successfully
	syscall.Syscall(syscall.SYS_EXIT, 0, 0, 0)
}
