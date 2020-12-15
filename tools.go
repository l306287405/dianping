package dianping

import (
	"io/ioutil"
	"math/rand"
	"net/http"
)

func PostForm(url string,R *ReqParams)([]byte,error){
	r,err:=http.PostForm(url,R.Values)
	if err!=nil{
		return nil,err
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	return body,err
}

func RandStr(len int) string {
	bytes := make([]byte, len)
	r := &rand.Rand{}
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}