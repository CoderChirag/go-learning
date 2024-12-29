package http

import (
	"encoding/json"
	"net/http"
)

func data(w http.ResponseWriter, _ *http.Request) {
	resp, err := http.Get("https://dummyjson.com/products")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	w.Header().Add("content-type", "application/json")
	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(DecodeBodyUsingJsonDecoder(&resp.Body))
}

func headers(w http.ResponseWriter, req *http.Request){
	headers := make(map[string]string)
	for k, v := range req.Header {
		for _, val := range v {
			headers[k] = val
		}
	}
	w.Header().Add("content-type", "application/json")
	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(headers)
}

func Server(){
	http.HandleFunc("/data", data)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8080", nil)
}