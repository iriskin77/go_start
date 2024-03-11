package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func ArticlesCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ProductHandler"))
}

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ArticleHandler"))
}
