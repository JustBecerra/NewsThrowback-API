package routes

import (
	"net/http"

	"NewsThrowback-API/handlers"
)

func Register(mux *http.ServeMux, articleHandler *handlers.ArticleHandler, searchHandler *handlers.SearchHandler) {
	mux.HandleFunc("/articles", articleHandler.GetAll)
	mux.HandleFunc("/articles/", articleHandler.GetByID)
	mux.HandleFunc("/article", searchHandler.Search)
}
