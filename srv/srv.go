package main

import (
	"log"
	"net/http"
)

func main() {
  http.ListenAndServe(":8081", http.HandlerFunc(func (rw http.ResponseWriter, req *http.Request) {
    log.Printf("received request at: %s with id: %s", req.URL.Path, req.Header.Get("RequestID"))

    rw.WriteHeader(http.StatusOK)
    rw.Write([]byte("hello world\n"))

    log.Printf("response sent at: %s with id: %s", req.URL.Path, req.Header.Get("RequestID"))
  }))
}
