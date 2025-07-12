package main

import (
	"strings"
	"testing"
)

func TesWordCount(t *testing.T) {
	text := "Go, Go! Go? go."
	expected := map[string]int{"go": 4}
	text = strings.ToLower(replaceP(text))
	words := strings.Fields(text)

	actual := make(map[string]int)
	for _, word := range words {
		actual[word]++
	}

	for word, expCount := range expected {
		if actual[word] != expCount {
			t.Errorf("For word %q, expected %d but got %d", word, expCount, actual[word])
		}
	}
}
