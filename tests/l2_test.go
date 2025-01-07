package main

import "testing"

func TestL2Tunnel(t *testing.T) {
	data := []byte("test data")
	encapsulated := L2Tunnel(data)
	if len(encapsulated) <= len(data) {
		t.Errorf("Encapsulation failed, expected length greater than %d, got %d", len(data), len(encapsulated))
	}

	parsed := ParseL2Tunnel(encapsulated)
	if string(parsed) != string(data) {
		t.Errorf("Decapsulation failed, expected %s, got %s", string(data), string(parsed))
	}
}

func TestParseL2TunnelShortFrame(t *testing.T) {
	frame := []byte{0xAA, 0xBB} // Too short
	parsed := ParseL2Tunnel(frame)
	if parsed != nil {
		t.Error("Expected nil for short frame, got non-nil result")
	}
}

func L2Tunnel(data []byte) []byte {
	// Dummy implementation for L2Tunnel
	return append([]byte{0x00}, data...)
}

func ParseL2Tunnel(data []byte) []byte {
	// Dummy implementation for ParseL2Tunnel
	if len(data) < 3 {
		return nil
	}
	return data[1:]
}
