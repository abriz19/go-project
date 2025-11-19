// Fundamentals of Go Tasks
// Task:  Word Frequency Count
// Write a Go function that takes a string as input and returns a dictionary containing the frequency of each word in the string. Treat words in a case-insensitive manner and ignore punctuation marks.
// [Optional]: Write test for your function

// Task : Palindrome Check
// Write a Go function that takes a string as input and checks whether it is a palindrome or not. A palindrome is a word, phrase, number, or other sequence of characters that reads the same forward and backward (ignoring spaces, punctuation, and capitalization).
// [Optional]: Write test for your function

package main

import (
	"fmt"
	"regexp"
	"strings"
)

func WordFrequency(text string) map[string]int {
	text = strings.ToLower(text)

	re := regexp.MustCompile(`[^\w\s]`)
	text = re.ReplaceAllString(text, "")

	words := strings.Fields(text)

	frequency := make(map[string]int)
	for _, word := range words {
		frequency[word]++
	}

	return frequency
}


func IsPalindrome(s string) bool {
	s = strings.ToLower(s)

	re := regexp.MustCompile(`[^a-z0-9]`)
	s = re.ReplaceAllString(s, "")

	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}


func main() {
	// Test WordFrequency function
	text := "Hello, hello! How are you? You look well. Well, well, well..."
	freq := WordFrequency(text)
	fmt.Println(freq)

	// Test IsPalindrome function
	fmt.Println(IsPalindrome("Hello"))
	fmt.Println(IsPalindrome("A man, a plan, a canal: Panama")) 
}
