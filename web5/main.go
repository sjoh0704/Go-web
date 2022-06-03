// 데코레이터
// data => Encrypt => Zip => send
// recv => Unzip => Decrypt => data

package main

import (
	"github.com/tuckersGo/goweb/web9/lzw"	
	"github.com/tuckersGo/goweb/web9/cipher"
	"fmt"
)


type Component interface{
	Operator(string)
}

var sendData string

// 보내기 
type SendComponent struct {}

func (self *SendComponent) Operator(data string){
	sendData = data
}

// 압축하는 데코레이터 
type ZipComponent struct {
	com Component
}

func (self *ZipComponent) Operator(data string){
	dat, err := lzw.Write([]byte(data))
	if err != nil {
		panic(err)
	}

	self.com.Operator(string(dat))
}

// 암호화하는 데코레이터
type EncryptComponet struct {
	key string
	com Component
}

func (self *EncryptComponet) Operator(data string){
	dat, err := cipher.Encrypt([]byte(data), self.key)
	if err != nil{
		panic(err)
	}

	self.com.Operator(string(dat))
}

// 받기
var readData string

type ReadComponent struct {}

func (self *ReadComponent) Operator(data string){
	readData = data
}

// 압축해제하는 데코레이터 
type UnzipComponent struct {
	com Component
}

func (self *UnzipComponent) Operator(data string){
	dat, err := lzw.Read([]byte(data))
	if err != nil {
		panic(err)
	}

	self.com.Operator(string(dat))
}

// 복호화하는 데코레이터
type DecryptComponet struct {
	key string
	com Component
}

func (self *DecryptComponet) Operator(data string){
	dat, err := cipher.Decrypt([]byte(data), self.key)
	if err != nil{
		panic(err)
	}
	self.com.Operator(string(dat))
}


func main(){
	// data => Encrypt => Zip => send
	// recv => Unzip => Decrypt => data
	
	// 보내기 
	key := "abcde"
	sender := &EncryptComponet{key: key,
		com: &ZipComponent{
			com: &SendComponent{}}}

	sender.Operator("hello world") 
	fmt.Println(sendData)

	// 받기 
	reader := &UnzipComponent{
			com: &DecryptComponet{
				key:key,
				com: &ReadComponent{}}}
	
	reader.Operator(sendData)
	fmt.Println(readData)

}