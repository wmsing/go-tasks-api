package store

import (
	"errors"
	"sync"
	"time"

	"fmt"

	"github.com/portfolio/go-tasks-api/internal/models"
)

var ErrNotFound = errors.New("task not found")

var ErrTitleRequired = errors.New("title is required")

type Memory struct {
	mu     sync.RWMutex
	byID   map[string]models.Task
	order  []string
	nextID int
}

func NewMemory() *Memory {
	return &Memory{
		byID: make(map[string]models.Task),
	}
}

func (m *Memory) List() []models.Task {
	m.mu.RLock()
	defer m.mu.RUnlock()

	out := make([]models.Task, 0, len(m.order))
	for _, id := range m.order {
		out = append(out, m.byID[id])
	}
	return out
}

func (m *Memory) Create(title string) (models.Task, error) {
	title = trimSpace(title)
	if title == "" {
		return models.Task{}, ErrTitleRequired
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	m.nextID++
	id := fmt.Sprintf("task-%d", m.nextID)
	task := models.Task{
		ID:        id,
		Title:     title,
		Completed: false,
		CreatedAt: time.Now().UTC(),
	}
	m.byID[id] = task
	m.order = append(m.order, id)
	return task, nil
}

func (m *Memory) Complete(id string) (models.Task, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	task, ok := m.byID[id]
	if !ok {
		return models.Task{}, ErrNotFound
	}
	if !task.Completed {
		task.Completed = true
		m.byID[id] = task
	}
	return task, nil
}

func trimSpace(s string) string {
	start, end := 0, len(s)
	for start < end && (s[start] == ' ' || s[start] == '\t' || s[start] == '\n' || s[start] == '\r') {
		start++
	}
	for end > start && (s[end-1] == ' ' || s[end-1] == '\t' || s[end-1] == '\n' || s[end-1] == '\r') {
		end--
	}
	return s[start:end]
}
