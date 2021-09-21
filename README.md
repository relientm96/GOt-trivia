# GOt-trivia

Simple golang program game for trivia questions and answers.

Questions and answers taken from the [Open Trivia Database](https://opentdb.com/).

## Install

Either download the compiled binaries in the release, or manually compile them with the following make commands:

```bash
# Compile for Windows
make compile-windows

# Compile for Linux
make compile-linux

# Compile for MacOS
make compile-darwin
```

## Play the Game

```bash
# Play with 10 questions
./main -amount 10

# Run the game with custom args (refer below for usage)
./main -amount <amount> -difficulty <difficulty> -type <type>
```

## Usage

* Amount: `integer` (number of questions)
* Difficulty: `easy/medium/hard` (difficulty level)
* Trivia Type: `boolean/multiple` (true/false or multiple choice)

```bash
Usage of ./main:
  -amount string (optional)
        number of trivia questions (default "1")
  -difficulty string (optional)
        easy/medium/hard
  -type string (optional)
        boolean/multiple
```
