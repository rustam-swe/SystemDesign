package main

import (
	"syscall"
	"unsafe"
)

func main() {
	// The file we want to read from
	filename := "test.txt"

	// Step 1: Open the file for reading
	// Convert filename to format syscalls understand
	filenameBytes := []byte(filename + "\x00") // Add null terminator
	filenamePtr := uintptr(unsafe.Pointer(&filenameBytes[0]))

	// Make the syscall to open the file
	// Parameters: SYS_OPEN, filename, flags
	fd, _, err := syscall.Syscall(
		syscall.SYS_OPEN,     // The syscall we want
		filenamePtr,          // Pointer to filename
		syscall.O_RDONLY,     // Open for reading only
		0,                    // No permissions needed for reading
	)

	// Check if opening the file failed
	if err != 0 {
		// Write error message to stderr (file descriptor 2)
		errorMsg := "Error: Could not open file\n"
		errorBytes := []byte(errorMsg)
		errorPtr := uintptr(unsafe.Pointer(&errorBytes[0]))
		syscall.Syscall(syscall.SYS_WRITE, 2, errorPtr, uintptr(len(errorMsg)))
		syscall.Syscall(syscall.SYS_EXIT, 1, 0, 0)
	}

	// Step 2: Read data from the file
	// Create a buffer to hold the data we read
	buffer := make([]byte, 100) // Buffer can hold up to 100 bytes
	bufferPtr := uintptr(unsafe.Pointer(&buffer[0]))

	// Make the syscall to read data
	// Parameters: SYS_READ, file_descriptor, buffer_pointer, buffer_size
	bytesRead, _, err := syscall.Syscall(
		syscall.SYS_READ,        // The syscall we want
		fd,                      // File descriptor from step 1
		bufferPtr,               // Pointer to our buffer
		uintptr(len(buffer)),    // Maximum bytes to read
	)

	// Check if reading failed
	if err != 0 {
		// Write error and exit
		errorMsg := "Error: Could not read file\n"
		errorBytes := []byte(errorMsg)
		errorPtr := uintptr(unsafe.Pointer(&errorBytes[0]))
		syscall.Syscall(syscall.SYS_WRITE, 2, errorPtr, uintptr(len(errorMsg)))
		syscall.Syscall(syscall.SYS_CLOSE, fd, 0, 0)
		syscall.Syscall(syscall.SYS_EXIT, 1, 0, 0)
	}

	// Step 3: Write what we read to stdout (screen)
	// We only want to write the actual bytes we read, not the whole buffer
	syscall.Syscall(
		syscall.SYS_WRITE,    // The syscall we want
		1,                    // File descriptor 1 = stdout (screen)
		bufferPtr,            // Pointer to our data
		bytesRead,            // How many bytes we actually read
	)

	// Step 4: Close the file
	syscall.Syscall(syscall.SYS_CLOSE, fd, 0, 0)

	// Step 5: Exit the program successfully
	syscall.Syscall(syscall.SYS_EXIT, 0, 0, 0)
}
