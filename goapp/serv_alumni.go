package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type (
	AlumniService struct {
		repo *AlumniRepository
	}
)

var (
	sitesallowed = []string{"", "localhost"}
)

func NewAlumniService(r *AlumniRepository) *AlumniService {
	return &AlumniService{
		repo: r,
	}
}

func (s *AlumniService) Get(w http.ResponseWriter, r *http.Request) {
	s.repo.Db.AutoMigrate(Alumni{})
	w.Header().Set("Content-Type", "application/json")
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	rec, err := s.repo.Get(id)
	err = json.NewEncoder(w).Encode(rec)
	if err != nil {
		http.Error(w, "JSON Encoding error", http.StatusInternalServerError)
		return
	}
}

func (s *AlumniService) List(w http.ResponseWriter, r *http.Request) {
	s.repo.Db.AutoMigrate(Alumni{})
	w.Header().Set("Content-Type", "application/json")
	recs, err := s.repo.List([]string{""}, 0)
	err = json.NewEncoder(w).Encode(recs)
	if err != nil {
		http.Error(w, "JSON Encoding error", http.StatusInternalServerError)
		return
	}
}

func (s *AlumniService) Count(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := NewDatabase("mysql")
	repo := NewAlumniRepository(db)
	recs := repo.Count([]string{""})
	err = json.NewEncoder(w).Encode(recs)
	if err != nil {
		http.Error(w, "JSON Encoding error", http.StatusInternalServerError)
		return
	}
}

func (s *AlumniService) Delete(w http.ResponseWriter, r *http.Request) {
	s.repo.Db.AutoMigrate(Alumni{})
	w.Header().Set("Content-Type", "application/json")
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	rec, err := s.repo.Get(id)
	err = json.NewEncoder(w).Encode(rec)
	if err != nil {
		http.Error(w, "JSON Encoding error", http.StatusInternalServerError)
		return
	}
}
