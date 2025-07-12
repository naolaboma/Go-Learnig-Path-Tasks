package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func createStGrade() Grade {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Enter your name please: ", reader)
	noSubjStr, _ := getInput("Enter Number of Subject: ", reader)
	noSubj, err := strconv.Atoi(noSubjStr)
	if err != nil || noSubj <= 0 {
		fmt.Println("Invalid number of subjects. Please enter a positive integer.")
		os.Exit(1)
	}
	g := newStGrade(name)

	for limit := 0; limit < noSubj; limit++ {
		subject, _ := getInput("Enter the subject name: ", reader)
		gradeStr, _ := getInput("Enter the grade: ", reader)

		gradeVal, err := strconv.ParseFloat(gradeStr, 64)
		if err != nil || gradeVal < 0 || gradeVal > 100  {
			fmt.Println("invalid! Grade must be between 0 and 100")
			os.Exit(1)
		}
		g.addName(subject, gradeVal)
	}

	g.save()
	return g
}

func main() {
	mygrade := createStGrade()
	fmt.Println(mygrade.format())
}
