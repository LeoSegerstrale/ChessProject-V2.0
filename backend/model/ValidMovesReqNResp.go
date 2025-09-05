package model

type VMoveReq struct {
	From         string     `json:"from" validate:"required"`
	Board        [][]*Piece `json:"board" validate:"required"`
	EnPassantReq string     `json:"enPassant"`
	CastleStatus []bool     `json:"castleStatus"`
}

type VMoveResp struct {
	ValidSquares []string `json:"validSquares"`
}

type Piece struct {
	Piece  string
	Colour string
}
