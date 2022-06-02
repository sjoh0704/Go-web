package myapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type User struct{
	ID int	`json:"id"`
	FirstName string `json:"first_name"`
	LastName string	`json:"last_name"`
	Email string	`json:"email"`
	CreatedAt time.Time	`json:"created_at"`
}

var userMap map[int] *User // user정보를 담을 map 생성
var lastID int

func indexHandler(w http.ResponseWriter, r *http.Request){

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "index")
	
}

func usersHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Get user info of all users")

}


func usersGetInfoHandler(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r) // vars안에 id값이 있음 
	id, err := strconv.Atoi(vars["id"])

	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return 
	}
	
	// user가 있을 때 
	if user, ok := userMap[id]; ok{
		// user를 json으로 변환
		data, _ := json.Marshal(user)
		fmt.Fprint(w, string(data))
		return
	}

	// user가 없을 때 
	w.WriteHeader(http.StatusOK)
	w.Header().Add("content-type", "application/json")
	fmt.Fprintf(w, "No user id: %d", id)

}

func createUserHandler(w http.ResponseWriter, r *http.Request){
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)

	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	
	// 유저 생성
	user.ID = lastID
	lastID += 1
	user.CreatedAt = time.Now()
	userMap[user.ID] = user

	// data를 json을 바꿔서 응답
	w.WriteHeader(http.StatusCreated)
	data, _ := json.Marshal(user) // 바이트를 리턴하므로 string으로 바꿔주자.
	fmt.Fprint(w, string(data))
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r) // vars안에 id값이 있음 
	id, err := strconv.Atoi(vars["id"])

	if err != nil{
		fmt.Fprint(w, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, ok := userMap[id]
	if !ok{ // 지울 유저가 없으면 
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "No User id: %d", id)
		return 
	}

	delete(userMap, id)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Deleted User id: %d", id)

	

}



func NewHandler() http.Handler{
	userMap = make(map[int]*User) // mapUser 초기화 \
	lastID = 1

	mux := mux.NewRouter() //고릴라 Mux에 있는 메서드 
	
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/users", usersHandler).Methods("GET") //고릴라 mux의 기능
	mux.HandleFunc("/users", createUserHandler).Methods("POST")
	mux.HandleFunc("/users/{id:[0-9]+}", usersGetInfoHandler).Methods("GET") // 고릴라 mux를 사용하면 자동으로 파싱
	mux.HandleFunc("/users/{id:[0-9]+}", DeleteUserHandler).Methods("DELETE") // 고릴라 mux를 사용하면 자동으로 파싱


	return mux
}