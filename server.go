package main

import(
	"http"
	"io"
)

func HelloServer(c *http.Conn, req *http.Request) {
	io.WriteString(c, "Hello world!");
}

func main() {
	http.Handle("/", http.HandlerFunc(HelloServer))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.String())
	}
}
