/*
 * 03_http_server.go
 *
 * HTTP server using only syscalls
 * This program demonstrates how to:
 * - Create a TCP server socket
 * - Bind to an address and port
 * - Listen for incoming connections
 * - Accept client connections
 * - Parse basic HTTP requests
 * - Send HTTP responses
 * - Handle multiple clients (one at a time)
 *
 * Syscalls used:
 * - SYS_SOCKET (create socket)
 * - SYS_BIND (bind to address/port)
 * - SYS_LISTEN (listen for connections)
 * - SYS_ACCEPT (accept client connections)
 * - SYS_READ (receive HTTP requests)
 * - SYS_WRITE (send HTTP responses)
 * - SYS_CLOSE (close sockets)
 */

package main

import (
	"syscall"
	"unsafe"
)

func main() {
	// Step 1: Create a TCP server socket
	serverFD, _, err := syscall.Syscall(
		syscall.SYS_SOCKET,
		syscall.AF_INET,     // IPv4
		syscall.SOCK_STREAM, // TCP
		0,                   // Default protocol
	)

	if err != 0 {
		errorMsg := "Error: Could not create server socket\n"
		errorBytes := []byte(errorMsg)
		errorPtr := uintptr(unsafe.Pointer(&errorBytes[0]))
		syscall.Syscall(syscall.SYS_WRITE, 2, errorPtr, uintptr(len(errorMsg)))
		syscall.Syscall(syscall.SYS_EXIT, 1, 0, 0)
	}

	// Step 2: Set up server address (bind to 0.0.0.0:8080)
	// This means listen on all network interfaces on port 8080
	serverAddr := make([]byte, 16) // Size of sockaddr_in

	// sin_family = AF_INET
	serverAddr[0] = byte(syscall.AF_INET & 0xFF)
	serverAddr[1] = byte((syscall.AF_INET >> 8) & 0xFF)

	// sin_port = 8080 in network byte order (big endian)
	port := uint16(8080)
	serverAddr[2] = byte((port >> 8) & 0xFF) // High byte first
	serverAddr[3] = byte(port & 0xFF)        // Low byte second

	// sin_addr = 0.0.0.0 (INADDR_ANY - listen on all interfaces)
	serverAddr[4] = 0 // 0.0.0.0
	serverAddr[5] = 0
	serverAddr[6] = 0
	serverAddr[7] = 0

	// Step 3: Bind socket to address and port
	serverAddrPtr := uintptr(unsafe.Pointer(&serverAddr[0]))
	_, _, err = syscall.Syscall(
		syscall.SYS_BIND,
		serverFD,                 // Server socket file descriptor
		serverAddrPtr,            // Server address
		uintptr(len(serverAddr)), // Address length
	)

	if err != 0 {
		errorMsg := "Error: Could not bind to port 8080\n"
		errorMsg += "Make sure port 8080 is not already in use\n"
		errorBytes := []byte(errorMsg)
		errorPtr := uintptr(unsafe.Pointer(&errorBytes[0]))
		syscall.Syscall(syscall.SYS_WRITE, 2, errorPtr, uintptr(len(errorMsg)))
		syscall.Syscall(syscall.SYS_CLOSE, serverFD, 0, 0)
		syscall.Syscall(syscall.SYS_EXIT, 1, 0, 0)
	}

	// Step 4: Listen for incoming connections
	_, _, err = syscall.Syscall(
		syscall.SYS_LISTEN,
		serverFD, // Server socket
		5,        // Backlog (max 5 pending connections)
		0,
	)

	if err != 0 {
		errorMsg := "Error: Could not listen on socket\n"
		errorBytes := []byte(errorMsg)
		errorPtr := uintptr(unsafe.Pointer(&errorBytes[0]))
		syscall.Syscall(syscall.SYS_WRITE, 2, errorPtr, uintptr(len(errorMsg)))
		syscall.Syscall(syscall.SYS_CLOSE, serverFD, 0, 0)
		syscall.Syscall(syscall.SYS_EXIT, 1, 0, 0)
	}

	// Server is ready
	startMsg := "HTTP Server started on http://127.0.0.1:8080\n"
	startMsg += "Press Ctrl+C to stop the server\n"
	startMsg += "Waiting for connections...\n\n"
	startMsgBytes := []byte(startMsg)
	startMsgPtr := uintptr(unsafe.Pointer(&startMsgBytes[0]))
	syscall.Syscall(syscall.SYS_WRITE, 1, startMsgPtr, uintptr(len(startMsg)))

	// Step 5: Main server loop - accept and handle clients
	for {
		// Accept incoming connection
		clientAddr := make([]byte, 16) // Client address will be stored here
		clientAddrLen := uintptr(len(clientAddr))
		clientAddrPtr := uintptr(unsafe.Pointer(&clientAddr[0]))
		clientAddrLenPtr := uintptr(unsafe.Pointer(&clientAddrLen))

		clientFD, _, err := syscall.Syscall(
			syscall.SYS_ACCEPT,
			serverFD,         // Server socket
			clientAddrPtr,    // Client address (output)
			clientAddrLenPtr, // Address length (input/output)
		)

		if err != 0 {
			errorMsg := "Error: Could not accept client connection\n"
			errorBytes := []byte(errorMsg)
			errorPtr := uintptr(unsafe.Pointer(&errorBytes[0]))
			syscall.Syscall(syscall.SYS_WRITE, 2, errorPtr, uintptr(len(errorMsg)))
			continue // Try to accept next connection
		}

		// Client connected
		clientMsg := "Client connected, handling request...\n"
		clientMsgBytes := []byte(clientMsg)
		clientMsgPtr := uintptr(unsafe.Pointer(&clientMsgBytes[0]))
		syscall.Syscall(syscall.SYS_WRITE, 1, clientMsgPtr, uintptr(len(clientMsg)))

		// Step 6: Read HTTP request from client
		requestBuffer := make([]byte, 4096) // Buffer for HTTP request
		requestPtr := uintptr(unsafe.Pointer(&requestBuffer[0]))

		bytesReceived, _, err := syscall.Syscall(
			syscall.SYS_READ,
			clientFD,                      // Client socket
			requestPtr,                    // Buffer to read into
			uintptr(len(requestBuffer)-1), // Leave space for null terminator
		)

		if err != 0 || bytesReceived == 0 {
			errorMsg := "Error: Could not read from client\n"
			errorBytes := []byte(errorMsg)
			errorPtr := uintptr(unsafe.Pointer(&errorBytes[0]))
			syscall.Syscall(syscall.SYS_WRITE, 2, errorPtr, uintptr(len(errorMsg)))
			syscall.Syscall(syscall.SYS_CLOSE, clientFD, 0, 0)
			continue
		}

		// Null-terminate the request
		requestBuffer[bytesReceived] = 0

		// Print received request
		reqMsg := "Received HTTP request:\n"
		reqMsgBytes := []byte(reqMsg)
		reqMsgPtr := uintptr(unsafe.Pointer(&reqMsgBytes[0]))
		syscall.Syscall(syscall.SYS_WRITE, 1, reqMsgPtr, uintptr(len(reqMsg)))
		syscall.Syscall(syscall.SYS_WRITE, 1, requestPtr, bytesReceived)
		syscall.Syscall(syscall.SYS_WRITE, 1, uintptr(unsafe.Pointer(&[]byte{'\n'}[0])), 1)

		// Step 7: Create and send HTTP response
		// Simple HTTP/1.1 response with HTML content
		httpResponse := "HTTP/1.1 200 OK\r\n"
		httpResponse += "Content-Type: text/html\r\n"
		httpResponse += "Connection: close\r\n"
		httpResponse += "Server: Simple-Syscall-Server/1.0\r\n"

		// HTML content
		htmlContent := "<html><head><title>Syscall HTTP Server</title></head>"
		htmlContent += "<body>"
		htmlContent += "<h1>Hello from Syscall HTTP Server!</h1>"
		htmlContent += "<p>This response was generated using only system calls.</p>"
		htmlContent += "<p>No standard library HTTP functions were used!</p>"
		htmlContent += "<hr>"
		htmlContent += "<p><strong>Server Info:</strong></p>"
		htmlContent += "<ul>"
		htmlContent += "<li>Language: Go</li>"
		htmlContent += "<li>Method: Direct syscalls only</li>"
		htmlContent += "<li>Port: 8080</li>"
		htmlContent += "</ul>"
		htmlContent += "</body></html>"

		// Add Content-Length header
		contentLengthStr := "Content-Length: "
		contentLength := len(htmlContent)

		// Convert content length to string manually
		if contentLength == 0 {
			contentLengthStr += "0"
		} else {
			lengthDigits := make([]byte, 0, 10)
			temp := contentLength
			for temp > 0 {
				lengthDigits = append([]byte{byte('0' + temp%10)}, lengthDigits...)
				temp /= 10
			}
			contentLengthStr += string(lengthDigits)
		}
		contentLengthStr += "\r\n"

		httpResponse += contentLengthStr
		httpResponse += "\r\n" // Empty line separates headers from body
		httpResponse += htmlContent

		// Send the response
		responseBytes := []byte(httpResponse)
		responsePtr := uintptr(unsafe.Pointer(&responseBytes[0]))

		_, _, err = syscall.Syscall(
			syscall.SYS_WRITE,
			clientFD,                   // Client socket
			responsePtr,                // Response to send
			uintptr(len(httpResponse)), // Response length
		)

		if err != 0 {
			errorMsg := "Error: Could not send response to client\n"
			errorBytes := []byte(errorMsg)
			errorPtr := uintptr(unsafe.Pointer(&errorBytes[0]))
			syscall.Syscall(syscall.SYS_WRITE, 2, errorPtr, uintptr(len(errorMsg)))
		} else {
			sentMsg := "Sent HTTP response to client\n\n"
			sentMsgBytes := []byte(sentMsg)
			sentMsgPtr := uintptr(unsafe.Pointer(&sentMsgBytes[0]))
			syscall.Syscall(syscall.SYS_WRITE, 1, sentMsgPtr, uintptr(len(sentMsg)))
		}

		// Step 8: Close client connection
		syscall.Syscall(syscall.SYS_CLOSE, clientFD, 0, 0)

		// Ready for next connection
		readyMsg := "Ready for next connection...\n"
		readyMsgBytes := []byte(readyMsg)
		readyMsgPtr := uintptr(unsafe.Pointer(&readyMsgBytes[0]))
		syscall.Syscall(syscall.SYS_WRITE, 1, readyMsgPtr, uintptr(len(readyMsg)))
	}

	// This code is never reached, but good practice
	syscall.Syscall(syscall.SYS_CLOSE, serverFD, 0, 0)
	syscall.Syscall(syscall.SYS_EXIT, 0, 0, 0)
}
