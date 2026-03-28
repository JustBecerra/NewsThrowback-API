package models

type ArticleResult struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Date        string   `json:"date"`
	City        []string `json:"location_city"`
	State       []string `json:"location_state"`
	URL         string   `json:"url"`
	Description []string `json:"description"`
}

type Pagination struct {
	Total int `json:"total"`
}

type SearchResponse struct {
	Pagination Pagination      `json:"pagination"`
	Results    []ArticleResult `json:"results"`
}
