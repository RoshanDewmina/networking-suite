package main

import (
	"crypto/tls"
	"net"
	"testing"
	"time"
)

func TestUDPIntegration(t *testing.T) {
	// Start UDP server in a goroutine
	go reliableUDPServer()
	// Allow server to start
	time.Sleep(1 * time.Second)

	// Client setup
	serverAddr, _ := net.ResolveUDPAddr("udp", "localhost:9999")
	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		t.Fatalf("Failed to connect to UDP server: %v", err)
	}
	defer conn.Close()

	// Send a test message
	message := "Integration test message"
	_, err = conn.Write([]byte(message))
	if err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	// Receive acknowledgment
	buf := make([]byte, 1024)
	conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	n, _, err := conn.ReadFromUDP(buf)
	if err != nil {
		t.Fatalf("Failed to receive acknowledgment: %v", err)
	}

	response := string(buf[:n])
	if response != "ACK" {
		t.Errorf("Expected 'ACK', but got '%s'", response)
	}
}

func startTLSServer() {
	// Implementation of the TLS server
}

func TestTLSIntegration(t *testing.T) {
	// Start TLS server in a goroutine
	go startTLSServer()
	// Allow server to start
	time.Sleep(1 * time.Second)

	// Client setup
	config := &tls.Config{InsecureSkipVerify: true} // Skip verification for testing
	conn, err := tls.Dial("tcp", "localhost:8443", config)
	if err != nil {
		t.Fatalf("Failed to connect to TLS server: %v", err)
	}
	defer conn.Close()

	// Send a test message
	message := "TLS integration test"
	_, err = conn.Write([]byte(message))
	if err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	// Receive acknowledgment
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		t.Fatalf("Failed to receive acknowledgment: %v", err)
	}

	response := string(buf[:n])
	expected := "[TLS] Acknowledged"
	if response != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, response)
	}
}
