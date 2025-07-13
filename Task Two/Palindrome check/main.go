package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func getInput(prompt string, r *bufio.Reader) string {
	fmt.Print(prompt)
	input, _ := r.ReadString('\n')
	return strings.TrimSpace(input)
}

func clearSt(s string) string {
	var b strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteRune(unicode.ToLower(r))
		}
	}
	return b.String()
}

func isPalindrome(str string) bool {
	str = clearSt(str)
	l, r := 0, len(str)-1
	for l < r {
		if str[l] != str[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	str := getInput("Enter a string: ", reader)
	if isPalindrome(str) {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not a Palindrome")
	}
}
