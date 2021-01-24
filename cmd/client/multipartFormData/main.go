package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
)

// HTTP1.0のHTMLで一番複雑な送信処理がmultipart/form-dataっぽい
// curl -F "name=Michael Jackson" -F "thumbnail=@photo.jpg" http://localhost:8888
func main() {
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)
	writer.WriteField("name", "Wichael Jackson")
	part := make(textproto.MIMEHeader)
	part.Set("Context-Type", "image/jpeg")
	part.Set("Context-Disposition", `form-data; name="thumbnail"; filename="photo.jpg"`)
	fWriter, err := writer.CreatePart(part)
	if err != nil {
		panic(err)
	}
	// fWriter, err := writer.CreateFormFile("thumbnail", "photo.jpg")
	// if err != nil {
	// 	panic(err)
	// }
	rFile, err := os.Open("photo.jpg")
	defer rFile.Close()
	io.Copy(fWriter, rFile)
	writer.Close() // bufに一気に書き込む

	resp, err := http.Post("http://localhost:8888", writer.FormDataContentType(), &buf)
	if err != nil {
		panic(err)
	}
	log.Println("Status: ", resp.Status)

}
