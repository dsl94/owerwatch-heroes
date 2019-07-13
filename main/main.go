package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

import "owerwatch-heroes/service"

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/hero", service.AllHeroes)
	myRouter.HandleFunc("/hero/{id}", service.OneHero)
	myRouter.HandleFunc("/ability", service.AllAbilities)
	myRouter.HandleFunc("/ability/{id}", service.OneAbility)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	handleRequests()
}
