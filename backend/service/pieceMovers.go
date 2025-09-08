package service

import (
	"ChessWeb/backend/model"
	"strconv"
)

func BishopMover(location string, board [][]*model.Piece, kingLoc string) []string {

	currY := int(location[0] - '0')
	currX := int(location[1] - '0')

	directions := [][]int{
		{-1, -1},
		{1, -1},
		{-1, 1},
		{1, 1},
	}

	currColour := board[currY][currX].Colour

	possibleMoves := []string{}

	for _, direction := range directions {

		for i := 1; ; i++ {

			newY := direction[0]*i + currY
			newX := direction[1]*i + currX

			if newY < 0 || newY >= len(board) || newX < 0 || newX >= len(board[newY]) {
				break
			}

			if board[newY][newX] == nil {
				newPos := strconv.Itoa(newY) + strconv.Itoa(newX)
				possibleMoves = append(possibleMoves, newPos)

			} else if board[newY][newX].Colour != currColour {
				newPos := strconv.Itoa(newY) + strconv.Itoa(newX)
				possibleMoves = append(possibleMoves, newPos)
				break
			} else {
				break
			}

		}

	}

	return possibleMoves
}

func RookMover(location string, board [][]*model.Piece, kingLoc string) []string {

	currY := int(location[0] - '0')
	currX := int(location[1] - '0')

	directions := [][]int{
		{-1, 0},
		{1, 0},
		{0, 1},
		{0, -1},
	}

	currColour := board[currY][currX].Colour

	possibleMoves := []string{}

	for _, direction := range directions {

		for i := 1; ; i++ {

			newY := direction[0]*i + currY
			newX := direction[1]*i + currX

			if newY < 0 || newY >= len(board) || newX < 0 || newX >= len(board[newY]) {
				break
			}

			if board[newY][newX] == nil {
				newPos := strconv.Itoa(newY) + strconv.Itoa(newX)
				possibleMoves = append(possibleMoves, newPos)

			} else if board[newY][newX].Colour != currColour {
				newPos := strconv.Itoa(newY) + strconv.Itoa(newX)
				possibleMoves = append(possibleMoves, newPos)
				break
			} else {
				break
			}

		}

	}

	return possibleMoves
}

func KnightMover(location string, board [][]*model.Piece, kingLoc string) []string {

	currY := int(location[0] - '0')
	currX := int(location[1] - '0')

	directions := [][]int{
		{2, 1},
		{2, -1},
		{1, 2},
		{1, -2},
		{-2, 1},
		{-2, -1},
		{-1, 2},
		{-1, -2},
	}

	currColour := board[currY][currX].Colour

	possibleMoves := []string{}

	for _, direction := range directions {

		newY := direction[0] + currY
		newX := direction[1] + currX

		if newY < 0 || newY >= len(board) || newX < 0 || newX >= len(board[newY]) {

		} else if board[newY][newX] == nil {
			newPos := strconv.Itoa(newY) + strconv.Itoa(newX)
			possibleMoves = append(possibleMoves, newPos)

		} else if board[newY][newX].Colour != currColour {
			newPos := strconv.Itoa(newY) + strconv.Itoa(newX)
			possibleMoves = append(possibleMoves, newPos)

		}

	}

	return possibleMoves
}

func KingMover(location string, board [][]*model.Piece, castleStat []bool, rookLocs []string, rec bool) []string {

	currY := int(location[0] - '0')
	currX := int(location[1] - '0')

	directions := [][]int{
		{-1, -1},
		{1, -1},
		{-1, 1},
		{1, 1},
		{-1, 0},
		{1, 0},
		{0, 1},
		{0, -1},
	}

	currColour := board[currY][currX].Colour

	possibleMoves := []string{}

	for _, direction := range directions {

		newY := direction[0] + currY
		newX := direction[1] + currX

		if newY < 0 || newY >= len(board) || newX < 0 || newX >= len(board[newY]) {

		} else if board[newY][newX] == nil {
			newPos := strconv.Itoa(newY) + strconv.Itoa(newX)
			possibleMoves = append(possibleMoves, newPos)

		} else if board[newY][newX].Colour != currColour {
			newPos := strconv.Itoa(newY) + strconv.Itoa(newX)
			possibleMoves = append(possibleMoves, newPos)

		}

	}

	if castleStat[1] {
		castleLeft := false
		castleRight := false

		if castleStat[0] {
			castleLeft = true
			rRookLoc := int(rookLocs[1][1] - '0')
			for i := currX - 1; i > 1; i-- {
				if (board[currY][i] != nil || i == rRookLoc) || !kingVMoveChecker(board, currY, i, board[currY][currX].Colour) {
					castleLeft = false

				}
			}
		}
		if castleStat[2] {
			castleRight = true
			lRookLoc := int(rookLocs[0][1] - '0')
			for i := currX + 1; i < 7; i++ {
				if (board[currY][i] != nil || i == lRookLoc) || !kingVMoveChecker(board, currY, i, board[currY][currX].Colour) {
					castleRight = false

				}
			}
		}
		if castleLeft {
			possibleMoves = append(possibleMoves, strconv.Itoa(currY)+strconv.Itoa(2))

		}
		if castleRight {
			possibleMoves = append(possibleMoves, strconv.Itoa(currY)+strconv.Itoa(6))

		}
	}

	return possibleMoves
}

