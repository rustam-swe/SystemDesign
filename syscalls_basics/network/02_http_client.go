/*
 * 02_http_client.go
 *
 * HTTP client using only syscalls
 * This program demonstrates how to:
 * - Create a TCP socket
 * - Connect to an HTTP server
 * - Send a proper HTTP GET request
 * - Receive and parse HTTP response
 * - Handle basic HTTP protocol
 *
 * Syscalls used:
 * - SYS_SOCKET (create socket)
 * - SYS_CONNECT (connect to server)
 * - SYS_WRITE (send HTTP request)
 * - SYS_READ (receive HTTP response)
 * - SYS_CLOSE (close socket)
 */

package main

import (
	"syscall"
	"unsafe"
)

func main() {
	// Step 1: Create a TCP socket
	socketFD, _, err := syscall.Syscall(
		syscall.SYS_SOCKET,
		syscall.AF_INET,     // IPv4
		syscall.SOCK_STREAM, // TCP
		0,                   // Default protocol
	)

	if err != 0 {
		errorMsg := "Error: Could not create socket\n"
		errorBytes := []byte(errorMsg)
		errorPtr := uintptr(unsafe.Pointer(&errorBytes[0]))
		syscall.Syscall(syscall.SYS_WRITE, 2, errorPtr, uintptr(len(errorMsg)))
		syscall.Syscall(syscall.SYS_EXIT, 1, 0, 0)
	}

	// Step 2: Set up server address (127.0.0.1:8080)
	serverAddr := make([]byte, 16) // Size of sockaddr_in

	// sin_family = AF_INET
	serverAddr[0] = byte(syscall.AF_INET & 0xFF)
	serverAddr[1] = byte((syscall.AF_INET >> 8) & 0xFF)

	// sin_port = 8080 in network byte order (big endian)
	port := uint16(8080)
	serverAddr[2] = byte((port >> 8) & 0xFF) // High byte first
	serverAddr[3] = byte(port & 0xFF)        // Low byte second

	// sin_addr = 127.0.0.1 (localhost)
	serverAddr[4] = 127 // 127.0.0.1
	serverAddr[5] = 0
	serverAddr[6] = 0
	serverAddr[7] = 1

	// Step 3: Connect to the server
	serverAddrPtr := uintptr(unsafe.Pointer(&serverAddr[0]))
	_, _, err = syscall.Syscall(
		syscall.SYS_CONNECT,
		socketFD,
		serverAddrPtr,
		uintptr(len(serverAddr)),
	)

	if err != 0 {
		errorMsg := "Error: Could not connect to server\n"
		errorMsg += "Make sure HTTP server is running on localhost:8080\n"
		errorBytes := []byte(errorMsg)
		errorPtr := uintptr(unsafe.Pointer(&errorBytes[0]))
		syscall.Syscall(syscall.SYS_WRITE, 2, errorPtr, uintptr(len(errorMsg)))
		syscall.Syscall(syscall.SYS_CLOSE, socketFD, 0, 0)
		syscall.Syscall(syscall.SYS_EXIT, 1, 0, 0)
	}

	// Step 4: Send HTTP GET request
	// A proper HTTP/1.1 GET request format:
	// GET /path HTTP/1.1\r\n
	// Host: hostname\r\n
	// Connection: close\r\n
	// \r\n
	httpRequest := "GET / HTTP/1.1\r\n"
	httpRequest += "Host: localhost:8080\r\n"
	httpRequest += "User-Agent: Simple-Syscall-Client/1.0\r\n"
	httpRequest += "Accept: text/html\r\n"
	httpRequest += "Connection: close\r\n"
	httpRequest += "\r\n" // Empty line indicates end of headers

	requestBytes := []byte(httpRequest)
	requestPtr := uintptr(unsafe.Pointer(&requestBytes[0]))

	_, _, err = syscall.Syscall(
		syscall.SYS_WRITE,
		socketFD,
		requestPtr,
		uintptr(len(httpRequest)),
	)

	if err != 0 {
		errorMsg := "Error: Could not send HTTP request\n"
		errorBytes := []byte(errorMsg)
		errorPtr := uintptr(unsafe.Pointer(&errorBytes[0]))
		syscall.Syscall(syscall.SYS_WRITE, 2, errorPtr, uintptr(len(errorMsg)))
		syscall.Syscall(syscall.SYS_CLOSE, socketFD, 0, 0)
		syscall.Syscall(syscall.SYS_EXIT, 1, 0, 0)
	}

	// Print what we sent
	sentMsg := "Sent HTTP request:\n"
	sentMsgBytes := []byte(sentMsg)
	sentMsgPtr := uintptr(unsafe.Pointer(&sentMsgBytes[0]))
	syscall.Syscall(syscall.SYS_WRITE, 1, sentMsgPtr, uintptr(len(sentMsg)))
	syscall.Syscall(syscall.SYS_WRITE, 1, requestPtr, uintptr(len(httpRequest)))

	// Step 5: Receive HTTP response
	response := make([]byte, 4096) // Buffer for HTTP response
	responsePtr := uintptr(unsafe.Pointer(&response[0]))
	totalReceived := 0

	// Read response in chunks (HTTP responses can be large)
	for totalReceived < len(response)-1 {
		bytesReceived, _, err := syscall.Syscall(
			syscall.SYS_READ,
			socketFD,
			uintptr(unsafe.Pointer(&response[totalReceived])),
			uintptr(len(response)-totalReceived-1),
		)

		if err != 0 {
			// Error occurred
			break
		}

		if bytesReceived == 0 {
			// Server closed connection (normal for Connection: close)
			break
		}

		totalReceived += int(bytesReceived)
	}

	if totalReceived == 0 {
		errorMsg := "Error: No data received from server\n"
		errorBytes := []byte(errorMsg)
		errorPtr := uintptr(unsafe.Pointer(&errorBytes[0]))
		syscall.Syscall(syscall.SYS_WRITE, 2, errorPtr, uintptr(len(errorMsg)))
		syscall.Syscall(syscall.SYS_CLOSE, socketFD, 0, 0)
		syscall.Syscall(syscall.SYS_EXIT, 1, 0, 0)
	}

	// Null-terminate the response for safety
	response[totalReceived] = 0

	// Step 6: Display HTTP response
	responseHeaderMsg := "\n=== HTTP RESPONSE ===\n"
	responseHeaderBytes := []byte(responseHeaderMsg)
	responseHeaderPtr := uintptr(unsafe.Pointer(&responseHeaderBytes[0]))
	syscall.Syscall(syscall.SYS_WRITE, 1, responseHeaderPtr, uintptr(len(responseHeaderMsg)))

	// Print the response
	syscall.Syscall(syscall.SYS_WRITE, 1, responsePtr, uintptr(totalReceived))

	// Step 7: Basic HTTP response parsing
	// Let's extract the status code for educational purposes
	parseMsg := "\n=== BASIC RESPONSE PARSING ===\n"
	parseMsgBytes := []byte(parseMsg)
	parseMsgPtr := uintptr(unsafe.Pointer(&parseMsgBytes[0]))
	syscall.Syscall(syscall.SYS_WRITE, 1, parseMsgPtr, uintptr(len(parseMsg)))

	// Look for "HTTP/1.1 " at the beginning
	if totalReceived >= 12 {
		statusMsg := "HTTP Status: "
		statusMsgBytes := []byte(statusMsg)
		statusMsgPtr := uintptr(unsafe.Pointer(&statusMsgBytes[0]))
		syscall.Syscall(syscall.SYS_WRITE, 1, statusMsgPtr, uintptr(len(statusMsg)))

		// Extract status code (characters 9, 10, 11 after "HTTP/1.1 ")
		statusCode := make([]byte, 4)
		statusCode[0] = response[9]  // First digit
		statusCode[1] = response[10] // Second digit
		statusCode[2] = response[11] // Third digit
		statusCode[3] = '\n'         // Newline

		statusCodePtr := uintptr(unsafe.Pointer(&statusCode[0]))
		syscall.Syscall(syscall.SYS_WRITE, 1, statusCodePtr, 4)
	}

	// Step 8: Close the socket
	syscall.Syscall(syscall.SYS_CLOSE, socketFD, 0, 0)

	// Success message
	doneMsg := "\nHTTP client completed successfully!\n"
	doneMsgBytes := []byte(doneMsg)
	doneMsgPtr := uintptr(unsafe.Pointer(&doneMsgBytes[0]))
	syscall.Syscall(syscall.SYS_WRITE, 1, doneMsgPtr, uintptr(len(doneMsg)))

	// Exit successfully
	syscall.Syscall(syscall.SYS_EXIT, 0, 0, 0)
}
