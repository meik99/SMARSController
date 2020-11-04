package main

import (
	"log"
	"net/http"
)

func alarm(writer http.ResponseWriter, request *http.Request) {

}

func handleRequests() {
	http.HandleFunc("/apps/coffeetogo/api/v1/alarm", alarm)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
