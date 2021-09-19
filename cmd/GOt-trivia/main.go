package main

import (
	"fmt"
	"log"
	"strconv"
)

const (
	baseURL = "https://opentdb.com/api.php?amount="
)

type TriviaResponse struct {
	Response_code int                 `json:"response_code"`
	Results       []map[string]string `json:"results"`
}

func main() {
	amount := strconv.Itoa(1)
	fetchUrl := fmt.Sprintf("%s%s", baseURL, amount)
	log.Printf("Fetching trivia questions from: %s", fetchUrl)

	triviaResponse := new(TriviaResponse)
	err := getJson(fetchUrl, triviaResponse)
	if err != nil {
		panic(err)
	}

	log.Println(triviaResponse.Response_code)
	log.Println(triviaResponse.Results)
}
