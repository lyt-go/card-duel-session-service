package handler

import (
	"net/http"

	"cardgame/internal/model"
	"cardgame/pkg/httpx"
)

func (s *Server) registerPlayRecordRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/play-records", s.createPlayRecord)
	mux.HandleFunc("GET /api/play-records", s.listPlayRecord)
	mux.HandleFunc("GET /api/play-records/{id}", s.getPlayRecord)
	mux.HandleFunc("DELETE /api/play-records/{id}", s.deletePlayRecord)
}

type createPlayRecordRequest struct {
	GameID string `json:"game_id"`
	DeckID string `json:"deck_id"`
	CardID string `json:"card_id"`
	Turn   int    `json:"turn"`
}

func (s *Server) createPlayRecord(w http.ResponseWriter, r *http.Request) {
	var req createPlayRecordRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	x, err := s.svc.CreatePlayRecord(model.PlayRecord{GameID: req.GameID, DeckID: req.DeckID, CardID: req.CardID, Turn: req.Turn})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.Created(w, x)
}

func (s *Server) listPlayRecord(w http.ResponseWriter, r *http.Request) {
	pp := httpx.ParsePagination(r, 20, s.maxPageSize())
	filter := model.PlayRecordFilter{
		GameID: r.URL.Query().Get("game_id"),
		DeckID: r.URL.Query().Get("deck_id"),
	}
	items, total, err := s.svc.ListPlayRecords(filter, pp.Page, pp.Size)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, httpx.PageResult{
		Items:      items,
		Pagination: httpx.Pagination{Page: pp.Page, Size: pp.Size, Total: total},
	})
}

func (s *Server) getPlayRecord(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	x, err := s.svc.GetPlayRecord(id)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, x)
}

func (s *Server) deletePlayRecord(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := s.svc.DeletePlayRecord(id); err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.NoContent(w)
}
