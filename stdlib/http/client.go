package http

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func DecodeBodyUsingScanner(body *io.ReadCloser) map[string]interface{} {
	scanner := bufio.NewScanner(*body)
	var jsonString string
	for scanner.Scan() {
		jsonString = jsonString + scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	var jsonContent map[string]interface{}
	if err := json.Unmarshal([]byte(jsonString), &jsonContent); err != nil {
		panic(err)
	}
	return jsonContent
}

func DecodeBodyUsingJsonDecoder(body *io.ReadCloser) map[string]interface{} {
	jsonDecoder := json.NewDecoder(*body)
	var jsonContent map[string]interface{}
	err := jsonDecoder.Decode(&jsonContent)
	if err != nil {
		panic(err)
	}
	return jsonContent
}

func Client(){
	resp, err := http.Get("https://dummyjson.com/products")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// fmt.Println(decodeBodyUsingScanner(&resp.Body))
	fmt.Println(DecodeBodyUsingJsonDecoder(&resp.Body))
}