package stack

func ValidTicTacToe(board []string) bool {
	xCount, oCount := 0, 0

	// Count X and O
	for _, row := range board {
		for _, ch := range row {
			if ch == 'X' {
				xCount++
			} else if ch == 'O' {
				oCount++
			}
		}
	}

	// Basic count validation
	if oCount > xCount || xCount > oCount+1 {
		return false
	}

	xWin := isWinner(board, 'X')
	oWin := isWinner(board, 'O')

	// Both can't win
	if xWin && oWin {
		return false
	}

	// If X wins, must have one more move than O
	if xWin && xCount != oCount+1 {
		return false
	}

	// If O wins, must have same count
	if oWin && xCount != oCount {
		return false
	}

	return true
}

func isWinner(board []string, ch rune) bool {
	// Check rows and columns
	for i := 0; i < 3; i++ {
		if board[i][0] == byte(ch) && board[i][1] == byte(ch) && board[i][2] == byte(ch) {
			return true
		}
		if board[0][i] == byte(ch) && board[1][i] == byte(ch) && board[2][i] == byte(ch) {
			return true
		}
	}
	// Check diagonals
	if board[0][0] == byte(ch) && board[1][1] == byte(ch) && board[2][2] == byte(ch) {
		return true
	}
	if board[0][2] == byte(ch) && board[1][1] == byte(ch) && board[2][0] == byte(ch) {
		return true
	}
	return false
}
