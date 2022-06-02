// Rest API 예제
// TDD
// user 저장은 map을 통해서 함, DB 필요 X
// 특정 유저 정보 얻기: GET /users/id
// 유저 생성: POST /users


package main

import (
	"net/http"
	"webstudy/web4/myapp"
)

func main(){

	mux := myapp.NewHandler()
	http.ListenAndServe(":3000", mux)
}

