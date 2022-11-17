package main

import (
	// "errors"
	"fmt"
	"io"
	"net/http"
)


//getHello :handler in this program
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /ping request\n")
	io.WriteString(w, "pong\n")
}

func main() {

	http.HandleFunc("/ping", getHello)

	err := http.ListenAndServe(":3333", nil)

	// if errors.Is(err, http.ErrServerClosed) {
	// 	fmt.Printf("server closed\n")
	// } else
	if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		return
	}
}
