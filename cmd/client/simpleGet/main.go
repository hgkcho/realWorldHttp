package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

// curl http://localhost:8888
func main() {
	resp, err := http.Get("http://localhost:8888")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	log.Printf("%T\n", body)
	log.Println(body)
	log.Printf("%T\n", string(body))
	log.Println(string(body))

}
