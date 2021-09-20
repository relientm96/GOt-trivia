package trivia

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const (
	baseURL = "https://opentdb.com/api.php?"
)

// service provides the read-write implementation of TriviaService.
type TriviaService struct{}

// NewService returns a configured TriviaService implementation.
func NewService() *TriviaService {
	return &TriviaService{}
}

func (s *TriviaService) ConstructQuery(query TriviaQuery) string {
	urlWithQuery := baseURL

	if query.Amount != 0 {
		urlWithQuery = fmt.Sprintf("%samount=%s", urlWithQuery, strconv.Itoa(query.Amount))
	}

	if query.Difficulty != "" {
		urlWithQuery = fmt.Sprintf("%s&difficulty=%s", urlWithQuery, strings.ToLower(query.Difficulty))
	}

	if query.TriviaType != "" {
		urlWithQuery = fmt.Sprintf("%s&type=%s", urlWithQuery, strings.ToLower(query.TriviaType))
	}

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
