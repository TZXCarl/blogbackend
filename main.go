package main

import (
	"blogbackend/app/application"
	"blogbackend/app/db"
	"net/http"
)

type handler struct {
	next http.Handler
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	h.next.ServeHTTP(w, r)
}

func main() {

	//初始化数据库连接
	db.Init()

	var routeHandler http.Handler = &handler{
		next: application.NewRouter(),
	}

	// http.HandleFunc("/test", )
	http.ListenAndServe(":9000", routeHandler)
}
