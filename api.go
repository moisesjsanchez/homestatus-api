package main

import(
	"time"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

func postMeeting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	return
}

func postSleeping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	return
}

func postOOO(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	return
}

func postFree(w http.ResponseWriter, r *http.Request) {	
	w.Header().Set("Content-Type", "application/json")
	return
}

func main() {
	r := mux.NetRouter()

	r.HandleFunc("/api/meeting", postMeeting).Methods("POST")	
	r.HandleFunc("/api/sleeping", postSleeping).Methods("POST")
	r.HandleFunc("/api/ooo", postOOO).Methods("POST")
	r.HandleFunc("/api/free", postFree).Methods("POST")

	http.ListenAndServer(":4566",r)
}