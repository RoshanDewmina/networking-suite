package main

import (
	"crypto/tls"
	"io"
	"log"
	"net"
)

// startTLSServer starts a TLS server
func startTLSServer() {
	// Load server certificate and key
	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Fatalf("[TLS] Failed to load server certificate and key: %v", err)
	}

	// Configure TLS
	config := &tls.Config{Certificates: []tls.Certificate{cert}}

	// Start listening on port 8443
	ln, err := tls.Listen("tcp", ":8443", config)
	if err != nil {
		log.Fatalf("[TLS] Failed to start TLS server: %v", err)
	}
	defer ln.Close()

	log.Println("[TLS] Server started on port 8443")
	for {
		// Accept incoming connections
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("[TLS] Connection error: %v\n", err)
			continue
		}
		go handleTLSConnection(conn)
	}
}

// handleTLSConnection handles an incoming TLS client connection
func handleTLSConnection(conn net.Conn) {
	defer conn.Close()
	log.Println("[TLS] Client connected")

	buf := make([]byte, 1024)
	for {
		// Read data from the client
		n, err := conn.Read(buf)
		if err == io.EOF {
			log.Println("[TLS] Client disconnected")
			break
		}
		if err != nil {
			log.Printf("[TLS] Read error: %v\n", err)
			break
		}
		message := string(buf[:n])
		log.Printf("[TLS] Received: %s\n", message)

		// Respond to the client
		_, err = conn.Write([]byte("[TLS] Acknowledged"))
		if err != nil {
			log.Printf("[TLS] Write error: %v\n", err)
			break
		}
	}
}

// startTLSClient connects to a TLS server and exchanges messages
func startTLSClient() {
	// Configure TLS
	config := &tls.Config{InsecureSkipVerify: true} // Skip verification for simplicity (not for production)

	// Connect to the server
	conn, err := tls.Dial("tcp", "localhost:8443", config)
	if err != nil {
		log.Fatalf("[TLS] Failed to connect to server: %v", err)
	}
	defer conn.Close()

	log.Println("[TLS] Connected to server")

	// Send a message to the server
	message := "Hello from TLS Client"
	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Fatalf("[TLS] Failed to send message: %v", err)
	}
	log.Printf("[TLS] Sent: %s\n", message)

	// Read the server's response
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatalf("[TLS] Failed to read response: %v", err)
	}
	log.Printf("[TLS] Server response: %s\n", string(buf[:n]))
}
