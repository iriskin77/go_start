package main

import (
	"log"
	"net/http"
	"time"
	//"router/routers.go"
)

func main() {
	// r := mux.NewRouter()
	// r.HandleFunc("/products/{key}", ProductHandler)
	// r.HandleFunc("/articles/{category}/", ArticlesCategoryHandler)
	// r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)
	http.Handle("/", r)

	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:9090",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	//http.ListenAndServe("127.0.0.1:9090", r)

	log.Fatal(srv.ListenAndServe())

}
