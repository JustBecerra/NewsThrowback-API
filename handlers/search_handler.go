package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"NewsThrowback-API/models"
)

type SearchHandler struct {
	baseURL string
}

func NewSearchHandler(baseURL string) *SearchHandler {
	return &SearchHandler{baseURL: baseURL}
}

func (h *SearchHandler) Search(w http.ResponseWriter, r *http.Request) {
	keyword := r.URL.Query().Get("keyword")
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")

	if keyword == "" {
		http.Error(w, "keyword query parameter is required", http.StatusBadRequest)
		return
	}

	searchURL := fmt.Sprintf(
		"%s/collections/chronicling-america/?q=%s&dates=%s/%s&fo=json&c=5",
		h.baseURL,
		url.QueryEscape(keyword),
		from,
		to,
	)

	resp, err := http.Get(searchURL)
	if err != nil {
		http.Error(w, "failed to reach Chronicling America API", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "failed to read API response", http.StatusInternalServerError)
		return
	}
	log.Printf("Chronicling America raw response: %s", string(body))

	var result models.SearchResponse
	if err := json.Unmarshal(body, &result); err != nil {
		http.Error(w, "failed to parse API response", http.StatusInternalServerError)
		return
	}

	if len(result.Results) > 5 {
		result.Results = result.Results[:5]
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
