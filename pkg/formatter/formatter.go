package formatter

import (
	"fmt"

	"github.com/relientm96/GOt-trivia/pkg/trivia"
)

// service provides the read-write implementation of FormatterService.
type FormatterService struct{}

// NewService returns a configured FormatterService implementation.
func NewService() *FormatterService {
	return &FormatterService{}
}

func (s *FormatterService) PrintTrivias(triviaResponse trivia.TriviaResponse) {
	for _, response := range triviaResponse.Results {
		fmt.Println(response.Category)
		fmt.Println(response.Type)
		fmt.Println(response.Difficulty)
		fmt.Println(response.Question)
		fmt.Println(response.Correct_answer)
		fmt.Println(response.Incorrect_answers)
	}
}
