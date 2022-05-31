package myapp

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func indexHandler(w http.ResponseWriter, r *http.Request){

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "index")
	
}

func usersHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Get user info of all users")

}


func usersGetInfoHandler(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r) // vars안에 id값이 있음 
	fmt.Fprintf(w, "Get user info by user id: %s", vars["id"])

}


func NewHandler() http.Handler{
	mux := mux.NewRouter() //고릴라 Mux에 있는 메서드 
	
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/users", usersHandler)
	mux.HandleFunc("/users/{id:[0-9]+}", usersGetInfoHandler) // 고릴라 mux를 사용하면 자동으로 파싱



	return mux
}