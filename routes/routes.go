package routes

import (
	"net/http"

	"NewsThrowback-API/handlers"
)

func Register(mux *http.ServeMux, articleHandler *handlers.ArticleHandler) {
	mux.HandleFunc("/articles", articleHandler.GetAll)
	mux.HandleFunc("/articles/", articleHandler.GetByID)
}
