package myapp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFooHandler(t *testing.T){
	assert := assert.New(t)
	// res, req를 받는다. 
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/foo", nil) // 제일 끝은 바디 

	FooHandler(res, req)

	// assert로 표현하면 다음과 같다. 
	assert.Equal(http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello foo!", string(data))
}

func TestFooHandler2(t *testing.T){
	assert := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/?name=sseung", nil) // 쿼리 

	FooHandler(res, req)

	// 상태 코드 확인 
	assert.Equal(http.StatusOK, res.Code)
	
	// 응답 값 확인 
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello sseung!", string(data))
}

func TestUserHandler(t *testing.T){
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/user", 
		strings.NewReader(`{"first_name": "sseung", "last_name": "aa", "email": "example@naver.com"}`))

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)
	// user := new(User)
	// user.ServeHTTP(res, req)
	
	assert.Equal(http.StatusOK, res.Code)
	// data := ioutil.ReadAll(res.Body)
	// assert.Equal("df", string(data))

	user := new(User)
	err := json.NewDecoder(res.Body).Decode(user)
	assert.Nil(err) // 에러가 없어야 한다. 
	assert.Equal("sseung", user.FirstName)
	assert.Equal("aa", user.LastName)
	assert.Equal("example@naver.com", user.Email)
	


}