package main

import (
  "net/http"
  "log"
  "./routes"
)

func main() {
  router := routes.Routes()

  log.Fatal(http.ListenAndServe(":8080", router))
}
