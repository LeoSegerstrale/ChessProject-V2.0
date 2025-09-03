package service

import (
	"ChessWeb/backend/model"
	"strconv"
)

func BishopMover(location string, board [][]*model.Piece) []string {

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

func RookMover(location string, board [][]*model.Piece) []string {

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

func KnightMover(location string, board [][]*model.Piece) []string {

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

func KingMover(location string, board [][]*model.Piece) []string {

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

	return possibleMoves
}

func PawnMover(location string, board [][]*model.Piece) []string {
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

	return possibleMoves

}
