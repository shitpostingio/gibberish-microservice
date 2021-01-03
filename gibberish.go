package main

import (
	"encoding/json"
	"github.com/AlessandroPomponio/go-gibberish/gibberish"
	"gitlab.com/shitposting/analysis-api/services/structs"
	"io/ioutil"
	"log"
	"net/http"
)

func handleGibberish(w http.ResponseWriter, r *http.Request) {

	data, err := ioutil.ReadAll(r.Body)
	defer func() {
		err = r.Body.Close()
		if err != nil {
			log.Println("Error while closing request body:", err)
		}
	}()

	isGibberish := gibberish.IsGibberish(string(data), knowledgeBase)
	err = json.NewEncoder(w).Encode(&structs.GibberishResponse{IsGibberish: isGibberish})
	if err != nil {
		log.Println("Unable to send response: ", err)
	}

}
