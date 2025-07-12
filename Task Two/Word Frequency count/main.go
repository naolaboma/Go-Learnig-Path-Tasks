package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func createSentence() string {
	reader := bufio.NewReader(os.Stdin)
	inpt, _ := getInput("Enter Your Sentence: ", reader)

	return inpt
}
func replaceP(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsPunct(r) {
			return -1
		}
		return r
	}, s)
}
func Counterf() map[string]int {
	ride := createSentence()
	removedPunc := strings.ToLower(replaceP(ride))
	words := strings.Fields(removedPunc)

	outp := map[string]int{}
	for _, word := range words {
		outp[word]++
	}
	return outp
}

func main() {
	myres := Counterf()
	fmt.Println(myres)
}
