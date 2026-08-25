package handler

import (
	"net/http"

	"cardgame/internal/model"
	"cardgame/pkg/httpx"
)

func (s *Server) registerCollectionRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/collections", s.createCollection)
	mux.HandleFunc("GET /api/collections", s.listCollection)
	mux.HandleFunc("GET /api/collections/{id}", s.getCollection)
	mux.HandleFunc("DELETE /api/collections/{id}", s.deleteCollection)
}

type createCollectionRequest struct {
	PlayerID string `json:"player_id"`
	CardID   string `json:"card_id"`
	Count    int    `json:"count"`
}

func (s *Server) createCollection(w http.ResponseWriter, r *http.Request) {
	var req createCollectionRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	x, err := s.svc.CreateCollection(model.Collection{PlayerID: req.PlayerID, CardID: req.CardID, Count: req.Count})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.Created(w, x)
}

func (s *Server) listCollection(w http.ResponseWriter, r *http.Request) {
	pp := httpx.ParsePagination(r, 20, s.maxPageSize())
	filter := model.CollectionFilter{
		PlayerID: r.URL.Query().Get("player_id"),
		CardID:   r.URL.Query().Get("card_id"),
	}
	items, total, err := s.svc.ListCollections(filter, pp.Page, pp.Size)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, httpx.PageResult{
		Items:      items,
		Pagination: httpx.Pagination{Page: pp.Page, Size: pp.Size, Total: total},
	})
}

func (s *Server) getCollection(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	x, err := s.svc.GetCollection(id)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, x)
}

func (s *Server) deleteCollection(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := s.svc.DeleteCollection(id); err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.NoContent(w)
}
