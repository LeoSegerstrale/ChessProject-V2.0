package service

import (
	"ChessWeb/backend/model"
	"strconv"
)

func BishopMover(location string, board [][]*model.Piece, kingLoc string, rec bool) ([]string, [][][]*model.Piece) {

	currY := int(location[0] - '0')
	currX := int(location[1] - '0')
	var listOBoard [][][]*model.Piece

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

	if rec && kingLoc != "" {
		possibleMoves, listOBoard = validMovesAndBoard(board, []int{currY, currX}, possibleMoves, board[currY][currX], kingLoc)
	}

	return possibleMoves, listOBoard
}

func RookMover(location string, board [][]*model.Piece, kingLoc string, rec bool) ([]string, [][][]*model.Piece) {

	currY := int(location[0] - '0')
	currX := int(location[1] - '0')

	var listOBoard [][][]*model.Piece

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

	if rec && kingLoc != "" {
		possibleMoves, listOBoard = validMovesAndBoard(board, []int{currY, currX}, possibleMoves, board[currY][currX], kingLoc)
	}

	return possibleMoves, listOBoard
}

func KnightMover(location string, board [][]*model.Piece, kingLoc string, rec bool) ([]string, [][][]*model.Piece) {

	currY := int(location[0] - '0')
	currX := int(location[1] - '0')
	var listOBoard [][][]*model.Piece

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

	if rec && kingLoc != "" {
		possibleMoves, listOBoard = validMovesAndBoard(board, []int{currY, currX}, possibleMoves, board[currY][currX], kingLoc)
	}

	return possibleMoves, listOBoard
}

func KingMover(location string, board [][]*model.Piece, castleStat []bool, rookLocs []string, rec bool) ([]string, [][][]*model.Piece) {

	currY := int(location[0] - '0')
	currX := int(location[1] - '0')
	var listOBoard [][][]*model.Piece

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

	if rec {
		possibleMoves, listOBoard = validMovesAndBoard(board, []int{currY, currX}, possibleMoves, board[currY][currX], "!")
	}

	return possibleMoves, listOBoard
}

func PawnMover(location string, board [][]*model.Piece, enPassant string, kingLoc string, rec bool) ([]string, [][][]*model.Piece) {
	currY := int(location[0] - '0')
	currX := int(location[1] - '0')
	var listOBoard [][][]*model.Piece

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

	if rec && kingLoc != "" {
		possibleMoves, listOBoard = validMovesAndBoard(board, []int{currY, currX}, possibleMoves, board[currY][currX], kingLoc)
	}

	return possibleMoves, listOBoard

}

