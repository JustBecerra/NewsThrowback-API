package repositories

import (
	"database/sql"

	"NewsThrowback-API/models"
)

type ArticleRepository struct {
	db *sql.DB
}

func NewArticleRepository(db *sql.DB) *ArticleRepository {
	return &ArticleRepository{db: db}
}

func (r *ArticleRepository) GetAll() ([]models.Article, error) {
	rows, err := r.db.Query(`
		SELECT id, title, date, newspaper, content, source_url
		FROM articles
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []models.Article
	for rows.Next() {
		var a models.Article
		if err := rows.Scan(&a.ID, &a.Title, &a.Date, &a.Newspaper, &a.Content, &a.SourceURL); err != nil {
			return nil, err
		}
		articles = append(articles, a)
	}

	return articles, nil
}

func (r *ArticleRepository) GetByID(id int) (*models.Article, error) {
	var a models.Article
	err := r.db.QueryRow(`
		SELECT id, title, date, newspaper, content, source_url
		FROM articles WHERE id = $1
	`, id).Scan(&a.ID, &a.Title, &a.Date, &a.Newspaper, &a.Content, &a.SourceURL)
	if err != nil {
		return nil, err
	}
	return &a, nil
}
