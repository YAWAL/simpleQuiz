package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Question is a structure which represents info about one question in a quiz
type Question struct {
	Question      string   `json:"question,omitempty"`
	Answers       *Answers `json:"answers,omitempty"`
	CorrectAnswer string   `json:"correctAnswer,omitempty"`
}

// Answers is a structure which holds all answers related to question
type Answers struct {
	A string `json:"a,omitempty"`
	B string `json:"b,omitempty"`
	C string `json:"c,omitempty"`
	D string `json:"d,omitempty"`
}

var quiz []Question

func quizHandler(w http.ResponseWriter, req *http.Request) {
	enableCORS(&w)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quiz)
}

func startQuiz() {
	quiz = collectQuestions()

	router := mux.NewRouter()
	router.HandleFunc("/quiz", quizHandler).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8888", router))
}

// enable cross-origin resource sharing
func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func collectQuestions() []Question {
	return []Question{
		{Question: "What is the capital of the Malta?",
			Answers:       &Answers{A: "Milan", B: "Valletta", C: "London", D: "Paris"},
			CorrectAnswer: "b"},
		{Question: "Who was the first chess champion?",
			Answers:       &Answers{A: "Michail Tall", B: "Johannes Zukertort", C: "Wilhelm Steinitz", D: "Volodymyr Klitchko"},
			CorrectAnswer: "c"},
		{Question: "When Swedish became Sweden's official language?",
			Answers:       &Answers{A: "1 June 2008", B: "1 June 2009", C: "1 July 2007", D: "1 July 2009"},
			CorrectAnswer: "d"},
		{Question: "What is not an OOP principle?",
			Answers:       &Answers{A: "inheritance", B: "encapsulation", C: "polyglot", D: "abstraction"},
			CorrectAnswer: "c"},
	}
}

func main() {
	startQuiz()
}
