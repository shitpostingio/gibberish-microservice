package main

import (
	gStructs "github.com/AlessandroPomponio/go-gibberish/structs"
	"github.com/gorilla/mux"
	healthcheck "github.com/shitpostingio/analysis-commons/health-check"
	"log"
	"net/http"
)

const (
	kbPathKey      = "GB_KNOWLEDGE_PATH"
	bindAddressKey = "GB_BIND_ADDRESS"
)

var (
	knowledgeBasePath = "knowledge.json"
	bindAddress       = "localhost:10002"
	knowledgeBase     *gStructs.GibberishData
	r                 *mux.Router
)

func main() {

	r.HandleFunc("/gibberish", handleGibberish).Methods("POST")
	r.HandleFunc("/healthy", healthcheck.ConfirmServiceHealth).Methods("GET")
	log.Println("Gibberish server powered on!")
	log.Fatal(http.ListenAndServe(bindAddress, r))

}
