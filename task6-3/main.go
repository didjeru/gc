package main

import (
	"io"
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
	http.ListenAndServe(":3000", nil)
}
