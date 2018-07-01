package main

import (
	"net/http"
	"fmt"
	"io"
	"os"
	"github.com/satori/go.uuid"
	"file/utils"
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
			utils.HandleHTTPError(w, err);
		}
		defer file.Close()
		fileId, _ := uuid.NewV4()
		f, err := os.OpenFile("/data/upload_files/"+fileId.String(), os.O_WRONLY|os.O_CREATE, 0666)
		fmt.Println(fileId.String())
		if err != nil {
			utils.HandleServerError(w, err);
		}
		defer f.Close()
		io.Copy(f, file)
		fmt.Fprintln(w, "upload success!")
	}
}

func main() {
	//var routeHandler http.Handler = &handler{
	//		next: application.NewRouter(),
	//}

	fs := http.FileServer(http.Dir("/data/upload_files"))
	// http.Handle("/", routeHandler)
	http.Handle("/", http.StripPrefix("/", fs))
	http.HandleFunc("/upload", upload)
	http.ListenAndServe(":9010", nil)
}
