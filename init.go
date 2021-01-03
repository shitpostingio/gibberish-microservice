package main

import (
	"github.com/AlessandroPomponio/go-gibberish/persistence"
	"github.com/gorilla/mux"
	"log"
	"os"
)

func init() {

	setEnvVars()

	var err error
	knowledgeBase, err = persistence.LoadKnowledgeBase(knowledgeBasePath)
	if err != nil {
		log.Fatal("Unable to start gibberish server: ", err)
	}

	r = mux.NewRouter()

}

func setEnvVars() {

	kb := os.Getenv(kbPathKey)
	if kb != "" {
		knowledgeBasePath = kb
	}

	add := os.Getenv(bindAddressKey)
	if add != "" {
		bindAddress = add
	}

}
