package main

import (
	"encoding/json"
	"net/http"
	"log"
	"fmt"
)

func main(){
	http.HandleFunc("/users", getUsers)

	fmt.Println("Api running on 8081")

	log.Fatal(http.ListenAndServe(":8081", nil))
}

type User struct {
	ID int
	Name string
}

func getUsers(w http.ResponseWriter, r *http.Request){
	 w.Header().Set("Content=Type", "application/json")
	 json.NewEncoder(w).Encode([]User{{
		ID: 1,
		Name: "Lucas",
	 }})
}