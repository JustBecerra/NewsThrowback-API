package models

type Article struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Date      string `json:"date"`
	Newspaper string `json:"newspaper"`
	Content   string `json:"content"`
	SourceURL string `json:"source_url"`
}
