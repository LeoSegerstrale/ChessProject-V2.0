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
		possibleMoves, _ = BishopMover(req.From, board, req.KingLoc, true)

	} else if piece.Piece == "rook" {
		possibleMoves, _ = RookMover(req.From, board, req.KingLoc, true)
	} else if piece.Piece == "queen" {
		possibleMoves, _ = QueenMover(req.From, board, req.KingLoc, true)
	} else if piece.Piece == "knight" {
		possibleMoves, _ = KnightMover(req.From, board, req.KingLoc, true)
	} else if piece.Piece == "king" {
		possibleMoves, _ = KingMover(req.From, board, req.CastleStatus, req.RookLocs, true)
	} else if piece.Piece == "pawn" {
		possibleMoves, _ = PawnMover(req.From, board, req.EnPassantReq, req.KingLoc, true)
	}

	return possibleMoves
}
