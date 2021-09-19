package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const (
	baseURL = "https://opentdb.com/api.php?amount="
)

func GetJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

type TriviaResponse struct {
	Response_code int                 `json:"response_code"`
	Results       []map[string]string `json:"results"`
}

func main() {
	amount := strconv.Itoa(1)
	fetchUrl := fmt.Sprintf("%s%s", baseURL, amount)
	log.Printf("Fetching trivia questions from: %s", fetchUrl)

	triviaResponse := new(TriviaResponse)
	err := GetJson(fetchUrl, triviaResponse)
	if err != nil {
		panic(err)
	}

	log.Println(triviaResponse.Response_code)
	log.Println(triviaResponse.Results)
}
