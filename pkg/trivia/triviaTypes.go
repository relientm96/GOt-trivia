package trivia

type TriviaResult struct {
	Category          string   `json:"category"`
	Type              string   `json:"type"`
	Difficulty        string   `json:"difficulty"`
	Question          string   `json:"question"`
	Correct_answer    string   `json:"correct_answer"`
	Incorrect_answers []string `json:"incorrect_answers"`
}

type TriviaResponse struct {
	Response_code int            `json:"response_code"`
	Results       []TriviaResult `json:"results"`
}

type TriviaQuery struct {
	Amount     string
	Difficulty string
	TriviaType string
}
