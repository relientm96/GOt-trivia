package game

import (
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"time"

	"github.com/relientm96/GOt-trivia/pkg/trivia"
)

// service provides the read-write implementation of GameService.
type GameService struct{}

// NewService returns a configured GameService implementation.
func NewService() *GameService {
	return &GameService{}
}

func (s *GameService) Start(triviaResponse trivia.TriviaResponse) {
	// Print Common info
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
		fmt.Printf("------------------------------------------------\n\n")
	}
}

func handleQuestion(triviaResult trivia.TriviaResult, answers []string, corectIndex int) error {
	var input string
	var selectedByUser int
	for {
		fmt.Print("Enter selection: ")
		fmt.Scanln(&input)

		parsedInput, err := strconv.Atoi(input)
		if err != nil {
			return err
		}

		if parsedInput < len(answers) && parsedInput >= 0 {
			selectedByUser = parsedInput
			// User has successfully selected a legitimate selection
			break
		}
	}

	if corectIndex == selectedByUser {
		fmt.Println("Good Job!")
	} else {
		fmt.Printf("Wrong Answer! The answer is %d) - %s \n", corectIndex, answers[corectIndex])
	}

	return nil
}

func makeAnswers(triviaResult trivia.TriviaResult) ([]string, int) {
	if triviaResult.Type == "boolean" {
		answers := make([]string, 2)
		answers[0] = "true"
		answers[1] = "false"

		var correctAnswerIndex int
		if triviaResult.Correct_answer == "True" {
			correctAnswerIndex = 0
		} else {
			correctAnswerIndex = 1
		}
		return answers, correctAnswerIndex
	}

	// Handle selections for multiple choice type
	max := len(triviaResult.Incorrect_answers) + 1 // number of incorrect ans + correct ans
	rand.Seed(time.Now().UnixNano())
	correctAnswerIndex := rand.Intn(max)

	// Create list of answers for selection
	answers := make([]string, max)
	incorrectAnsCounter := 0
	for i := range answers {
		if i == correctAnswerIndex {
			answers[i] = triviaResult.Correct_answer
			continue
		}
		answers[i] = triviaResult.Incorrect_answers[incorrectAnsCounter]
		incorrectAnsCounter++
	}

	return answers, correctAnswerIndex
}

func printQuestion(triviaResult trivia.TriviaResult, answers []string, questionNumber int) {
	fmt.Printf("====== Question %d ======\n", questionNumber)
	fmt.Println(applyFormat(triviaResult.Question))

	fmt.Println()

	for i, answer := range answers {
		fmt.Printf("%d) %s\n", i, applyFormat(answer))
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
