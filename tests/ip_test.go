package main

import (
	"testing"
)

// ParseIPPacket extracts the payload from a given IP packet
func ParseIPPacket(packet []byte) []byte {
	if len(packet) < 20 {
		return nil // Packet too short to contain a valid header
	}
	return packet[20:] // Extract payload assuming the header is 20 bytes
}

// GenerateIPPacket creates a simple IP packet with the given payload
func GenerateIPPacket(payload []byte) []byte {
	header := make([]byte, 20) // Simplified header with a fixed length of 20 bytes
	return append(header, payload...)
}

func TestGenerateIPPacket(t *testing.T) {
	payload := []byte("test payload")
	packet := GenerateIPPacket(payload)

	// Check if the packet length is correct
	expectedLength := len(payload) + 20 // 20 bytes for the header
	if len(packet) != expectedLength {
		t.Errorf("Incorrect packet length, expected %d, got %d", expectedLength, len(packet))
	}

	// Verify the payload is correctly embedded
	if string(packet[20:]) != string(payload) {
		t.Errorf("Payload mismatch, expected %s, got %s", string(payload), string(packet[20:]))
	}
}

func TestParseIPPacket(t *testing.T) {
	payload := []byte("test payload")
	packet := GenerateIPPacket(payload)

	// Extract payload from the packet
	extractedPayload := ParseIPPacket(packet)
	if string(extractedPayload) != string(payload) {
		t.Errorf("Payload mismatch, expected %s, got %s", string(payload), string(extractedPayload))
	}
}

func TestParseIPPacketShortPacket(t *testing.T) {
	shortPacket := []byte{0x45} // Too short to be valid
	extractedPayload := ParseIPPacket(shortPacket)
	if extractedPayload != nil {
		t.Error("Expected nil for short packet, got non-nil result")
	}
}
