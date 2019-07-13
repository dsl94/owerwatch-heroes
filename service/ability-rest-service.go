package service

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

import myData "owerwatch-heroes/data"

func AllAbilities(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	response, err := http.Get("https://overwatch-api.net/api/v1/ability/")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
		var response myData.AbilitiesResponse
		json.Unmarshal(data, &response);
		json.NewEncoder(w).Encode(response)
	}
}

func OneAbility(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	w.Header().Set("Content-Type", "application/json")
	response, err := http.Get("https://overwatch-api.net/api/v1/ability/" + id)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
		var response myData.AbilityResponse
		json.Unmarshal(data, &response);
		json.NewEncoder(w).Encode(response)
	}
}