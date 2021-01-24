package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Postメソッドで任意のBodyを送信
// curl -T main.go -H "Content-Type: text/plain" http://localhost:8888
func main() {
	f, err := os.Open("main.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// stirngs.NewReader()やbytes.Bufferはよく使うみたい
	// reader := strings.NewReader("asvasvd")

	// io.Readerのまま渡す
	resp, err := http.Post("http://localhost:8888", "text/plain", f)
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
