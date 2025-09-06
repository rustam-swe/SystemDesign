/*
 * 01_tcp_client.go
 *
 * Basic TCP client using only syscalls
 * This program demonstrates how to:
 * - Create a TCP socket
 * - Connect to a server
 * - Send data
 * - Receive data
 * - Close the connection
 *
 * Syscalls used:
 * - SYS_SOCKET (create socket)
 * - SYS_CONNECT (connect to server)
 * - SYS_WRITE (send data)
 * - SYS_READ (receive data)
 * - SYS_CLOSE (close socket)
 */

package main

import (
	"syscall"
	"unsafe"
)

func main() {
	// Step 1: Create a TCP socket
	// socket(domain, type, protocol)
	// AF_INET = IPv4, SOCK_STREAM = TCP, 0 = default protocol
	socketFD, _, err := syscall.Syscall(
		syscall.SYS_SOCKET,
		syscall.AF_INET,     // IPv4
		syscall.SOCK_STREAM, // TCP
		0,                   // Default protocol (TCP for SOCK_STREAM)
	)

	if err != 0 {
		errorMsg := "Error: Could not create socket\n"
		errorBytes := []byte(errorMsg)
		errorPtr := uintptr(unsafe.Pointer(&errorBytes[0]))
		syscall.Syscall(syscall.SYS_WRITE, 2, errorPtr, uintptr(len(errorMsg)))
		syscall.Syscall(syscall.SYS_EXIT, 1, 0, 0)
	}

	// Step 2: Set up server address
	// We'll connect to 127.0.0.1:8080 (localhost:8080)
	// This is where our HTTP server will be running

	// Create sockaddr_in structure for IPv4
	// struct sockaddr_in {
	//     sa_family_t    sin_family; // Address family (AF_INET)
	//     in_port_t      sin_port;   // Port number (in network byte order)
	//     struct in_addr sin_addr;   // Internet address (in network byte order)
	//     char           sin_zero[8]; // Padding
	// };

	serverAddr := make([]byte, 16) // Size of sockaddr_in

	// sin_family = AF_INET (2 bytes, little endian on most systems)
	serverAddr[0] = byte(syscall.AF_INET & 0xFF)
	serverAddr[1] = byte((syscall.AF_INET >> 8) & 0xFF)

	// sin_port = 8080 in network byte order (big endian)
	port := uint16(8080)
	serverAddr[2] = byte((port >> 8) & 0xFF) // High byte first
	serverAddr[3] = byte(port & 0xFF)        // Low byte second

	// sin_addr = 127.0.0.1 (localhost) in network byte order
	serverAddr[4] = 127 // 127.0.0.1
	serverAddr[5] = 0
	serverAddr[6] = 0
	serverAddr[7] = 1

	// sin_zero = padding (already zero-initialized)

	// Step 3: Connect to the server
	serverAddrPtr := uintptr(unsafe.Pointer(&serverAddr[0]))
	_, _, err = syscall.Syscall(
		syscall.SYS_CONNECT,
		socketFD,                 // Socket file descriptor
		serverAddrPtr,            // Server address
		uintptr(len(serverAddr)), // Address length
	)

	if err != 0 {
		errorMsg := "Error: Could not connect to server\n"
		errorMsg += "Make sure the server is running on localhost:8080\n"
		errorBytes := []byte(errorMsg)
		errorPtr := uintptr(unsafe.Pointer(&errorBytes[0]))
		syscall.Syscall(syscall.SYS_WRITE, 2, errorPtr, uintptr(len(errorMsg)))
		syscall.Syscall(syscall.SYS_CLOSE, socketFD, 0, 0)
		syscall.Syscall(syscall.SYS_EXIT, 1, 0, 0)
	}

	// Step 4: Send a simple message
	message := "Hello from TCP client!\n"
	messageBytes := []byte(message)
	messagePtr := uintptr(unsafe.Pointer(&messageBytes[0]))

	_, _, err = syscall.Syscall(
		syscall.SYS_WRITE,
		socketFD,              // Socket file descriptor
		messagePtr,            // Message to send
		uintptr(len(message)), // Message length
	)

	if err != 0 {
		errorMsg := "Error: Could not send data\n"
		errorBytes := []byte(errorMsg)
		errorPtr := uintptr(unsafe.Pointer(&errorBytes[0]))
		syscall.Syscall(syscall.SYS_WRITE, 2, errorPtr, uintptr(len(errorMsg)))
		syscall.Syscall(syscall.SYS_CLOSE, socketFD, 0, 0)
		syscall.Syscall(syscall.SYS_EXIT, 1, 0, 0)
	}

	// Print success message
	successMsg := "Sent message to server\n"
	successBytes := []byte(successMsg)
	successPtr := uintptr(unsafe.Pointer(&successBytes[0]))
	syscall.Syscall(syscall.SYS_WRITE, 1, successPtr, uintptr(len(successMsg)))

	// Step 5: Receive response from server
	buffer := make([]byte, 1024) // Buffer to receive data
	bufferPtr := uintptr(unsafe.Pointer(&buffer[0]))

	bytesReceived, _, err := syscall.Syscall(
		syscall.SYS_READ,
		socketFD,             // Socket file descriptor
		bufferPtr,            // Buffer to receive into
		uintptr(len(buffer)), // Maximum bytes to receive
	)

	if err != 0 || bytesReceived == 0 {
		errorMsg := "Error: Could not receive data or connection closed\n"
		errorBytes := []byte(errorMsg)
		errorPtr := uintptr(unsafe.Pointer(&errorBytes[0]))
		syscall.Syscall(syscall.SYS_WRITE, 2, errorPtr, uintptr(len(errorMsg)))
		syscall.Syscall(syscall.SYS_CLOSE, socketFD, 0, 0)
		syscall.Syscall(syscall.SYS_EXIT, 1, 0, 0)
	}

	// Display received data
	responseMsg := "Received from server:\n"
	responseMsgBytes := []byte(responseMsg)
	responseMsgPtr := uintptr(unsafe.Pointer(&responseMsgBytes[0]))
	syscall.Syscall(syscall.SYS_WRITE, 1, responseMsgPtr, uintptr(len(responseMsg)))

	// Print the actual response
	syscall.Syscall(syscall.SYS_WRITE, 1, bufferPtr, bytesReceived)

	// Step 6: Close the socket
	syscall.Syscall(syscall.SYS_CLOSE, socketFD, 0, 0)

	// Success message
	doneMsg := "\nTCP client completed successfully!\n"
	doneMsgBytes := []byte(doneMsg)
	doneMsgPtr := uintptr(unsafe.Pointer(&doneMsgBytes[0]))
	syscall.Syscall(syscall.SYS_WRITE, 1, doneMsgPtr, uintptr(len(doneMsg)))

	// Exit successfully
	syscall.Syscall(syscall.SYS_EXIT, 0, 0, 0)
}
