package main

import "log"

// L2Tunnel encapsulates data into a Layer 2 frame
func L2Tunnel(data []byte) []byte {
	log.Println("[L2] Encapsulating data into L2 frame")
	// Add a mock Ethernet header (6 bytes source MAC, 6 bytes dest MAC, 2 bytes ethertype)
	header := []byte{0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xFF, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x08, 0x00}
	return append(header, data...)
}

// ParseL2Tunnel decapsulates a Layer 2 frame
func ParseL2Tunnel(frame []byte) []byte {
	log.Println("[L2] Decapsulating L2 frame")
	// Validate frame length
	if len(frame) < 14 {
		log.Println("[L2] Frame too short!")
		return nil
	}
	// Strip the Ethernet header
	return frame[14:]
}
