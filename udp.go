package main

import (
	"log"
	"net"
	"time"
)

// reliableUDPServer starts a reliable UDP server
func reliableUDPServer() {
	log.Println("[UDP] Starting Reliable UDP Server")
	// Listen on UDP port
	addr, err := net.ResolveUDPAddr("udp", ":9999")
	if err != nil {
		log.Fatalf("[UDP] Failed to resolve UDP address: %v", err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalf("[UDP] Failed to start server: %v", err)
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		n, clientAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Printf("[UDP] Read error: %v\n", err)
			continue
		}
		log.Printf("[UDP] Received: %s\n", string(buf[:n]))

		// Send acknowledgment back to the client
		ack := []byte("ACK")
		_, err = conn.WriteToUDP(ack, clientAddr)
		if err != nil {
			log.Printf("[UDP] Failed to send ACK: %v\n", err)
		} else {
			log.Printf("[UDP] Sent ACK to %s\n", clientAddr.String())
		}
	}
}

// reliableUDPClient sends a message and waits for an acknowledgment
func reliableUDPClient() {
	log.Println("[UDP] Starting Reliable UDP Client")
	// Resolve server address
	serverAddr, err := net.ResolveUDPAddr("udp", "localhost:9999")
	if err != nil {
		log.Fatalf("[UDP] Failed to resolve server address: %v", err)
	}
	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		log.Fatalf("[UDP] Failed to connect to server: %v", err)
	}
	defer conn.Close()

	// Send a message to the server
	message := "Hello from UDP Client"
	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Fatalf("[UDP] Failed to send message: %v", err)
	}
	log.Printf("[UDP] Sent: %s\n", message)

	// Wait for acknowledgment
	buf := make([]byte, 1024)
	conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	n, _, err := conn.ReadFromUDP(buf)
	if err != nil {
		log.Printf("[UDP] Failed to receive acknowledgment: %v\n", err)
		return
	}
	log.Printf("[UDP] Received acknowledgment: %s\n", string(buf[:n]))
}
