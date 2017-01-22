package controllers

import (
  "github.com/goserver/models"
  "github.com/ant0ine/go-json-rest/rest"
)

type (
  ExerciseController struct {}
)

func NewExerciseController() *ExerciseController {
  return &ExerciseController{}
}

func (ec ExerciseController) GetExercises(w rest.ResponseWriter, r *rest.Request) {
  ecs := models.Exercise{
    Title: "Overhead Press",
    Id: 1,
  }

  w.WriteJson(ecs)
}
