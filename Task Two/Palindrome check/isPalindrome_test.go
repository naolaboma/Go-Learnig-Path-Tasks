package main

import (
	"testing"
)

func TestClearSt(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello, World!", "helloworld"},
		{"123abc!!", "123abc"},
		{"Go-Lang_123", "golang123"},
		{"", ""},
		{"😊123", "123"}, // emoji is not letter/digit
	}

	for _, test := range tests {
		result := clearSt(test.input)
		if result != test.expected {
			t.Errorf("clearSt(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"Madam", true},
		{"Racecar", true},
		{"A man, a plan, a canal: Panama", true},
		{"No lemon, no melon", true},
		{"Hello, World!", false},
		{"12321", true},
		{"Not a palindrome", false},
	}

	for _, test := range tests {
		result := isPalindrome(test.input)
		if result != test.expected {
			t.Errorf("isPalindrome(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}
