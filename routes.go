package main

import (
  "github.com/julienschmidt/httprouter"
)

func (*httprouter.Router)Routes(){
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	return router
}
