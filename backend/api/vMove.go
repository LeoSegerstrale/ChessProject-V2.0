package api

import (
	"ChessWeb/backend/model"
	"ChessWeb/backend/service"
	"encoding/json"
	"net/http"
)

func VMoveCheck(w http.ResponseWriter, r *http.Request) {

	var req model.VMoveReq
	var resp model.VMoveResp

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	moves := service.GetValidMoves(req)

	resp = model.VMoveResp{
		ValidSquares: moves,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)

	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func BotMove(w http.ResponseWriter, r *http.Request) {

	var req model.BotMoveReq
	var resp model.BotMoveResp
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}

	board := [][]*model.Piece{}

	resp = model.BotMoveResp{
		Board: board,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)

	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
}
