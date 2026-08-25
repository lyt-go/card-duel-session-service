package handler

import (
	"net/http"

	"cardgame/internal/model"
	"cardgame/pkg/httpx"
)

func (s *Server) registerGameRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/games", s.createGame)
	mux.HandleFunc("GET /api/games", s.listGame)
	mux.HandleFunc("GET /api/games/{id}", s.getGame)
	mux.HandleFunc("PUT /api/games/{id}", s.updateGame)
	mux.HandleFunc("DELETE /api/games/{id}", s.deleteGame)
	mux.HandleFunc("PATCH /api/games/{id}/status", s.transitionGame)
}

type createGameRequest struct {
	DeckAID string `json:"deck_a_id"`
	DeckBID string `json:"deck_b_id"`
}

func (s *Server) createGame(w http.ResponseWriter, r *http.Request) {
	var req createGameRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	x, err := s.svc.CreateGame(model.Game{DeckAID: req.DeckAID, DeckBID: req.DeckBID})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.Created(w, x)
}

func (s *Server) listGame(w http.ResponseWriter, r *http.Request) {
	pp := httpx.ParsePagination(r, 20, s.maxPageSize())
	filter := model.GameFilter{
		Status:  r.URL.Query().Get("status"),
		DeckAID: r.URL.Query().Get("deck_a_id"),
	}
	items, total, err := s.svc.ListGames(filter, pp.Page, pp.Size)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, httpx.PageResult{
		Items:      items,
		Pagination: httpx.Pagination{Page: pp.Page, Size: pp.Size, Total: total},
	})
}

func (s *Server) getGame(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	x, err := s.svc.GetGame(id)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, x)
}

type updateGameRequest struct {
	Turn     int    `json:"turn"`
	WinnerID string `json:"winner_id"`
}

func (s *Server) updateGame(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var req updateGameRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	x, err := s.svc.UpdateGame(id, model.Game{Turn: req.Turn, WinnerID: req.WinnerID})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, x)
}

func (s *Server) deleteGame(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := s.svc.DeleteGame(id); err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.NoContent(w)
}

type transitionGameRequest struct {
	Status string `json:"status"`
}

func (s *Server) transitionGame(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var req transitionGameRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	x, err := s.svc.TransitionGame(id, req.Status)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, x)
}
