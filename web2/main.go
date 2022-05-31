// json으로 요청 보내고, 응답 받기
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct{
	FirstName string `json:"first_name"` // json으로 어노테이션 => decode하고, marshal할때 어노테이션의 값과 동일하게 해준다. 
	LastName string	`json:"last_name"`
	Email string	`json:"email"`
	CreatedAt time.Time	`json:"created_at"`
}

func (u *User) ServeHTTP(w http.ResponseWriter, r *http.Request){

	// 1. 요청 처리하기 
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user) // request body에 있는 내용을 user에 디코딩하겠다. 
	
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	// 2. 비즈니스 로직  
	user.CreatedAt = time.Now()

	// 3. 응답하기  
	data, _ := json.Marshal(user) // user의 내용을 Json 형태로 인코딩한다. 
	w.Header().Add("content-type", "application/json") // reponse header에 content-type: application/json으로 명시하면 응답이 json 포멧을 보인다. 이를 명시하지만 text plain으로 처리 
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}

func main(){

	mux := http.NewServeMux()

	mux.Handle("/user", &User{})
	
	
	http.ListenAndServe(":3000", mux)

}