func PawnMover(location string, board [][]*model.Piece, enPassant string, kingLoc string) []string {
	currY := int(location[0] - '0')
	currX := int(location[1] - '0')

	var newY int
	var newX int

	blackTrue := board[currY][currX].Colour == "black"
	var oppColour string

	possibleMoves := []string{}

	var direction int

	if blackTrue {
		direction = 1
		oppColour = "white"

	} else {
		direction = -1
		oppColour = "black"
	}

	sUp := []int{1 * direction, 0}
	dUp := []int{2 * direction, 0}

	lTake := []int{1 * direction, -1}
	rTake := []int{1 * direction, 1}

	if board[currY+sUp[0]][currX] == nil {
		newY = currY + sUp[0]
		newX = currX + sUp[1]

		newPos := strconv.Itoa(newY) + strconv.Itoa(newX)
		possibleMoves = append(possibleMoves, newPos)

		if (currY == 1 && blackTrue || currY == 6 && !blackTrue) && board[currY+dUp[0]][currX+dUp[1]] == nil {
			newY = currY + dUp[0]
			newX = currX + dUp[1]

			newPos := strconv.Itoa(newY) + strconv.Itoa(newX)
			possibleMoves = append(possibleMoves, newPos)
		}

	}

	if currX != 0 && board[currY+lTake[0]][currX+lTake[1]] != nil && board[currY+lTake[0]][currX+lTake[1]].Colour == oppColour {
		newY = currY + lTake[0]
		newX = currX + lTake[1]

		newPos := strconv.Itoa(newY) + strconv.Itoa(newX)
		possibleMoves = append(possibleMoves, newPos)
	}
	if currX != 7 && board[currY+rTake[0]][currX+rTake[1]] != nil && board[currY+rTake[0]][currX+rTake[1]].Colour == oppColour {
		newY = currY + rTake[0]
		newX = currX + rTake[1]

		newPos := strconv.Itoa(newY) + strconv.Itoa(newX)
		possibleMoves = append(possibleMoves, newPos)
	}

	if enPassant != "" {
		enPasY := int(enPassant[0] - '0')
		enPasX := int(enPassant[1] - '0')

		if enPasY == currY {
			if enPasX == currX+1 {
				newY = currY + rTake[0]
				newX = currX + rTake[1]

				newPos := strconv.Itoa(newY) + strconv.Itoa(newX)
				possibleMoves = append(possibleMoves, newPos)
			} else if enPasX == currX-1 {
				newY = currY + lTake[0]
				newX = currX + lTake[1]

				newPos := strconv.Itoa(newY) + strconv.Itoa(newX)
				possibleMoves = append(possibleMoves, newPos)
			}
		}

	}

	return possibleMoves

}

func QueenMover(location string, board [][]*model.Piece, kingLoc string) []string {

	possibleMoves := BishopMover(location, board, kingLoc)
	extraMoves := RookMover(location, board, kingLoc)

	for _, move := range extraMoves {
		possibleMoves = append(possibleMoves, move)
	}
	return possibleMoves
}

func kingVMoveChecker(board [][]*model.Piece, currY int, currX int, colour string) bool {

	good := true

	strCurrY := strconv.Itoa(currY)
	strCurrX := strconv.Itoa(currX)
	location := strCurrY + strCurrX

	prev := board[currY][currX]

	board[currY][currX] = &model.Piece{
		Piece:  "",
		Colour: colour,
	}

	rookMoves := RookMover(location, board, location)

	if kingMoveHelper(board, rookMoves, colour, "rook") {
		good = false
	}

	knightMoves := KnightMover(location, board, location)

	if good && kingMoveHelper(board, knightMoves, colour, "knight") {
		good = false
	}

	bishopMoves := BishopMover(location, board, location)

	if good && kingMoveHelper(board, bishopMoves, colour, "bishop") {
		good = false
	}

	queenMoves := QueenMover(location, board, location)

	if good && kingMoveHelper(board, queenMoves, colour, "queen") {
		good = false
	}

	kingMoves := KingMover(location, board, []bool{false, false, false}, []string{}, true)

	if good && kingMoveHelper(board, kingMoves, colour, "king") {
		good = false
	}

	board[currY][currX] = prev

	return good
}

func kingMoveHelper(board [][]*model.Piece, moves []string, colour string, piece string) bool {
	for _, move := range moves {
		moveX, _ := strconv.Atoi(string(move[1]))
		moveY, _ := strconv.Atoi(string(move[0]))

		if board[moveY][moveX] != nil && board[moveY][moveX].Colour != colour && board[moveY][moveX].Piece == piece {
			return true
		}
	}

	return false
}

func returnBoard(board [][]*model.Piece, fromSquare []int, toSquare []int, piece *model.Piece, kingLoc string) (bool, [][]*model.Piece) {
	return false, board
}
