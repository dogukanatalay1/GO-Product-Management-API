package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	helpers "product-api/helpers"
	. "product-api/models"
	"strconv"
	"time"
	"github.com/gorilla/mux"
)

var productStore = make(map[string]Product)

// Make method allows you to create and initialize an object of type slice, map, or chan.

var id int = 0

// HTTP Post - /api/products
func PostProductHandler(w http.ResponseWriter, r *http.Request) {
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	helpers.CheckError(err)

	product.CreatedOn = time.Now()
	id++
	product.ID = id
	key := strconv.Itoa(id)
	productStore[key] = product

	data, err := json.Marshal(product)
	helpers.CheckError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

// HTTP Get - /api/products
func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	var products []Product
	for _, product := range productStore {
		products = append(products, product)
	}

	data, err := json.Marshal(products)
	helpers.CheckError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// HTTP Get - /api/products/{id}
func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	var product Product
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])
	for _, prd := range productStore {
		if prd.ID == key {
			product = prd
		}
	}

	data, err := json.Marshal(product)
	helpers.CheckError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// HTTP Put - /api/products
func PutProductHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	key := vars["id"]

	var prodUpdate Product
	err = json.NewDecoder(r.Body).Decode(&prodUpdate)
	helpers.CheckError(err)

	if _, ok := productStore[key]; ok {
		prodUpdate.ID, _ = strconv.Atoi(key)
		prodUpdate.ChangedOn = time.Now()
		delete(productStore, key)
		productStore[key] = prodUpdate
	} else {
		log.Printf("Deger bulunamadi : %s", key)
		w.WriteHeader(http.StatusNoContent)
	}
	w.WriteHeader(http.StatusOK)
}

// HTTP Delete - /api/products/{id}
func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  key := vars["id"]
  if _, ok := productStore[key]; ok {
    delete(productStore, key)
  } else {
    log.Printf("Deger bulunamadi : %s", key)
  }
  w.WriteHeader(http.StatusOK)
}
