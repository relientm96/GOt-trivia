# GOt-trivia

Simple golang program game for trivia questions and answers.

Questions and answers taken from the [Open Trivia Database](https://opentdb.com/).

## Usage

* Amount: `integer` (number of questions)
* Difficulty: `easy/medium/hard` (difficulty level)
* Trivia Type: `boolean/multiple` (true/false ormultiple choice)

```bash
Usage of ./main:
  -amount string (optional)
        number of trivia questions (default "1")
  -difficulty string (optional)
        easy/medium/hard
  -type string (optional)
        boolean/multiple
```

## Play the Game!

```bash
# Build binary
go build cmd/GOt-trivia/main

./main -amount <amount> -difficulty <difficulty> -type <type>
```