package main

import (
	"io"
	"net/http"
)

func run() {
	http.HandleFunc("/",blockchainGetHandle)
	http.ListenAndServe("localhost:8000",nil)
}

func blockchainGetHandle(w http.ResponseWriter,r *http.Request) {
	io.WriteString(w,"<h1>hello , i am webserver v1</h1>")
}

func main() {
	run()
}