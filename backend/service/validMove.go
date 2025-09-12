package service

import (
	"ChessWeb/backend/model"
	"strconv"
)

func GetBotMove(req model.BotMoveReq) ([][]*model.Piece, []string) {

	ogBoard := req.Board
	currColour := req.Colour
	var possibleBoards [][][]*model.Piece

	for y, row := range ogBoard {
		for x, square := range row {
			if square != nil && square.Colour == currColour {

				currY := strconv.Itoa(y)
				currX := strconv.Itoa(x)
				currLoc := currY + currX
				var boards [][][]*model.Piece

				if square.Piece == "bishop" {
					_, boards = BishopMover(currLoc, ogBoard, req.KingLoc, true)

				} else if square.Piece == "rook" {
					_, boards = RookMover(currLoc, ogBoard, req.KingLoc, true)
				} else if square.Piece == "queen" {
					_, boards = QueenMover(currLoc, ogBoard, req.KingLoc, true)
				} else if square.Piece == "knight" {
					_, boards = KnightMover(currLoc, ogBoard, req.KingLoc, true)
				} else if square.Piece == "king" {
					_, boards = KingMover(currLoc, ogBoard, req.CastleStatus, req.RookLocs, true)
				} else if square.Piece == "pawn" {
					_, boards = PawnMover(currLoc, ogBoard, req.EnPassantReq, req.KingLoc, true)
				}

				for _, board := range boards {
					possibleBoards = append(possibleBoards, board)
				}

			}
		}
	}
	return possibleBoards[0], []string{""}

}

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
