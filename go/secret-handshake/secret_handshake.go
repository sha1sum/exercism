package secret

// Return a set of "moves" for a secret handshake
func Handshake(handshake int) []string {
	if handshake < 0 { handshake = 0 }
	// Make a slice to hold our moves
	var moves []string = make([]string, 0)
	// Iterate over bits and if they're on then add the move
	if handshake&1 == 1 { moves = append(moves, "wink") }
	if handshake&2 == 2 { moves = append(moves, "double blink") }
	if handshake&4 == 4 { moves = append(moves, "close your eyes") }
	if handshake&8 == 8 { moves = append(moves, "jump") }
	// Reverse the order if the reverse bit is on
	if handshake&16 == 16 {
		for i, j := 0, len(moves) - 1; i < j; i, j = i + 1, j - 1 {
			moves[i], moves[j] = moves[j], moves[i]
		}
	}
	return moves
}