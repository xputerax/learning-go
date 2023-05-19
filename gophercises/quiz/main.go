package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Question struct {
	Name   string
	Answer string
}

func ReadQuestionsFromFile(fh *os.File) ([]Question, error) {
	csvLines, csvErr := csv.NewReader(fh).ReadAll()

	if csvErr != nil {
		return nil, fmt.Errorf("Failed to read csv: %s\n", csvErr)
	}

	questions := []Question{}

	for _, line := range csvLines {
		name := line[0]
		answer := line[1]
		q := Question{
			Name:   name,
			Answer: answer,
		}

		questions = append(questions, q)
	}

	return questions, nil
}

func main() {
	questionFile := "problems.csv"

	if len(os.Args) > 1 {
		questionFile = os.Args[1]
	}

	fmt.Printf("Reading questions from %s....\n", questionFile)

	fh, fErr := os.Open(questionFile)

	if fErr != nil {
		fmt.Errorf("Failed to open file for reading: %s", fErr)
		os.Exit(1)
	}

	questions, questionErr := ReadQuestionsFromFile(fh)

	if questionErr != nil {
		fmt.Errorf("Question read error: %s\n", questionErr)
		os.Exit(1)
	}

	totalQuestions := len(questions)
	correctCount := 0
	wrongCount := 0

	for idx, question := range questions {
		fmt.Printf("%d) %s = ", idx+1, question.Name)

		answer := ""
		_, scanErr := fmt.Scanln(&answer)

		if scanErr != nil {
			fmt.Printf("Scan error: %s\n", scanErr)
		} else {
			if answer == question.Answer {
				correctCount++
			} else {
				wrongCount++
			}
		}
	}

	var correctPercentage float32 = float32(correctCount) / float32(totalQuestions) * 100
	var wrongPercentage float32 = float32(wrongCount) / float32(totalQuestions) * 100

	fmt.Println("\n=== Finished ===")
	fmt.Printf("Correct: %d (%.2f%%)\nWrong: %d (%.2f%%)\n",
		correctCount, correctPercentage,
		wrongCount, wrongPercentage)
}
