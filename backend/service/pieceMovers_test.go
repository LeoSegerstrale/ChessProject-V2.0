package service

import (
	"ChessWeb/backend/model"
	"reflect"
	"testing"
)

func TestBishopMover(t *testing.T) {
	ogLoc := []int{3, 4}
	tests := []struct {
		name     string
		location string
		board    [][]*model.Piece
		want     []string
		kingLoc  string
	}{
		{
			name:     "Empty board center",
			location: "34",
			board:    emptyBoard(ogLoc),
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
			got := BishopMover(tt.location, tt.board, tt.kingLoc)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRookMover(t *testing.T) {
	ogLoc := []int{3, 4}
	tests := []struct {
		name     string
		location string
		board    [][]*model.Piece
		want     []string
		kingLoc  string
	}{
		{
			name:     "Empty board center",
			location: "34",
			board:    emptyBoard(ogLoc),
			want:     []string{"24", "14", "04", "44", "54", "64", "74", "35", "36", "37", "33", "32", "31", "30"},
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
			want:     []string{"24", "44", "35", "33"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RookMover(tt.location, tt.board, tt.kingLoc)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKnightMover(t *testing.T) {
	ogLoc := []int{3, 4}
	tests := []struct {
		name     string
		location string
		board    [][]*model.Piece
		want     []string
		kingLoc  string
	}{
		{
			name:     "Empty board center",
			location: "34",
			board:    emptyBoard(ogLoc),
			want:     []string{"55", "53", "46", "42", "15", "13", "26", "22"},
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
			want:     []string{"55", "53", "46", "42", "15", "13", "26", "22"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := KnightMover(tt.location, tt.board, tt.kingLoc)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKingMover(t *testing.T) {
	ogLoc := []int{3, 4}
	casLoc := []int{7, 4}

	tests := []struct {
		name       string
		location   string
		board      [][]*model.Piece
		want       []string
		castleStat []bool
		rookLocs   []string
	}{
		{
			name:       "If Castling blocked by pieceMovement",
			location:   "74",
			board:      boardWithBishopBlockCastleV2(),
			want:       []string{"63", "65", "64", "75", "73", "72"},
			castleStat: []bool{true, true, true},
			rookLocs:   []string{"70", "77"},
		},
		{
			name:       "Empty board center",
			location:   "34",
			board:      emptyBoard(ogLoc),
			want:       []string{"23", "43", "25", "45", "24", "44", "35", "33"},
			castleStat: []bool{false, false, false},
			rookLocs:   []string{},
		},
		{
			name:       "Blocked by same color",
			location:   "74",
			board:      boardWithFBlockers(),
			want:       []string{},
			castleStat: []bool{true, true, true},
			rookLocs:   []string{"70", "77"},
		},
		{
			name:       "Partially blocked by opposite colour",
			location:   "34",
			board:      boardWithPBlockers(),
			want:       []string{"23", "43", "25", "45", "24", "44", "35", "33"},
			castleStat: []bool{false, false, false},
			rookLocs:   []string{},
		}, {
			name:       "Castling works",
			location:   "74",
			board:      emptyBoard(casLoc),
			want:       []string{"63", "65", "64", "75", "73", "72", "76"},
			castleStat: []bool{true, true, true},
			rookLocs:   []string{"70", "77"},
		},
		{
			name:       "If Castling unavailable it does not say it is",
			location:   "74",
			board:      emptyBoard(casLoc),
			want:       []string{"63", "65", "64", "75", "73"},
			castleStat: []bool{false, true, false},
			rookLocs:   []string{"70", "77"},
		},
		{
			name:       "If Castling unavailable in one direction, it can only castle in one direction",
			location:   "74",
			board:      emptyBoard(casLoc),
			want:       []string{"63", "65", "64", "75", "73", "72"},
			castleStat: []bool{true, true, false},
			rookLocs:   []string{"70", "77"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := KingMover(tt.location, tt.board, tt.castleStat, tt.rookLocs, false)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKingMoverHelpers(t *testing.T) {
	tests := []struct {
		name   string
		currY  int
		currX  int
		colour string
		board  [][]*model.Piece
		want   bool
	}{
		{
			name:   "Empty board center",
			board:  emptyBoard([]int{7, 5}),
			currY:  7,
			currX:  5,
			want:   true,
			colour: "white",
		},
		{
			name:   "Blocked by piece",
			board:  boardWithBishopBlockCastle(),
			currY:  7,
			currX:  5,
			want:   false,
			colour: "white",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := kingVMoveChecker(tt.board, tt.currY, tt.currX, tt.colour)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestPawnMover(t *testing.T) {
	ogLoc := []int{6, 0}
	tests := []struct {
		name      string
		location  string
		board     [][]*model.Piece
		want      []string
		enPassant string
		kingLoc   string
	}{
		{
			name:      "pawn in second rank can jump two or one square",
			location:  "60",
			board:     emptyBoard(ogLoc),
			want:      []string{"50", "40"},
			enPassant: "",
		},
		{
			name:      "Pawn with blockers either side can take",
			location:  "34",
			board:     boardWithPBlockers(),
			want:      []string{"23", "25"},
			enPassant: "",
		},
		{
			name:      "Pawn with no available squares has no available squares",
			location:  "34",
			board:     boardWithFBlockers(),
			want:      []string{},
			enPassant: "",
		},
		{
			name:      "En Passant test baby",
			location:  "34",
			board:     boardWithEnPassant(),
			want:      []string{"23"},
			enPassant: "33",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PawnMover(tt.location, tt.board, tt.enPassant, tt.kingLoc)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got = %v, want %v", got, tt.want)
			}
		})
	}
}

//HELPER FUNCTIONS

func emptyBoard(pPos []int) [][]*model.Piece {
	board := make([][]*model.Piece, 8)
	for i := range board {
		board[i] = make([]*model.Piece, 8)
	}
	board[pPos[0]][pPos[1]] = &model.Piece{
		Piece:  "",
		Colour: "white",
	}
	return board
}

func boardWithBishopBlockCastleV2() [][]*model.Piece {
	board := make([][]*model.Piece, 8)
	for i := range board {
		board[i] = make([]*model.Piece, 8)
	}
	board[4][2] = &model.Piece{
		Piece:  "bishop",
		Colour: "black",
	}
	board[7][4] = &model.Piece{
		Piece:  "king",
		Colour: "white",
	}
	return board
}

func boardWithBishopBlockCastle() [][]*model.Piece {
	board := make([][]*model.Piece, 8)
	for i := range board {
		board[i] = make([]*model.Piece, 8)
	}
	board[4][2] = &model.Piece{
		Piece:  "bishop",
		Colour: "black",
	}
	board[7][5] = &model.Piece{
		Piece:  "king",
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

func boardWithEnPassant() [][]*model.Piece {

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
			} else if i == 3 && j == 3 {
				board[i][j] = oppPiece
			} else if i == 2 && j == 4 {
				board[i][j] = oppPiece
			}

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
