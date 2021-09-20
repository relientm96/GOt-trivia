package main

import (
	"log"

	"github.com/relientm96/GOt-trivia/pkg/formatter"
	"github.com/relientm96/GOt-trivia/pkg/trivia"
)

func main() {
	triviaService := trivia.NewService()

	query := trivia.TriviaQuery{
		Amount:     1,
		Difficulty: "Hard",
		TriviaType: "multiple",
	}
	fetchUrl := triviaService.ConstructQuery(query)

	triviaResponse, err := triviaService.FetchTrivias(fetchUrl)
	if err != nil {
		log.Fatal(err)
	}

	formatterService := formatter.NewService()
	formatterService.PrintTrivias(*triviaResponse)
}
