package main

import (
	"net"
	"testing"
	"time"
)

func reliableUDPServer() {
	addr, err := net.ResolveUDPAddr("udp", ":9999")
	if err != nil {
		panic(err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		n, clientAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			continue
		}

		message := string(buf[:n])
		if message == "Test UDP message" {
			conn.WriteToUDP([]byte("ACK"), clientAddr)
		}
	}
}

func TestReliableUDPServerClient(t *testing.T) {
	// Start UDP server in a goroutine
	go reliableUDPServer()
	// Allow the server some time to start
	time.Sleep(1 * time.Second)

	// Connect to the server as a client
	serverAddr, _ := net.ResolveUDPAddr("udp", "localhost:9999")
	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	// Send a message to the server
	message := "Test UDP message"
	_, err = conn.Write([]byte(message))
	if err != nil {
		t.Fatalf("Failed to send message to server: %v", err)
	}

	// Wait for acknowledgment from the server
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
