package models

// QuestionsAnswers is a structure which represents info about one question in a Quiz
type QuestionsAnswers struct {
	Question      string   `json:"question"`
	Answers       *Answers `json:"answers"`
	CorrectAnswer string   `json:"correctAnswer"`
}

// RenderQuestions is a structure which contains info about one question and no correct answer
type RenderQuestions struct {
	Questions string   `json:"question"`
	Answers   *Answers `json:"answers"`
}

// Questions is a structure which holds questions and corresponding correct answers
type Questions struct {
	Questions []QuestionsAnswers
}

// Answers is a structure which holds all answers related to question
type Answers struct {
	A string `json:"a"`
	B string `json:"b"`
	C string `json:"c"`
	D string `json:"d"`
}

// Quiz is a structure which holds questions fora quiz
type Quiz struct {
	Questions `json:"questions"`
}

// Player holds info about player
type Player struct {
	Name      string            `json:"name"`
	PlAnswers map[string]string `json:"playerAnswers"`
}

// Player holds info about all players in a quiz
type Players struct {
	Players []Player `json:"players"`
}

// SavePlayerPayload is a structure for storing player's name for saving
type SavePlayerPayload struct {
	Name string `json:"name"`
}
