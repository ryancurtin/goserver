package main

import (
  "log"
  "net/http"

  "github.com/ant0ine/go-json-rest/rest"
  "github.com/goserver/controllers"
  // "gopkg.in/mgo.v2"
  // "gopkg.in/mgo.v2/bson"
)

// func getSession() *mgo.Session {
//   // connect to DB
//   s, err := mgo.Dial('mongodb://localhost')

//   // error check on connection
//   if err != nil {
//     panic(err)
//   }

//   return s
// }

func main() {
  api := rest.NewApi()
  api.Use(rest.DefaultDevStack...)

  ec := controllers.NewExerciseController()

  router, err := rest.MakeRouter(
    rest.Get("/exercises", ec.GetExercises),
  )
  if err != nil {
    log.Fatal(err)
  }
  api.SetApp(router)
  log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}
