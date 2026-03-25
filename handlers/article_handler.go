package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"NewsThrowback-API/repositories"
)

type ArticleHandler struct {
	repo *repositories.ArticleRepository
}

func NewArticleHandler(repo *repositories.ArticleRepository) *ArticleHandler {
	return &ArticleHandler{repo: repo}
}

func (h *ArticleHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	articles, err := h.repo.GetAll()
	if err != nil {
		http.Error(w, "failed to fetch articles", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

func (h *ArticleHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	// Extract the ID from the URL path: /articles/{id}
	parts := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		http.Error(w, "invalid article ID", http.StatusBadRequest)
		return
	}

	article, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, "article not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(article)
}
