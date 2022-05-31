package myapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// user
type User struct{
	FirstName string `json:"first_name"`
	LastName string	`json:"last_name"`
	Email string	`json:"email"`
	CreatedAt time.Time	`json:"created_at"`
}

func (u *User) ServeHTTP(w http.ResponseWriter, r *http.Request){
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	user.CreatedAt = time.Now()

	data, _ := json.Marshal(user)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}


// BarHandler struct
type BarHandler struct{}

func  (b *BarHandler) ServeHTTP(w  http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "hello bar")
}


func FooHandler(w http.ResponseWriter, r *http.Request){
	name := r.URL.Query().Get("name")
	if name == ""{
		fmt.Fprint(w, "Hello foo!")
		return
	}
	fmt.Fprintf(w, "Hello %s!", name)
}


// mux 
func NewHttpHandler() http.Handler { // mux는 http.Handler를 인터페이스로 사용

	mux := http.NewServeMux()

	mux.HandleFunc("/foo", FooHandler)

	mux.Handle("/bar", new(BarHandler))

	mux.Handle("/user", &User{})

	return mux
}