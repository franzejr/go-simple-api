package routes

import (
  "github.com/julienschmidt/httprouter"
  "./../controllers"
)

func () Routes() {
  router := httprouter.New()
  router.GET("/", controllers.Index)
  router.GET("/hello/:name", controllers.Hello)
  return router
}
