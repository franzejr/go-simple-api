package main

import (
  "github.com/julienschmidt/httprouter"
  "net/http"
  "log"
)

func main() {
  router := httprouter.New()
  router.GET("/", Index)
  router.GET("/hello/:name", Hello)

  log.Fatal(http.ListenAndServe(":8080", router))
}
