package service

import (
	"ChessWeb/backend/model"
	"strconv"
)

func GetValidMoves(req model.VMoveReq) []string {

	board := req.Board
	currY, _ := strconv.Atoi(string(req.From[0]))

	currX, _ := strconv.Atoi(string(req.From[1]))

	piece := req.Board[currY][currX]

	if piece == nil {

		return nil
	}

	var possibleMoves []string

	if piece.Piece == "bishop" {
		possibleMoves = BishopMover(req.From, board)

	} else if piece.Piece == "rook" {
		possibleMoves = RookMover(req.From, board)
	} else if piece.Piece == "queen" {
		possibleMoves = BishopMover(req.From, board)
		extraMoves := RookMover(req.From, board)

		for _, move := range extraMoves {
			possibleMoves = append(possibleMoves, move)
		}
	} else if piece.Piece == "knight" {
		possibleMoves = KnightMover(req.From, board)
	} else if piece.Piece == "king" {
		possibleMoves = KingMover(req.From, board)
	}

	return possibleMoves
}
