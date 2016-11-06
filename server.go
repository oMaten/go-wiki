package main

import (
  "fmt"
  "net/http"
)

func requestHanler(res http.ResponseWriter, req *http.Request) {
  fmt.Fprintf(res, "Hi there, I love %s!", req.URL.Path[1:])
}

func main() {
  http.HandleFunc("/", requestHanler)
  http.ListenAndServe(":8080", nil)
}
