package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

import "owerwatch-heroes/service"

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("api//heros", service.AllHeroes)
	myRouter.HandleFunc("api/hers/{id}", service.OneHero)
	myRouter.HandleFunc("api/heros/{id}/abilities", service.HeroAbilities)
	myRouter.HandleFunc("api/ability", service.AllAbilities)
	myRouter.HandleFunc("api/ability/{id}", service.OneAbility)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	handleRequests()
}
