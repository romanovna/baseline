package handler

import (
	"crud-api/models"
	"crud-api/repository"
	"crud-api/repository/post"
	"encoding/json"
	"github.com/go-chi/chi"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type Handler struct {
	repo repository.PostRepo
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{
		repo: post.NewSQLPostRepo(db),
	}
}

// swagger:route GET /posts post fetchPosts
//
// Fetches 5 posts.
//
// Responses:
//        200: postListResponse
func (p *Handler) Fetch(w http.ResponseWriter, r *http.Request) {
	payload, _ := p.repo.Fetch(r.Context(), 5)
	respondWithJSON(w, http.StatusOK, payload)
}

// swagger:route POST /posts post createPost
//
// Creates an post.
//
// Responses:
//        200: postSingleResponse
//        500: errorResponse
func (p *Handler) Create(w http.ResponseWriter, r *http.Request) {
	respPost := models.Post{}
	json.NewDecoder(r.Body).Decode(&respPost)

	err := p.repo.Create(r.Context(), &respPost)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondWithJSON(w, http.StatusCreated, respPost)
}

// swagger:route PUT /posts/{id} post updatePost
//
// Updates an post.
//
// Responses:
//        200: post
//        500: errorResponse
func (p *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data := models.Post{ID: int64(id)}
	json.NewDecoder(r.Body).Decode(&data)
	err := p.repo.Update(r.Context(), &data)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Successfully Updated"})
}

// swagger:route GET /posts/{id} post getPostByID
//
// Get post by ID.
//
// Responses:
//        200: postSingleResponse
//        204: noContentResponse
func (p *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	payload, err := p.repo.GetByID(r.Context(), int64(id))

	if err != nil {
		respondWithError(w, http.StatusNoContent, "Content not found")
	}

	respondWithJSON(w, http.StatusOK, payload)
}

// swagger:route DELETE /posts/{id} post deletePost
//
// Delete post by ID.
//
// Responses:
//        200: post
//        204: noContentResponse
//        301: errorResponse
//        500: errorResponse

func (p *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	res, err := p.repo.GetByID(r.Context(), int64(id))
	if res == nil {
		respondWithError(w, http.StatusNotFound, "Not found")
		return
	}
	err = p.repo.Delete(r.Context(), id)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
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
