package main

import (
	"fmt"
	"net/http"
)

func myWeb(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(w, "Welcome to littie cookie web")
	if err != nil {
		return
	}
}

func main() {

	http.HandleFunc("/", myWeb)

	fmt.Println("init little cookie web server at http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("服务器开启错误: ", err)
	}
}
