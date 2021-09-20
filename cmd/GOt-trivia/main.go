package main

import (
	"flag"
	"log"

	"github.com/relientm96/GOt-trivia/pkg/formatter"
	"github.com/relientm96/GOt-trivia/pkg/trivia"
)

func main() {
	amount := flag.String("amount", "1", "number of trivia questions")
	difficulty := flag.String("difficulty", "", "easy/medium/hard")
	triviaType := flag.String("type", "", "boolean/multiple")

	flag.Parse()

	query := trivia.TriviaQuery{
		Amount:     *amount,
		Difficulty: *difficulty,
		TriviaType: *triviaType,
	}

	triviaService := trivia.NewService()
	fetchUrl := triviaService.ConstructQuery(query)

	triviaResponse, err := triviaService.FetchTrivias(fetchUrl)
	if err != nil {
		log.Fatal(err)
	}

	formatterService := formatter.NewService()
	formatterService.PrintTrivias(*triviaResponse)
}
