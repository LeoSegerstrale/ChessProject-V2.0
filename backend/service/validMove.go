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
		possibleMoves = BishopMover(req.From, board, req.KingLoc)

	} else if piece.Piece == "rook" {
		possibleMoves = RookMover(req.From, board, req.KingLoc)
	} else if piece.Piece == "queen" {
		possibleMoves = QueenMover(req.From, board, req.KingLoc)
	} else if piece.Piece == "knight" {
		possibleMoves = KnightMover(req.From, board, req.KingLoc)
	} else if piece.Piece == "king" {
		possibleMoves = KingMover(req.From, board, req.CastleStatus, req.RookLocs, false)
	} else if piece.Piece == "pawn" {
		possibleMoves = PawnMover(req.From, board, req.EnPassantReq, req.KingLoc)
	}

	return possibleMoves
}
