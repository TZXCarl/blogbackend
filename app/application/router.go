package application

import (
	// "encoding/json"
	"fmt"
	// "io/ioutil"
	"net/http"
	// "runtime"
	// "strconv"
	"os"
	"io"

	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {  
    w.Write([]byte("hello world"))  
}

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("enter")

	r.ParseMultipartForm(32 << 20) 
	file, handler, err := r.FormFile("uploadfile") 
	fmt.Println(handler.Filename)
	 
    if err != nil {
        fmt.Println(err)  
        return  
    }  
    defer file.Close()  
    f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)  
    if err != nil {  
        fmt.Println(err)  
        return  
    }  
    defer f.Close()  
    io.Copy(f, file)  
    fmt.Fprintln(w, "upload ok!")  
}

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/upload", upload).Methods(http.MethodGet)
	r.HandleFunc("/", index).Methods(http.MethodGet)	
	return r
}
