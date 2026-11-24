package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type ShortenRequest struct {
	URL       string `json:"url"`
	ExpiresIn int    `json:"expires_in_days"`
}

type ShorenResponse struct {
	ShortURL  string     `json:"short_url"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

type Server struct {
	// store   Store
	// keygen  *Keygen
	// cache   *LRU
	// baseURL string
	// clicks  chan Click
	// db      *sql.DB
}

func main() {
	s := &Server{}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /shorten", s.handleShorten)
	mux.HandleFunc("GET /{key}", s.handleRedirect)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func (s *Server) handleShorten(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var req ShortenRequest

	if err := dec.Decode(&req); err != nil {
		http.Error(w, "bad request body", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ShorenResponse{ShortURL: "http://localhost:8080.fake12"})
}

func (s *Server) handleRedirect(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.PathValue("key"))
}
