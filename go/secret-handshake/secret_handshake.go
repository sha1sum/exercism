package secret

import "fmt"

// A string of all-on bits (all moves and reversed)
const allOn string = "11111"

// Return a set of "moves" for a secret handshake
func Handshake(handshake int) []string {
	if handshake < 0 { handshake = 0 }
	handshakeString := fmt.Sprintf("%05b", handshake)
	return binaryHandshake(handshakeString)
}

// Take a string representation of binary and convert it into a moves
func binaryHandshake(handshake string) []string {
	// Whether or not we should reverse the moves (initializing here)
	reverse := false
	// Truncate the handshake to a maximum of 5 "bits"/"moves"
	if len(handshake) > 5 {
		handshake = handshake[len(handshake) - 5:]
	}
	// Make a slice to hold our moves
	var moves []string = make([]string, 0)
	// Iterate over each bit
	for i := 0; i < 5; i++ {
		// Whether the current bit is on
		match := handshake[i] == allOn[i]
		// Iterate over bits and if they're on then add the move
		switch {
		case i == 0 && match:
			reverse = true
		case i == 1 && match:
			moves = append(moves, "jump")
		case i == 2 && match:
			moves = append(moves, "close your eyes")
		case i == 3 && match:
			moves = append(moves, "double blink")
		case i == 4 && match:
			moves = append(moves, "wink")
		}
	}
	// Since we went start to end in the binary string, we want to reverse the moves unless the fifth bit was on.
	if !reverse {
		for i, j := 0, len(moves)-1; i < j; i, j = i+1, j-1 {
			moves[i], moves[j] = moves[j], moves[i]
		}
	}
	return moves
}