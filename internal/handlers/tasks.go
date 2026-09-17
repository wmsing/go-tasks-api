package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/portfolio/go-tasks-api/internal/models"
	"github.com/portfolio/go-tasks-api/internal/store"
)

type TaskStore interface {
	List() []models.Task
	Create(title string) (models.Task, error)
	Complete(id string) (models.Task, error)
}

type Tasks struct {
	Store TaskStore
}

func (h *Tasks) List(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, h.Store.List())
}

func (h *Tasks) Create(w http.ResponseWriter, r *http.Request) {
	var req models.CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	task, err := h.Store.Create(req.Title)
	if err != nil {
		if errors.Is(err, store.ErrTitleRequired) {
			writeError(w, http.StatusBadRequest, err.Error())
			return
		}
		writeError(w, http.StatusInternalServerError, "could not create task")
		return
	}
	writeJSON(w, http.StatusCreated, task)
}

func (h *Tasks) Complete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		writeError(w, http.StatusBadRequest, "invalid task id")
		return
	}

	task, err := h.Store.Complete(id)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			writeError(w, http.StatusNotFound, "task not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "could not complete task")
		return
	}
	writeJSON(w, http.StatusOK, task)
}

func (h *Tasks) Health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}
