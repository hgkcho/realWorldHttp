package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// x-www-form-urlencoded形式のPOSTメソッドの送信
// curl -d test=value http://localhost:8888
func main() {
	values := url.Values{
		"test": {"value"},
	}
	resp, err := http.PostForm("http://localhost:8888", values)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	log.Println(resp.Status)
}
