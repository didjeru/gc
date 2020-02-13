package main

import (
	"io"
	"log"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")

	res.Header().Set("Content-Type", "text/html")
	io.WriteString(res, `
		<doctype html>
		<html>
			<head>
				<title>Hello `+name+`!</title>
			</head>
			<body>
				Hello `+name+`!
			</body>
		</html>
	`)
}

func main() {
	http.HandleFunc("/hello", hello)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}
