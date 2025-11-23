package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		clientIp := r.RemoteAddr
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, `
			<!DOCTYPE html>
			<html>
			<head>
				<title>Your IP address</title>
			</head>
			<body>
				<h1>Your ip address is %s</h1>
			</body>
			</html>`, strings.Split(clientIp, ":")[0])
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
