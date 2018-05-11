package main

import (
	"blogbackend/app/application"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type handler struct {
	next http.Handler
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	w.Header().Add("Access-Control-Allow-Headers", "x-requested-with")
	// w.Header.
	// w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	// w.Header().Set("Content-Type", "application/json;charset=utf-8")
	h.next.ServeHTTP(w, r)
}

func main() {
	var routeHandler http.Handler = &handler{
		next: application.NewRouter(),
	}

	http.ListenAndServe(":9010", routeHandler)
}
