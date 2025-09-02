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
