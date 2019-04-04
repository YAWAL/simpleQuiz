package main

import (
	"log"
	"net/http"

	"github.com/YAWAL/simpleQuiz/config"
	"github.com/YAWAL/simpleQuiz/domain"
	"github.com/YAWAL/simpleQuiz/router"
)

const configFile = "config.json"

func main() {
	//read config
	cnf, err := config.Load(configFile)
	if err != nil {
		log.Printf("can not load config from file: %s , err: %s", configFile, err)
	}
	log.Println("config loaded")

	// from requirement: Database - Just in-memory , so no database
	// init in-memory storage for quiz
	players, err := domain.InitQuiz(cnf.QuestionAnswersFilePath)
	if err != nil {
		log.Printf("can not init quiz: %s", err)
	}
	log.Println("quiz initialized")

	// init router
	r := router.InitRouter(cnf.QuestionFilePath, players)
	log.Printf("router initialized on port: %s", cnf.Port)

	log.Fatal(http.ListenAndServe(cnf.Port, r))

}
