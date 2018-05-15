package main

import (
	// "blogbackend/app/application"
	"net/http"
	// "runtime"
	// "strconv"
	"fmt"
	"io"
	"os"
	//	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	// _ "github.com/go-sql-driver/mysql"
)

type handler struct {
	next http.Handler
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "x-requested-with")
	w.Header().Add("Content-Type", "application/x-msdownload")

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
	} else {
		// w.Header.
		// w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
		// w.Header().Set("Content-Type", "application/json;charset=utf-8")
		h.next.ServeHTTP(w, r)

	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "x-requested-with")
	w.Header().Add("Content-Type", "application/x-msdownload")

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("file")
		fmt.Println(handler.Filename)

		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fileId, _ := uuid.NewV4()
		f, err := os.OpenFile("/data/upload_files/"+fileId.String()+"__"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		fmt.Println(fileId.String() + "__" + handler.Filename)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
		fmt.Fprintln(w, "upload ok!")
	}
}

func main() {
	//var routeHandler http.Handler = &handler{
	//		next: application.NewRouter(),
	//}

	fs := http.FileServer(http.Dir("/data/upload_files"))
	// http.Handle("/", routeHandler)
	http.Handle("/files/", http.StripPrefix("/files/", fs))
	http.HandleFunc("/upload", upload)
	http.ListenAndServe(":9010", nil)
}
