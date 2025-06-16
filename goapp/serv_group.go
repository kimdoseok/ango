package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type (
	GroupService struct {
		repo *GroupRepository
	}
)

func NewGroupService(r *GroupRepository) *GroupService {
	s.repo.Db.AutoMigrate(Group{})
	return &GroupService{
		repo: r,
	}
}

func (s *GroupService) List(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", "Title", "Body")
}

func (s *GroupService) Get(w http.ResponseWriter, r *http.Request) {
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

func (s *GroupService) XList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	recs, err := s.repo.List([]string{""}, 0)
	err = json.NewEncoder(w).Encode(recs)
	if err != nil {
		http.Error(w, "JSON Encoding error", http.StatusInternalServerError)
		return
	}
}

func (s *GroupService) Count(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := NewDatabase("mysql")
	repo := NewGroupRepository(db)
	recs := repo.Count([]string{""})
	err = json.NewEncoder(w).Encode(recs)
	if err != nil {
		http.Error(w, "JSON Encoding error", http.StatusInternalServerError)
		return
	}
}

func (s *GroupService) Delete(w http.ResponseWriter, r *http.Request) {
	s.repo.Db.AutoMigrate(Group{})
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
