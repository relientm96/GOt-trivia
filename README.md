# GOt-trivia

Simple golang program that polls [Open Trivia Database](https://opentdb.com/) for multiple choice type trivia questions and answers.

## Usage

* Amount: `integer` (number of questions)
* Difficulty: `easy/medium/hard` (difficulty level)
* Trivia Type: `boolean/multiple` (true/false ormultiple choice)

```bash
Usage of ./main:
  -amount string
        number of trivia questions (default "1")
  -difficulty string
        easy/medium/hard
  -type string
        boolean/multiple
```

## Running

```bash
# Build binary
go build cmd/GOt-trivia/main

./main -amount <amount> -difficulty <difficulty> -type <type>
```