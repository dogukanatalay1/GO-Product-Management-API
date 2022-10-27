package handlers

import (
	"encoding/json"
	"net/http"
	helpers "product-api/helpers"

)

func MainHandler(w http.ResponseWriter, r *http.Request) {

  data , err := json.Marshal("Welcome to homepage!")
  helpers.CheckError(err)

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)
  w.Write(data)
}

