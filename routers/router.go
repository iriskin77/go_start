package routers

import (
	"github.com/gorilla/mux"
)

type Handler struct {
}

func (h *Handler) InitRouters() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/products/{key}", ProductHandler)
	router.HandleFunc("/articles/{category}/", ArticlesCategoryHandler)
	router.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)
	return router
}
