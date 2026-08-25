package handler

import (
	"net/http"

	"cardgame/internal/model"
	"cardgame/pkg/httpx"
)

func (s *Server) registerDeckRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/decks", s.createDeck)
	mux.HandleFunc("GET /api/decks", s.listDeck)
	mux.HandleFunc("GET /api/decks/{id}", s.getDeck)
	mux.HandleFunc("PUT /api/decks/{id}", s.updateDeck)
	mux.HandleFunc("DELETE /api/decks/{id}", s.deleteDeck)
	mux.HandleFunc("PATCH /api/decks/{id}/status", s.transitionDeck)
}

type createDeckRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	CardCount   int    `json:"card_count"`
	PlayerID    string `json:"player_id"`
}

func (s *Server) createDeck(w http.ResponseWriter, r *http.Request) {
	var req createDeckRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	x, err := s.svc.CreateDeck(model.Deck{Name: req.Name, Description: req.Description, CardCount: req.CardCount, PlayerID: req.PlayerID})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.Created(w, x)
}

func (s *Server) listDeck(w http.ResponseWriter, r *http.Request) {
	pp := httpx.ParsePagination(r, 20, s.maxPageSize())
	filter := model.DeckFilter{
		Status:   r.URL.Query().Get("status"),
		PlayerID: r.URL.Query().Get("player_id"),
		Keyword:  r.URL.Query().Get("keyword"),
	}
	items, total, err := s.svc.ListDecks(filter, pp.Page, pp.Size)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, httpx.PageResult{
		Items:      items,
		Pagination: httpx.Pagination{Page: pp.Page, Size: pp.Size, Total: total},
	})
}

func (s *Server) getDeck(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	x, err := s.svc.GetDeck(id)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, x)
}

type updateDeckRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	CardCount   int    `json:"card_count"`
	PlayerID    string `json:"player_id"`
}

func (s *Server) updateDeck(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var req updateDeckRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	x, err := s.svc.UpdateDeck(id, model.Deck{Name: req.Name, Description: req.Description, CardCount: req.CardCount, PlayerID: req.PlayerID})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, x)
}

func (s *Server) deleteDeck(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := s.svc.DeleteDeck(id); err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.NoContent(w)
}

type transitionDeckRequest struct {
	Status string `json:"status"`
}

func (s *Server) transitionDeck(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var req transitionDeckRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	x, err := s.svc.TransitionDeck(id, req.Status)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, x)
}
