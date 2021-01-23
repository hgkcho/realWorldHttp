package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// curl -G --data "query=hello world" http://localhost:8888
func main() {
	values := url.Values{
		"query": {"hello world"},
	}
	resp, err := http.Get("http://localhost:8888" + "?" + values.Encode())
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Print(string(body))
}
