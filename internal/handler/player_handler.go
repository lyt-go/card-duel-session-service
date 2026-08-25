package handler

import (
	"net/http"

	"cardgame/internal/model"
	"cardgame/pkg/httpx"
)

func (s *Server) registerPlayerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/players", s.createPlayer)
	mux.HandleFunc("GET /api/players", s.listPlayer)
	mux.HandleFunc("GET /api/players/{id}", s.getPlayer)
	mux.HandleFunc("PUT /api/players/{id}", s.updatePlayer)
	mux.HandleFunc("DELETE /api/players/{id}", s.deletePlayer)
	mux.HandleFunc("PATCH /api/players/{id}/status", s.transitionPlayer)
}

type createPlayerRequest struct {
	Nickname string `json:"nickname"`
	Level    int    `json:"level"`
	Score    int    `json:"score"`
}

func (s *Server) createPlayer(w http.ResponseWriter, r *http.Request) {
	var req createPlayerRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	x, err := s.svc.CreatePlayer(model.Player{Nickname: req.Nickname, Level: req.Level, Score: req.Score})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.Created(w, x)
}

func (s *Server) listPlayer(w http.ResponseWriter, r *http.Request) {
	pp := httpx.ParsePagination(r, 20, s.maxPageSize())
	filter := model.PlayerFilter{
		Status:  r.URL.Query().Get("status"),
		Keyword: r.URL.Query().Get("keyword"),
	}
	items, total, err := s.svc.ListPlayers(filter, pp.Page, pp.Size)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, httpx.PageResult{
		Items:      items,
		Pagination: httpx.Pagination{Page: pp.Page, Size: pp.Size, Total: total},
	})
}

func (s *Server) getPlayer(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	x, err := s.svc.GetPlayer(id)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, x)
}

type updatePlayerRequest struct {
	Nickname string `json:"nickname"`
	Level    int    `json:"level"`
	Score    int    `json:"score"`
}

func (s *Server) updatePlayer(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var req updatePlayerRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	x, err := s.svc.UpdatePlayer(id, model.Player{Nickname: req.Nickname, Level: req.Level, Score: req.Score})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, x)
}

func (s *Server) deletePlayer(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := s.svc.DeletePlayer(id); err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.NoContent(w)
}

type transitionPlayerRequest struct {
	Status string `json:"status"`
}

func (s *Server) transitionPlayer(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var req transitionPlayerRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	x, err := s.svc.TransitionPlayer(id, req.Status)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, x)
}
