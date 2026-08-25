package handler

import (
	"net/http"

	"cardgame/internal/model"
	"cardgame/pkg/httpx"
)

func (s *Server) registerCardRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/cards", s.createCard)
	mux.HandleFunc("GET /api/cards", s.listCard)
	mux.HandleFunc("GET /api/cards/{id}", s.getCard)
	mux.HandleFunc("PUT /api/cards/{id}", s.updateCard)
	mux.HandleFunc("DELETE /api/cards/{id}", s.deleteCard)
}

type createCardRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Cost        int    `json:"cost"`
	Attack      int    `json:"attack"`
	Health      int    `json:"health"`
	Rarity      string `json:"rarity"`
	Type        string `json:"type"`
}

func (s *Server) createCard(w http.ResponseWriter, r *http.Request) {
	var req createCardRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	x, err := s.svc.CreateCard(model.Card{
		Name:        req.Name,
		Description: req.Description,
		Cost:        req.Cost,
		Attack:      req.Attack,
		Health:      req.Health,
		Rarity:      req.Rarity,
		Type:        req.Type,
	})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.Created(w, x)
}

func (s *Server) listCard(w http.ResponseWriter, r *http.Request) {
	pp := httpx.ParsePagination(r, 20, s.maxPageSize())
	filter := model.CardFilter{
		Rarity:  r.URL.Query().Get("rarity"),
		Type:    r.URL.Query().Get("type"),
		Keyword: r.URL.Query().Get("keyword"),
	}
	items, total, err := s.svc.ListCards(filter, pp.Page, pp.Size)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, httpx.PageResult{
		Items:      items,
		Pagination: httpx.Pagination{Page: pp.Page, Size: pp.Size, Total: total},
	})
}

func (s *Server) getCard(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	x, err := s.svc.GetCard(id)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, x)
}

type updateCardRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Cost        int    `json:"cost"`
	Attack      int    `json:"attack"`
	Health      int    `json:"health"`
	Rarity      string `json:"rarity"`
	Type        string `json:"type"`
}

func (s *Server) updateCard(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var req updateCardRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	x, err := s.svc.UpdateCard(id, model.Card{
		Name:        req.Name,
		Description: req.Description,
		Cost:        req.Cost,
		Attack:      req.Attack,
		Health:      req.Health,
		Rarity:      req.Rarity,
		Type:        req.Type,
	})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, x)
}

func (s *Server) deleteCard(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := s.svc.DeleteCard(id); err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.NoContent(w)
}
