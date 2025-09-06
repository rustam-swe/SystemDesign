package main

import (
	"syscall"
	"unsafe"
)

func main() {
	// The message we want to display
	message := "Hello, World from syscalls!\n"

	// Convert our message to the format syscalls need
	messageBytes := []byte(message)
	messagePtr := uintptr(unsafe.Pointer(&messageBytes[0]))

	// Use the SYS_WRITE syscall to write to stdout
	// Parameters: SYS_WRITE, file_descriptor, data_pointer, data_length
	// File descriptor 1 = stdout (the screen)
	syscall.Syscall(
		syscall.SYS_WRITE,     // What syscall we want to make
		1,                     // Write to stdout (screen)
		messagePtr,            // Pointer to our message
		uintptr(len(message)), // How many bytes to write
	)

	// Exit the program successfully
	// Parameter: exit code (0 = success)
	syscall.Syscall(syscall.SYS_EXIT, 0, 0, 0)
}