func QueenMover(location string, board [][]*model.Piece, kingLoc string, rec bool) ([]string, [][][]*model.Piece) {

	possibleMoves, boards := BishopMover(location, board, kingLoc, rec)
	extraMoves, extraboards := RookMover(location, board, kingLoc, rec)

	if rec {
		for i, move := range extraMoves {
			possibleMoves = append(possibleMoves, move)
			boards = append(boards, extraboards[i])
		}

	} else {
		for _, move := range extraMoves {
			possibleMoves = append(possibleMoves, move)
		}
	}

	return possibleMoves, boards
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

	rookMoves, _ := RookMover(location, board, location, false)

	if kingMoveHelper(board, rookMoves, colour, "rook") {
		good = false
		board[currY][currX] = prev

		return good

	}

	knightMoves, _ := KnightMover(location, board, location, false)

	if kingMoveHelper(board, knightMoves, colour, "knight") {
		good = false
		board[currY][currX] = prev

		return good
	}

	bishopMoves, _ := BishopMover(location, board, location, false)

	if kingMoveHelper(board, bishopMoves, colour, "bishop") {
		good = false
		board[currY][currX] = prev

		return good
	}

	queenMoves, _ := QueenMover(location, board, location, false)

	if kingMoveHelper(board, queenMoves, colour, "queen") {
		good = false
		board[currY][currX] = prev

		return good
	}

	kingMoves, _ := KingMover(location, board, []bool{false, false, false}, []string{}, false)

	if kingMoveHelper(board, kingMoves, colour, "king") {
		good = false
		board[currY][currX] = prev

		return good
	}

	pawnMoves, _ := PawnMover(location, board, "", location, false)
	if kingMoveHelper(board, pawnMoves, colour, "pawn") {
		good = false
		board[currY][currX] = prev

		return good
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

func validMovesAndBoard(board [][]*model.Piece, fromSquare []int, moves []string, piece *model.Piece, kingLoc string) ([]string, [][][]*model.Piece) {

	var kingLocX int
	var kingLocY int
	var oppColour string
	if kingLoc == "" { // for testing
		return moves, [][][]*model.Piece{}
	}

	possiblePromotions := []string{"queen", "rook", "bishop", "knight"}

	if kingLoc != "!" {
		kingLocX, _ = strconv.Atoi(string(kingLoc[1]))
		kingLocY, _ = strconv.Atoi(string(kingLoc[0]))

	}

	colour := board[fromSquare[0]][fromSquare[1]].Colour

	if colour == "white" {
		oppColour = "black"
	} else {
		oppColour = "white"
	}

	checkedMoves := []string{}
	listOBoards := [][][]*model.Piece{}
	var captured *model.Piece

	for _, move := range moves {
		enPassTrue := false
		promotion := false
		var enPassYSquare int
		if kingLoc == "!" {
			kingLocX, _ = strconv.Atoi(string(move[1]))
			kingLocY, _ = strconv.Atoi(string(move[0]))
		}
		moveX, _ := strconv.Atoi(string(move[1]))
		moveY, _ := strconv.Atoi(string(move[0]))

		if piece.Piece == "pawn" {
			if moveX != fromSquare[1] && board[moveY][moveX] == nil {

				if colour == "white" {
					enPassYSquare = moveY + 1
				} else {
					enPassYSquare = moveY - 1
				}
				board[enPassYSquare][moveX] = nil
				board[fromSquare[0]][fromSquare[1]] = nil
				board[moveY][moveX] = piece
				enPassTrue = true
			} else if (colour == "white" && moveY == 0) || (moveY == 7 && colour == "black") {

				captured = board[moveY][moveX]
				board[fromSquare[0]][fromSquare[1]] = nil
				board[moveY][moveX] = piece

				if kingVMoveChecker(board, kingLocY, kingLocX, colour) {
					for _, piecePromotion := range possiblePromotions {
						board[moveY][moveX] = &model.Piece{
							Piece:  piecePromotion,
							Colour: colour,
						}
						checkedMoves = append(checkedMoves, move)
						listOBoards = append(listOBoards, cloneBoard(board))
						promotion = true

					}
				}
				board[fromSquare[0]][fromSquare[1]] = piece
				board[moveY][moveX] = captured

			} else {

				captured = board[moveY][moveX]
				board[fromSquare[0]][fromSquare[1]] = nil
				board[moveY][moveX] = piece
			}
		} else {

			captured = board[moveY][moveX]
			board[fromSquare[0]][fromSquare[1]] = nil
			board[moveY][moveX] = piece
		}
		if !promotion {
			if kingVMoveChecker(board, kingLocY, kingLocX, colour) {
				checkedMoves = append(checkedMoves, move)
				listOBoards = append(listOBoards, cloneBoard(board))
			}
			board[fromSquare[0]][fromSquare[1]] = piece
			board[moveY][moveX] = captured
			if enPassTrue {
				board[enPassYSquare][moveX] = &model.Piece{
					Piece:  "pawn",
					Colour: oppColour,
				}
			}
		}

	}
	return checkedMoves, listOBoards
}

func cloneBoard(board [][]*model.Piece) [][]*model.Piece {
	newBoard := make([][]*model.Piece, len(board))
	for i := range board {
		newBoard[i] = make([]*model.Piece, len(board[i]))
		for j := range board[i] {
			if board[i][j] != nil {
				pieceCopy := *board[i][j]
				newBoard[i][j] = &pieceCopy
			}
		}
	}
	return newBoard
}
