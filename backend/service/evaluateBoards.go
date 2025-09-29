package service

import (
	"ChessWeb/backend/model"
)

func evaluateBoards(possibleBoards [][][]*model.Piece, colour string) int {
	var boardEvals []int
	var biggest int

	for _, board := range possibleBoards {
		eval := evaluateMaterial(board, colour)
		boardEvals = append(boardEvals, eval)
		if eval > biggest {
			biggest = eval
		}
	}
	for i, eval := range boardEvals {
		if eval == biggest {
			return i
		}
	}

	return 0

}

func evaluateMaterial(board [][]*model.Piece, colour string) int {
	pieceVals := map[string]int{
		"pawn":   1,
		"knight": 3,
		"bishop": 3,
		"rook":   5,
		"queen":  9,
	}
	var materialValue int
	for _, row := range board {
		for _, square := range row {
			if square != nil {
				if square.Colour == colour {
					materialValue += pieceVals[square.Piece]

				} else {
					materialValue -= pieceVals[square.Piece]
				}
			}

		}
	}
	return materialValue

}
