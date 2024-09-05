package handlers

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"tz-vg/internal/service"
)

type Handler struct {
	services *service.Service
	logger   *slog.Logger
	context  context.Context
}

type DeleteMessage struct {
	Message string `json:"message"`
}

func NewHandler(services *service.Service, logger *slog.Logger, context context.Context) *Handler {
	return &Handler{
		services: services,
		logger:   logger,
		context:  context,
	}
}

func (h *Handler) InitRouters() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("POST /tasks", h.create())
	router.HandleFunc("GET /tasks", h.getall())
	router.HandleFunc("GET /tasks/{id}", h.getone())
	router.HandleFunc("PUT /tasks/{id}", h.update())
	router.HandleFunc("DELETE /tasks/{id}", h.delete())

	return router
}

func (h *Handler) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	h.respond(w, r, code, map[string]string{"Ошибка": err.Error()})
}

func (h *Handler) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}
