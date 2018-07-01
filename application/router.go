package application

import (
	// "encoding/json"
	"fmt"
	// "io/ioutil"
	"net/http"
	// "runtime"
	// "strconv"
	"io"
	"os"
	"regexp"

	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"file/port"
	"file/utils"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func filePath() {

}

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("upload")
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("file")
	if err != nil {
		utils.HandleHTTPError(w, err)
	}
	defer file.Close()
	path, _ := uuid.NewV4()
	f, err := os.OpenFile("/data/upload_files/"+path.String(), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		utils.HandleHTTPError(w, err)
	}
	defer f.Close()
	io.Copy(f, file)
	fmt.Println(path.String(), handler.Filename, '', utils.GetTimestampString())
	insertResult := port.InsertFile(path.String(), handler.Filename, '', utils.GetTimestampString())
	if insertResult != true {
		utils.HandleServerError(w, err);
	}
	fmt.Fprintln(w, "upload success!")
}

//var staticHandler http.Handler

//func init() {
//	staticHandler = http.StripPrefix("/data/upload_fiels/", http.FileServer(http.Dir("radio")))
//}

// 静态文件处理
//func StaticServer(w http.ResponseWriter, req *http.Request) {
//	fmt.Println("path:" + req.URL.Path)
//	staticHandler.ServeHTTP(w, req)
//}

var uploadPath string = "/data/upload_files"
var rest regexp = /\*/
func NewRouter() *mux.Router {
	fs := http.FileServer(http.Dir(uploadPath))
	r := mux.NewRouter()
	r.HandleFunc("/upload", upload).Methods(http.MethodPost)
	//	r.HandleFunc("/files", StaticServer)
	r.Handle(`/files/{rest}`, http.StripPrefix("/files/", fs))
	r.HandleFunc("/", index).Methods(http.MethodGet)
	return r
}
