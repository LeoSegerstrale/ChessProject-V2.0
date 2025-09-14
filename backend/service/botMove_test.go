package service

import (
	"ChessWeb/backend/model"
	"reflect"
	"testing"
)

func TestBotMove(t *testing.T) {
	tests := []struct {
		name string
		req  model.BotMoveReq
		want *model.BotMoveResp
	}{
		{
			name: "test1",
			req: model.BotMoveReq{
				Board:        EmptyBoard([]int{6, 3}),
				Colour:       "white",
				EnPassantReq: "",
				CastleStatus: []bool{false, false, false},
				RookLocs:     []string{"70", "77"},
				KingLoc:      "74",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := GetBotMove(tt.req)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got = %v, want %v", got, tt.want)
			}
		})
	}

}
