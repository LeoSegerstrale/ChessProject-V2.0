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

	}

	return possibleMoves
}
