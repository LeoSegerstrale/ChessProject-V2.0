package model

type VMoveReq struct {
	From         string     `json:"from" validate:"required"`
	Board        [][]*Piece `json:"board" validate:"required"`
	EnPassantReq string     `json:"enPassant"`
	CastleStatus []bool     `json:"castleStatus"`
	RookLocs     []string   `json:"rookLocs"`
	KingLoc      string     `json:"kingLoc"`
}

type VMoveResp struct {
	ValidSquares []string `json:"validSquares"`
}

type BotMoveReq struct {
	Board        [][]*Piece `json:"board" validate:"required"`
	Colour       string     `json:"colour" validate:"required"`
	EnPassantReq string     `json:"enPassant"`
	CastleStatus []bool     `json:"castleStatus"`
	RookLocs     []string   `json:"rookLocs"`
	KingLoc      string     `json:"kingLoc"`
}

type BotMoveResp struct {
	Board     [][]*Piece `json:"board"`
	Checkmate []string   `json:"checkmate"`
}

type Piece struct {
	Piece  string
	Colour string
}
