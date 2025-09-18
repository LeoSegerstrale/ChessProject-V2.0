package service

import (
	"ChessWeb/backend/model"
	"math/rand"
	"strconv"
)

func GetBotMove(req model.BotMoveReq) ([][]*model.Piece, string, model.BotMoveReq) {

	ogBoard := req.Board
	currColour := req.Colour

	if currColour == "white-piece" {
		currColour = "white"
	} else if currColour == "black-piece" {
		currColour = "black"
	}
	var possibleBoards [][][]*model.Piece
	var possiblePieces []string
	var moves []string
	var possibleMoves []string
	var fromSquares []string

	for y, row := range ogBoard {
		for x, square := range row {
			if square != nil && square.Colour == currColour {

				currY := strconv.Itoa(y)
				currX := strconv.Itoa(x)
				currLoc := currY + currX
				var boards [][][]*model.Piece

				if square.Piece == "bishop" {
					moves, boards = BishopMover(currLoc, ogBoard, req.KingLoc, true)

				} else if square.Piece == "rook" {
					moves, boards = RookMover(currLoc, ogBoard, req.KingLoc, true)
				} else if square.Piece == "queen" {
					moves, boards = QueenMover(currLoc, ogBoard, req.KingLoc, true)
				} else if square.Piece == "knight" {
					moves, boards = KnightMover(currLoc, ogBoard, req.KingLoc, true)
				} else if square.Piece == "king" {
					moves, boards = KingMover(currLoc, ogBoard, req.CastleStatus, req.RookLocs, true)
				} else if square.Piece == "pawn" {
					moves, boards = PawnMover(currLoc, ogBoard, req.EnPassantReq, req.KingLoc, true)
				}

				if len(boards) != 0 {
					possiblePieces = append(possiblePieces, square.Piece)
				}

				for i, board := range boards {
					possibleBoards = append(possibleBoards, board)
					possiblePieces = append(possiblePieces, square.Piece)
					possibleMoves = append(possibleMoves, moves[i])
					fromSquares = append(fromSquares, strconv.Itoa(y)+strconv.Itoa(x))
				}
			}

		}
	}

	optimalBoard := rand.Intn(len(possibleBoards))

	bestBoard := possibleBoards[optimalBoard]

	if possiblePieces[optimalBoard] == "king" {
		req.KingLoc = possibleMoves[optimalBoard]
		req.CastleStatus[1] = false
	} else if possiblePieces[optimalBoard] == "rook" {

		if fromSquares[optimalBoard] == req.RookLocs[0] {
			req.CastleStatus[0] = false
		} else if fromSquares[optimalBoard] == req.RookLocs[1] {
			req.CastleStatus[2] = false
		}
	} else if possiblePieces[optimalBoard] == "pawn" {
		fromY, _ := strconv.Atoi(string(fromSquares[optimalBoard][0]))
		toY, _ := strconv.Atoi(string(possibleMoves[optimalBoard][0]))
		
		if fromY-toY == 2 || toY-fromY == 2 {
			req.EnPassantReq = possibleMoves[optimalBoard]
		}
	}

	return bestBoard, "", req

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
