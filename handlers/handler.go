package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/YAWAL/simpleQuiz/domain"
	"github.com/YAWAL/simpleQuiz/models"
)

// HomePage is a handler function which renders home page
func HomePage(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("Quiz home page"))

}

// loadQuestions performs reading questions from file and write them to Go's structure
func loadQuestions(path string) (data []byte, err error) {
	data, err = ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if !json.Valid(data) {
		return nil, err
	}
	return data, nil
}

// RenderQuestions is a handler which renders json representation of questions in quiz w/o correct answers
func RenderQuestions(qFilePath string) func(w http.ResponseWriter, req *http.Request) {
	questions, err := loadQuestions(qFilePath)
	if err != nil {
		return func(w http.ResponseWriter, req *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
	return func(w http.ResponseWriter, req *http.Request) {
		w.Write(questions)
	}
}

// SavePlayer is a handler functions which saves player
func SavePlayer(players models.Players) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		decoder := json.NewDecoder(req.Body)
		var body models.SavePlayerPayload
		err := decoder.Decode(&body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		player := domain.StorePlayer(body.Name)
		fmt.Println(player)
		players.Players = append(players.Players, player)
		fmt.Println(players)
		w.WriteHeader(http.StatusOK)
	}
}

// Players is a handler functions which renders info about all players
func Players(players *models.Players) func(w http.ResponseWriter, req *http.Request) {
	data, err := json.Marshal(&players)
	if err != nil {
		return func(w http.ResponseWriter, req *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
	return func(w http.ResponseWriter, req *http.Request) {
		w.Write(data)
	}
}

// PlayersAnswers is a handler functions which renders info about  player's answers
func PlayersAnswers(players models.Players) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		name := req.URL.Query().Get(domain.Name)
		if name == "" {
			w.WriteHeader(http.StatusBadRequest)
		}
		for _, p := range players.Players {
			if p.Name == name {
				w.WriteHeader(http.StatusOK)
				data, err := json.Marshal(p.PlAnswers)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
				}
				w.Write(data)
			}
		}
	}
}
