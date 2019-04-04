package router

import (
	"net/http"

	"github.com/YAWAL/simpleQuiz/handlers"
	"github.com/YAWAL/simpleQuiz/models"
	"github.com/gorilla/mux"
)

// InitRouter builds router for quiz
func InitRouter(qFilePath string, players *models.Players) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", handlers.HomePage).Methods(http.MethodGet)
	router.HandleFunc("/questions", handlers.RenderQuestions(qFilePath)).Methods(http.MethodGet)
	router.HandleFunc("/players", handlers.Players(players)).Methods(http.MethodGet)
	router.HandleFunc("/players", handlers.SavePlayer(*players)).Methods(http.MethodPost)
	router.HandleFunc("/players/{name}", handlers.SavePlayer(*players)).Methods(http.MethodPost)
	router.HandleFunc("/players/answers/{name}", handlers.PlayersAnswers(*players)).Methods(http.MethodGet)
	return router
}
