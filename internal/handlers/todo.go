package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"tz-vg/internal/entity"
	"tz-vg/internal/repository"
)

func (h *Handler) create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &entity.Todo{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			h.error(w, r, http.StatusBadRequest, repository.ErrInvalidCredentials)
			return
		}

		u := &entity.Todo{
			Title:       req.Title,
			Description: req.Description,
			DueDate:     req.DueDate,
		}

		if err := h.services.Create(h.context, u); err != nil {
			h.error(w, r, http.StatusUnprocessableEntity, repository.ErrInternalServerErr)
			return
		}

		h.respond(w, r, http.StatusCreated, u)
	}
}
func (h *Handler) getall() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todo, err := h.services.GetAll(h.context)
		if err != nil {
			h.error(w, r, http.StatusUnprocessableEntity, repository.ErrInternalServerErr)
			return
		}

		h.respond(w, r, http.StatusOK, todo)
	}
}
func (h *Handler) getone() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var id int

		idx := r.PathValue("id")
		id, err := strconv.Atoi(idx)
		if err != nil {
			h.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		todo, err := h.services.GetOne(h.context, id)
		if err != nil {
			if errors.Is(err, repository.ErrTaskNotFound) {
				h.error(w, r, http.StatusNotFound, repository.ErrTaskNotFound)
				return
			}
			h.error(w, r, http.StatusInternalServerError, repository.ErrInternalServerErr)
			return
		}
		h.respond(w, r, http.StatusOK, todo)
	}
}

func (h *Handler) update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var id int
		req := &entity.Todo{}

		idx := r.PathValue("id")
		id, err := strconv.Atoi(idx)

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			h.error(w, r, http.StatusBadRequest, repository.ErrInvalidCredentials)
			return
		}

		todo := &entity.Todo{
			Title:       req.Title,
			Description: req.Description,
			DueDate:     req.DueDate,
		}

		if err != nil {
			h.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		if err := h.services.Update(h.context, id, todo); err != nil {
			if errors.Is(err, repository.ErrTaskNotFound) {
				h.error(w, r, http.StatusNotFound, repository.ErrTaskNotFound)
				return
			}
			h.error(w, r, http.StatusInternalServerError, repository.ErrInternalServerErr)
			return
		}
		h.respond(w, r, http.StatusOK, todo)
	}
}

func (h *Handler) delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var id int

		idx := r.PathValue("id")
		id, err := strconv.Atoi(idx)
		if err != nil {
			h.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		if err := h.services.Delete(h.context, id); err != nil {
			if errors.Is(err, repository.ErrTaskNotFound) {
				h.error(w, r, http.StatusNotFound, repository.ErrTaskNotFound)
				return
			}
			h.error(w, r, http.StatusInternalServerError, repository.ErrInternalServerErr)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	}
}
