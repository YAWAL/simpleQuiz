package domain

import (
	"encoding/json"
	"io/ioutil"

	"github.com/YAWAL/simpleQuiz/models"
)

// InitQuiz perform reading file with questions and initiates quiz
func InitQuiz(qaFilePath string) ( *models.Players, error) {
	qa, err := ioutil.ReadFile(qaFilePath)
	if err != nil {
		return  nil, err
	}
	var QandA models.Questions
	err = json.Unmarshal(qa, &QandA)
	if err != nil {
		return  nil, err
	}
	var quiz models.Quiz
	quiz.Questions = QandA
	var players models.Players
	return &players, nil
}
