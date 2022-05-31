// go web 기본 
package main

import (
	"fmt"
	"net/http"
)

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "hello foo!")
}

type barHandler struct{}

func (b *barHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	name := r.URL.Query().Get("name") // http://localhost:3000/bar?name=sseung 이런식으로 접속 
	if name == ""{
		name = "world"
	}
	fmt.Fprintf(w, "Hello %s!", name)

}

func main(){

	// router를 만들어서 등록하기 
	mux := http.NewServeMux()

	// 어떤 request가 왔을 때 handler를 등록하는 함수
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){  
		fmt.Fprint(w, "hello world")
	})	

	mux.Handle("/foo", &fooHandler{}) // ServerHTTP 메서드를 갖는 인터페이스

	mux.Handle("/bar", &barHandler{})

	http.ListenAndServe(":3000", mux)
}