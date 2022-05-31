package myapp

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"

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

func TestUsersGetInfoHandler(t *testing.T){
	assert := assert.New(t)

	ts:= httptest.NewServer(NewHandler())
	defer ts.Close()

	res, err := http.Get(ts.URL + "/users/10")	
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	data, _ := ioutil.ReadAll(res.Body)
	assert.Contains(string(data), "user id: 10")

}