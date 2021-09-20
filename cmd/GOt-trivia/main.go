package main

import (
	"fmt"
	"log"

	"github.com/relientm96/GOt-trivia/pkg/trivia"
)

func main() {
	triviaService := trivia.NewService()

	fetchUrl := triviaService.ConstructQuery(1)
	triviaResponse, err := triviaService.FetchTrivias(fetchUrl)
	if err != nil {
		log.Fatal(err)
	}

	for _, response := range triviaResponse.Results {
		fmt.Println(response.Category)
		fmt.Println(response.Type)
		fmt.Println(response.Difficulty)
		fmt.Println(response.Question)
		fmt.Println(response.Correct_answer)
		fmt.Println(response.Incorrect_answers)
	}
}
