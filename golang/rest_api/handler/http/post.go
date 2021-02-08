package handler

import (
	driver "crud-api/driver/mysql"
	"crud-api/models"
	"crud-api/repository"
	"crud-api/repository/post"
	"encoding/json"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"strconv"
)

type Post struct {
	repo repository.PostRepo
}

func NewPostHandler(db *driver.DB) *Post {
	return &Post{
		repo: post.NewSQLPostRepo(db.SQL),
	}
}

func (p *Post) Init(w http.ResponseWriter, r *http.Request) {
	_, err := p.repo.Init(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
		return
	}
	respondWithJSON(w, http.StatusCreated, map[string]string{"message": "Table Successfully Created"})
}

func (p *Post) Fetch(w http.ResponseWriter, r *http.Request) {
	payload, _ := p.repo.Fetch(r.Context(), 5)
	respondWithJSON(w, http.StatusOK, payload)
}

func (p *Post) Create(w http.ResponseWriter, r *http.Request) {
	respPost := models.Post{}
	json.NewDecoder(r.Body).Decode(&respPost)

	newID, err := p.repo.Create(r.Context(), &respPost)
	log.Println(newID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]string{"message": "Successfully Created"})
}

func (p *Post) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data := models.Post{ID: int64(id)}
	json.NewDecoder(r.Body).Decode(&data)
	payload, err := p.repo.Update(r.Context(), &data)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondWithJSON(w, http.StatusOK, payload)
}

func (p *Post) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	payload, err := p.repo.GetByID(r.Context(), int64(id))

	if err != nil {
		respondWithError(w, http.StatusNoContent, "Content not found")
	}

	respondWithJSON(w, http.StatusOK, payload)
}

func (p *Post) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	_, err := p.repo.Delete(r.Context(), int64(id))

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
		return
	}

	respondWithJSON(w, http.StatusMovedPermanently, map[string]string{"message": "Delete Successfully"})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"message": msg})
}
