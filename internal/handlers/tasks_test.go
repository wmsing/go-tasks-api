package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/portfolio/go-tasks-api/internal/store"
)

func newTestServer(t *testing.T) *httptest.Server {
	t.Helper()
	mem := store.NewMemory()
	h := &Tasks{Store: mem}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", h.Health)
	mux.HandleFunc("GET /tasks", h.List)
	mux.HandleFunc("POST /tasks", h.Create)
	mux.HandleFunc("PATCH /tasks/{id}/complete", h.Complete)
	return httptest.NewServer(mux)
}

func TestCreateAndCompleteTask(t *testing.T) {
	srv := newTestServer(t)
	defer srv.Close()

	res, err := http.Post(srv.URL+"/tasks", "application/json", bytes.NewBufferString(`{"title":"Demo"}`))
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusCreated {
		t.Fatalf("create status: %d", res.StatusCode)
	}

	var created map[string]any
	if err := json.NewDecoder(res.Body).Decode(&created); err != nil {
		t.Fatal(err)
	}
	id, _ := created["id"].(string)

	req, _ := http.NewRequest(http.MethodPatch, srv.URL+"/tasks/"+id+"/complete", nil)
	res2, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res2.Body.Close()
	if res2.StatusCode != http.StatusOK {
		t.Fatalf("complete status: %d", res2.StatusCode)
	}
}

func TestCreateInvalidJSON(t *testing.T) {
	srv := newTestServer(t)
	defer srv.Close()

	res, err := http.Post(srv.URL+"/tasks", "application/json", bytes.NewBufferString(`not-json`))
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("status: %d", res.StatusCode)
	}
	var body errorBody
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil || body.Error == "" {
		t.Fatalf("error body: %+v err=%v", body, err)
	}
}
