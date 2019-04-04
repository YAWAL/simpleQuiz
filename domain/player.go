package domain

import (
	"fmt"

	"github.com/YAWAL/simpleQuiz/models"
)

type players models.Players

type player models.Player

// AddPlayer perform adding another player to players slice
func (p *players) AddPlayer(player models.Player) {
	p.Players = append(p.Players, player)
}

// PlayerAnswers returns map with answers for particular player
func (p *player) PlayerAnswers() map[string]string {
	return p.PlAnswers
}

// PlayersCorrectAnswers calculetes quantity of correct answers for particular player
func (p *player) PlayersCorrectAnswers() (counter int) {
	ca := correctAnswers()
	pa := p.PlayerAnswers()
	for k, v := range ca {
		if pa[k] == v {
			counter++
		}
	}
	return counter
}

// StorePlayer creates new player with given name
func StorePlayer(name string) models.Player {
	player := models.Player{Name: name}
	return player
}

// resultMessage creates string with information how better is particular player among all players
func (l *player) resultMessage(pl []player) string {
	result := l.PlayersCorrectAnswers()
	plCount := len(pl)
	var adv, percentage int
	for _, p := range pl {
		if p.PlayersCorrectAnswers() < result {
			adv++
		}
	}
	if adv == 0 {
		percentage = worstPlayer
	} else {
		percentage = allPlayers/plCount - adv
	}
	return fmt.Sprintf("You where better then %d of all quizer", percentage)
}
