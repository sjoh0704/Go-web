// rest api + test 
// GET /users/10

package main

import (
	"net/http"
	"webstudy/web4/myapp"
)

func main(){

	mux := myapp.NewHandler()
	http.ListenAndServe(":3000", mux)
}