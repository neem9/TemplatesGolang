package main

import (
  // "errors"
  "fmt"
  "io"
  "net/http"

  "github.com/gorilla/mux"
)

//getHello :handler in this program
func getHello(w http.ResponseWriter, r *http.Request) {
  fmt.Printf("got /ping request\n")
  io.WriteString(w, "pong\n")
}

func main() {
  r := mux.NewRouter()

  r.HandleFunc("/ping", getHello)

  err := http.ListenAndServe(":3333", r)
  // if errors.Is(err, http.ErrServerClosed) {
  //  fmt.Printf("server closed\n")
  // } else
  if err != nil {
    fmt.Printf("error starting server: %s\n", err)
    return
  }
}