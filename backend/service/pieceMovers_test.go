package service

import (
	"ChessWeb/backend/model"
	"reflect"
	"testing"
)

func TestBishopMover(t *testing.T) {
	tests := []struct {
		name     string
		location string
		board    [][]*model.Piece
		want     []string
	}{
		{
			name:     "Empty board center",
			location: "34",
			board:    emptyBoard(),
			want:     []string{"23", "12", "01", "43", "52", "61", "70", "25", "16", "07", "45", "56", "67"},
		},
		{
			name:     "Blocked by same color",
			location: "34",
			board:    boardWithFBlockers(),
			want:     []string{},
		},
		{
			name:     "Partially blocked by opposite colour",
			location: "34",
			board:    boardWithPBlockers(),
			want:     []string{"23", "43", "25", "45"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BishopMover(tt.location, tt.board)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got = %v, want %v", got, tt.want)
			}
		})
	}
}

//HELPER FUNCTIONS

func emptyBoard() [][]*model.Piece {
	board := make([][]*model.Piece, 8)
	for i := range board {
		board[i] = make([]*model.Piece, 8)
	}
	board[3][4] = &model.Piece{
		Piece:  "",
		Colour: "white",
	}
	return board
}

func boardWithFBlockers() [][]*model.Piece {

	oppPiece := &model.Piece{
		Piece:  "",
		Colour: "white",
	}
	board := make([][]*model.Piece, 8)
	for i := range board {
		board[i] = make([]*model.Piece, 8)
		for j := range board[i] {
			board[i][j] = oppPiece
		}
	}
	return board
}

func boardWithPBlockers() [][]*model.Piece {

	oppPiece := &model.Piece{
		Piece:  "",
		Colour: "black",
	}
	samePiece := &model.Piece{
		Piece:  "",
		Colour: "white",
	}

	board := make([][]*model.Piece, 8)
	for i := range board {
		board[i] = make([]*model.Piece, 8)
		for j := range board[i] {
			if i == 3 && j == 4 {
				board[i][j] = samePiece
			} else {
				board[i][j] = oppPiece
			}

		}
	}
	return board
}
