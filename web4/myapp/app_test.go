package myapp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T){
	assert := assert.New(t)

	// 실제 서버는 아니지만, 테스트 서버가 나온다. 
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	res, err := http.Get(ts.URL)
 	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("index", string(data))

}

// GET /users
func TestUsersHandler(t *testing.T){
	assert := assert.New(t)

	ts:= httptest.NewServer(NewHandler())
	defer ts.Close()

	res, err := http.Get(ts.URL + "/users") // 왜 패스가 됐지? => 하위에 있는 애가 정의되어 있지 않으면 상위에 있는애가 호출
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	data, _ := ioutil.ReadAll(res.Body) // 여기서 data는 바이트이므로 string으로 바꾸어주어야 함
	assert.Contains(string(data), "Get user info") // Get user info가 포함되어 있는 데이터인지 체크 =>
}

// GET /users/id
func TestUsersGetInfoHandler(t *testing.T){
	assert := assert.New(t)

	ts:= httptest.NewServer(NewHandler())
	defer ts.Close()

	res, err := http.Get(ts.URL + "/users/10")	
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	data, _ := ioutil.ReadAll(res.Body)
	assert.Contains(string(data), "No user id: 10")
}

// POST /users

func TestCreateUser(t *testing.T){
	assert := assert.New(t)
	ts := httptest.NewServer(NewHandler()) // mux가 인자로 들어감 
	defer ts.Close()

	res, err := http.Post(ts.URL + "/users" , "application/json",
	strings.NewReader(`{"first_name": "seung", "last_name": "oh", "email": "example.com"}`))

	assert.NoError(err)
	assert.Equal(http.StatusCreated, res.StatusCode)
	
	user := new(User)
	err = json.NewDecoder(res.Body).Decode(user)
	assert.NoError(err)
	assert.Equal("seung", user.FirstName)
	assert.Equal("oh", user.LastName)
	assert.Equal("example.com", user.Email)
	fmt.Println("user_id", user.ID)

}

func TestDeleteUser(t *testing.T){
	assert := assert.New(t)
	ts := httptest.NewServer(NewHandler()) 
	defer ts.Close()
	req, _ := http.NewRequest("DELETE", ts.URL + "/users/1", nil)
	res, err := http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	// 잠시 어떤 데이터가 찍히는지 확인 
	data, _ := ioutil.ReadAll(res.Body)
	log.Print(string(data))
	// getUserInfo가 호출되었음
	
	// assert.Contains(string(data), "No User id: 1")
}