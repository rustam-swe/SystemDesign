package main

import (
	"syscall"
	"unsafe"
)

func main() {
	filename := "demo.txt"
	message := "This is a complete syscall demo!\nWe created, wrote, and will read this file.\n"

	// === PART 1: CREATE AND WRITE ===

	// Convert filename to syscall format
	filenameBytes := []byte(filename + "\x00")
	filenamePtr := uintptr(unsafe.Pointer(&filenameBytes[0]))

	// Create/open file for writing
	fd, _, err := syscall.Syscall(
		syscall.SYS_OPEN,
		filenamePtr,
		syscall.O_CREAT|syscall.O_WRONLY|syscall.O_TRUNC,
		0644,
	)

	if err != 0 {
		errorMsg := "Failed to create file\n"
		errorBytes := []byte(errorMsg)
		errorPtr := uintptr(unsafe.Pointer(&errorBytes[0]))
		syscall.Syscall(syscall.SYS_WRITE, 2, errorPtr, uintptr(len(errorMsg)))
		syscall.Syscall(syscall.SYS_EXIT, 1, 0, 0)
	}

	// Write message to file
	messageBytes := []byte(message)
	messagePtr := uintptr(unsafe.Pointer(&messageBytes[0]))

	_, _, err = syscall.Syscall(
		syscall.SYS_WRITE,
		fd,
		messagePtr,
		uintptr(len(message)),
	)

	if err != 0 {
		errorMsg := "Failed to write to file\n"
		errorBytes := []byte(errorMsg)
		errorPtr := uintptr(unsafe.Pointer(&errorBytes[0]))
		syscall.Syscall(syscall.SYS_WRITE, 2, errorPtr, uintptr(len(errorMsg)))
		syscall.Syscall(syscall.SYS_CLOSE, fd, 0, 0)
		syscall.Syscall(syscall.SYS_EXIT, 1, 0, 0)
	}

	// Close file after writing
	syscall.Syscall(syscall.SYS_CLOSE, fd, 0, 0)

	// Print success message
	successMsg := "Successfully wrote file! Now reading it back...\n\n"
	successBytes := []byte(successMsg)
	successPtr := uintptr(unsafe.Pointer(&successBytes[0]))
	syscall.Syscall(syscall.SYS_WRITE, 1, successPtr, uintptr(len(successMsg)))

	// === PART 2: READ THE FILE ===

	// Open file for reading
	fd, _, err = syscall.Syscall(
		syscall.SYS_OPEN,
		filenamePtr,
		syscall.O_RDONLY,
		0,
	)

	if err != 0 {
		errorMsg := "Failed to open file for reading\n"
		errorBytes := []byte(errorMsg)
		errorPtr := uintptr(unsafe.Pointer(&errorBytes[0]))
		syscall.Syscall(syscall.SYS_WRITE, 2, errorPtr, uintptr(len(errorMsg)))
		syscall.Syscall(syscall.SYS_EXIT, 1, 0, 0)
	}

	// Read from file
	buffer := make([]byte, 200)
	bufferPtr := uintptr(unsafe.Pointer(&buffer[0]))

	bytesRead, _, err := syscall.Syscall(
		syscall.SYS_READ,
		fd,
		bufferPtr,
		uintptr(len(buffer)),
	)

	if err != 0 {
		errorMsg := "Failed to read file\n"
		errorBytes := []byte(errorMsg)
		errorPtr := uintptr(unsafe.Pointer(&errorBytes[0]))
		syscall.Syscall(syscall.SYS_WRITE, 2, errorPtr, uintptr(len(errorMsg)))
		syscall.Syscall(syscall.SYS_CLOSE, fd, 0, 0)
		syscall.Syscall(syscall.SYS_EXIT, 1, 0, 0)
	}

	// Display what we read
	headerMsg := "File contents:\n"
	headerBytes := []byte(headerMsg)
	headerPtr := uintptr(unsafe.Pointer(&headerBytes[0]))
	syscall.Syscall(syscall.SYS_WRITE, 1, headerPtr, uintptr(len(headerMsg)))

	// Print the actual file content
	syscall.Syscall(syscall.SYS_WRITE, 1, bufferPtr, bytesRead)

	// Print final message
	finalMsg := "\nDemo completed successfully!\n"
	finalBytes := []byte(finalMsg)
	finalPtr := uintptr(unsafe.Pointer(&finalBytes[0]))
	syscall.Syscall(syscall.SYS_WRITE, 1, finalPtr, uintptr(len(finalMsg)))

	// Close file and exit
	syscall.Syscall(syscall.SYS_CLOSE, fd, 0, 0)
	syscall.Syscall(syscall.SYS_EXIT, 0, 0, 0)
}
