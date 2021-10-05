package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func fileURLToBase64(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	mimeType := http.DetectContentType(bytes)

	var base64Encode string

	base64Encode += "data:" + mimeType + ";base64," + toBase64(bytes)

	return base64Encode, nil
}

func fileToBase64(filePath string) (string, error) {
	bytes, err := ioutil.ReadFile("./exfile/file.jpg")
	if err != nil {
		return "", nil
	}

	mimeType := http.DetectContentType(bytes)

	var base64Encoding string

	base64Encoding += "data:" + mimeType + ";base64," + toBase64(bytes)

	return base64Encoding, nil

}

func main() {

	u, err := fileURLToBase64("https://www.techhub.in.th/wp-content/uploads/2021/05/577280151-1.jpg")
	if err != nil {
		log.Fatal(err)
	}

	f, err := fileToBase64("./exfile/file.jpg")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(f)
	fmt.Println(u)

}
