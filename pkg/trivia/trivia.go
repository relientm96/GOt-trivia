package trivia

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	baseURL = "https://opentdb.com/api.php?"
)

// service provides the read-write implementation of triviaService.
type TriviaService struct{}

// NewService returns a configured triviaService implementation.
func NewService() *TriviaService {
	return &TriviaService{}
}

func (s *TriviaService) ConstructQuery(amount int) string {
	amountInStr := strconv.Itoa(amount)

	urlWithQuery := fmt.Sprintf("%samount=%s", baseURL, amountInStr)

	fmt.Printf("Fetching trivia questions from: %s\n", urlWithQuery)
	return urlWithQuery
}

func (s *TriviaService) FetchTrivias(fetchUrl string) (*TriviaResponse, error) {
	response, err := http.Get(fetchUrl)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	triviaObjectResponse := new(TriviaResponse)
	err = json.Unmarshal(body, triviaObjectResponse)
	if err != nil {
		return nil, err
	}

	return triviaObjectResponse, nil
}
