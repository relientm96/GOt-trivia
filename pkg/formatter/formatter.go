package formatter

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/relientm96/GOt-trivia/pkg/trivia"
)

// service provides the read-write implementation of FormatterService.
type FormatterService struct{}

// NewService returns a configured FormatterService implementation.
func NewService() *FormatterService {
	return &FormatterService{}
}

func (s *FormatterService) PrintTrivias(triviaResponse trivia.TriviaResponse) {
	// Common info
	fmt.Println()
	fmt.Println("Category:   ", triviaResponse.Results[0].Category)
	fmt.Println("Type:       ", triviaResponse.Results[0].Type)
	fmt.Println("Difficulty: ", triviaResponse.Results[0].Difficulty)
	fmt.Println()

	for i, result := range triviaResponse.Results {
		answers, correctIndex := makeAnswers(result)
		printQuestion(result, answers, i)
		err := handleQuestion(result, answers, correctIndex)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("------------------------\n\n")
	}
}

func handleQuestion(triviaResult trivia.TriviaResult, answers []string, corectIndex int) error {
	fmt.Print("Enter selection: ")
	var input string
	fmt.Scanln(&input)

	selectedByUser, err := strconv.Atoi(input)
	if err != nil {
		return err
	}

	if corectIndex == selectedByUser {
		fmt.Println("Good Job!")
	} else {
		fmt.Printf("Wrong Answer! The answer is %d), %s \n", corectIndex, answers[corectIndex])
	}

	return nil
}

func makeAnswers(triviaResult trivia.TriviaResult) ([]string, int) {
	answers := make([]string, 0)
	answers = append(answers, triviaResult.Incorrect_answers...)
	answers = append(answers, triviaResult.Correct_answer)
	return answers, len(answers) - 1
}

func printQuestion(triviaResult trivia.TriviaResult, answers []string, questionNumber int) {
	fmt.Printf("====== Question %d ======\n", questionNumber)
	fmt.Println(applyFormat(triviaResult.Question))

	fmt.Println()

	for i, answer := range answers {
		fmt.Printf("%d) %s\n", i, answer)
	}

	fmt.Println()
}

func applyFormat(text string) string {
	text = formatQuotes(text)
	text = formatApostrophe(text)
	return text
}

func formatQuotes(text string) string {
	var quotesRegex = regexp.MustCompile(`&quot;`)
	return quotesRegex.ReplaceAllString(text, "\"")
}

func formatApostrophe(text string) string {
	var apostropheRegex = regexp.MustCompile(`&#039;`)
	return apostropheRegex.ReplaceAllString(text, "'")
}
