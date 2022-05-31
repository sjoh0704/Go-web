//테스트 코드 작성하기 
package main

import (
	"net/http"
	"webstudy/web3/myapp"
)

func main(){

	http.ListenAndServe(":3000", myapp.NewHttpHandler())

}