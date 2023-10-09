package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

const (
	maxAttempts = 10
	numberRange = 100
)

func main() {
	randomNumber := generateRandomNumber(numberRange)

	fmt.Print("Try to guess the number (1-100): ")

	reader := bufio.NewReader(os.Stdin)

	for attempts := 0; attempts < maxAttempts; attempts++ {
		answer := getUserInput(reader)
		answerInt := parseUserInput(answer)

		result := checkGuess(randomNumber, answerInt)
		fmt.Println(result)

		if result == "You are right" {
			return
		}
	}
	fmt.Println("I'm sorry, but you lose")
}

func generateRandomNumber(max int) int {
	return rand.Intn(max) + 1
}

func getUserInput(reader *bufio.Reader) string {
	answer, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
	return strings.TrimSuffix(answer, "\r\n")
}

func parseUserInput(answer string) int {
	answerInt, err := strconv.Atoi(answer)
	if err != nil {
		fmt.Println("Error parsing input:", err)
		os.Exit(1)
	}
	return answerInt
}

func checkGuess(guess, target int) string {
	if guess == target {
		return "You are right"
	} else if guess > target {
		return "My number is bigger"
	} else {
		return "My number is smaller"
	}
}
