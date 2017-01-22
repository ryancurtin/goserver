package controllers

import (
  "log"

  "github.com/goserver/models"
  "github.com/ant0ine/go-json-rest/rest"
  "gopkg.in/mgo.v2"
)

type (
  ExerciseController struct {
    session *mgo.Session
  }
)

// Should use ENV variable here
const database string = "development"

func NewExerciseController(s *mgo.Session) *ExerciseController {
  return &ExerciseController{s}
}

func (ec ExerciseController) GetExercises(w rest.ResponseWriter, r *rest.Request) {
  var results []models.Exercise

  // Querying from specified database on exercises collection
  err := ec.session.DB(database).C("exercises").Find(nil).All(&results)
  if err != nil {
    log.Fatal(err)
  }

  w.WriteJson(results)
}
