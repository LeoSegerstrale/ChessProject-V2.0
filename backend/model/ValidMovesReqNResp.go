package model

type VMoveReq struct {
	From string `json:"from"`

	Board [][]*Piece `json:"board"`
}

type VMoveResp struct {
	ValidSquares []string `json:"validSquares"`
}

type Piece struct {
	Piece  string
	Colour string
}

type Board struct {
	Grid [][]*Piece
}
