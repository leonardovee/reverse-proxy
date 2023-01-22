package main

import (
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

func main() {
  url, err := url.Parse("http://127.0.0.1:8081")
  if err != nil {
    log.Fatal("invalid url")
  }

  http.ListenAndServe(":8080", http.HandlerFunc(func (rw http.ResponseWriter, req *http.Request) {
    id := uuid.NewString()
    path := req.URL.Path

    log.Printf("received request at: %s with id: %s", path, id) 

    req.Host = url.Host
    req.URL.Host = url.Host
    req.URL.Scheme = url.Scheme
    req.RequestURI = ""
    req.Header.Add("RequestID", id)

    res, err := http.DefaultClient.Do(req)
    if err != nil {
      log.Fatalf("couldn't forward request at: %s with id: %s", path, id)
    }

    rw.WriteHeader(res.StatusCode)
    io.Copy(rw, res.Body)

    log.Printf("response sent at: %s with id: %s", path, id) 
  }))
}

