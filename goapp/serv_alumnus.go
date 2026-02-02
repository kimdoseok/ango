package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/xuri/excelize/v2"
)

type (
	AlumnusService struct {
		repo *AlumnusRepository
	}
)

var (
	err error
)

func NewAlumnusService(r *AlumnusRepository) *AlumnusService {
	return &AlumnusService{
		repo: r,
	}
}

func (s *AlumnusService) List(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", "Title", "Body")
}

func (s *AlumnusService) Get(w http.ResponseWriter, r *http.Request) {
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
	rec, err := s.repo.Get(int32(id))
	err = json.NewEncoder(w).Encode(rec)
	if err != nil {
		http.Error(w, "JSON Encoding error", http.StatusInternalServerError)
		return
	}
}

func (s *AlumnusService) XList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	recs, err := s.repo.List([]string{""}, 0)
	err = json.NewEncoder(w).Encode(recs)
	if err != nil {
		http.Error(w, "JSON Encoding error", http.StatusInternalServerError)
		return
	}
}

func (s *AlumnusService) Count(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	repo := NewAlumnusRepository(db)
	recs := repo.Count([]string{""})
	err = json.NewEncoder(w).Encode(recs)
	if err != nil {
		http.Error(w, "JSON Encoding error", http.StatusInternalServerError)
		return
	}
}

func (s *AlumnusService) Delete(w http.ResponseWriter, r *http.Request) {
	s.repo.Db.AutoMigrate(Alumnus{})
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
	rec, err := s.repo.Get(int32(id))
	err = json.NewEncoder(w).Encode(rec)
	if err != nil {
		http.Error(w, "JSON Encoding error", http.StatusInternalServerError)
		return
	}
}

func (s *AlumnusService) Import(w http.ResponseWriter, r *http.Request) {
	xls, err := excelize.OpenFile("Book1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	s.repo.Import(xls)
	http.Redirect(w, r, "/alumnus", http.StatusSeeOther)
}

func (s *AlumnusService) Export(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "/alumnus", http.StatusSeeOther)
}
