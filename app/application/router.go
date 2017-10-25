package application

import (
	"blogbackend/app/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"strconv"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
)

var version = "v1"

func HandleServerError(w http.ResponseWriter, e error) {
	_, f, l, _ := runtime.Caller(1)
	errData := struct {
		Code   int
		Detail string
		Meta   map[string]interface{}
	}{
		Code:   500,
		Detail: e.Error(),
		Meta: map[string]interface{}{
			"file": f,
			"line": l,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)
	json, err := json.Marshal(errData)
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(w, string(json))
}

func testGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`hello world`))
}

func defaultFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`default`))
}
func createNote(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	result, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	note := &models.Note{}

	err = json.Unmarshal([]byte(result), &note)
	if err != nil {
		fmt.Println(note)
	}
	note.Id = bson.NewObjectId()
	now := time.Now()
	note.CreatedTime = strconv.FormatInt(now.UTC().UnixNano(), 10)
	note.UpdateTime = strconv.FormatInt(now.UTC().UnixNano(), 10)
	fmt.Println(note)

	res := InsertNote(*note)
	if res {
		json, _ := json.Marshal(note)
		wc := 0
		for wc < len(json) {
			n, err := w.Write(json)
			if err != nil {
				panic(err)
			}
			wc += n
		}
	}
}

func getNote(w http.ResponseWriter, r *http.Request) {
	// findNote()
}

func login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	result, _ := ioutil.ReadAll(r.Body)
	user := models.User{}
	json.Unmarshal([]byte(result), &user)
	username := user.Username
	password := user.Password
	user = FindUserByNameAndPwd(username, password)

	fmt.Println(user)
	json, err := json.Marshal(&struct {
		User models.User `json:"data"`
	}{
		User: user,
	})
	if err != nil {
		panic(err)
		// w.Write([]byte(""))
	}
	wc := 0
	for wc < len(json) {
		n, err := w.Write(json)
		if err != nil {
			panic(err)
		}
		wc += n
	}
}

func NewRouter() *mux.Router {

	r := mux.NewRouter()
	// r = r.PathPrefix("/" + version).Subrouter()

	//默认
	r.HandleFunc("/", defaultFunc).Methods(http.MethodGet)

	//用户登录
	r.HandleFunc("/login", login).Methods(http.MethodPost)

	//保存文档
	r.HandleFunc("/createNote", createNote).Methods(http.MethodPost)

	//hold库存
	// r.HandleFunc("/skus/hold", HoldPost).Methods(http.MethodPost)
	return r
}
