package main

import (
	"encoding/binary"
	"log"
)

// GenerateIPPacket creates a mock Layer 3 IP packet
func GenerateIPPacket(payload []byte) []byte {
	log.Println("[IP] Creating IP packet")
	// Mock IP header (20 bytes)
	header := make([]byte, 20)
	header[0] = 0x45 // Version (4) and Header Length (5)
	header[1] = 0x00 // Differentiated Services Field
	binary.BigEndian.PutUint16(header[2:4], uint16(len(payload)+20)) // Total Length
	header[4], header[5] = 0, 0   // Identification
	header[6], header[7] = 0x40, 0 // Flags and Fragment Offset
	header[8] = 64                 // TTL
	header[9] = 6                  // Protocol (TCP)
	binary.BigEndian.PutUint16(header[10:12], 0) // Header Checksum (not calculated here)
	// Mock Source and Destination IPs
	copy(header[12:16], []byte{192, 168, 0, 1}) // Source IP
	copy(header[16:20], []byte{192, 168, 0, 2}) // Destination IP

	return append(header, payload...)
}

// ParseIPPacket extracts the payload from an IP packet
func ParseIPPacket(packet []byte) []byte {
	log.Println("[IP] Parsing IP packet")
	// Validate packet length
	if len(packet) < 20 {
		log.Println("[IP] Packet too short!")
		return nil
	}
	// Extract and return the payload
	return packet[20:]
}
