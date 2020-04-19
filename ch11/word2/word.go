package word

import "unicode"

func IsPalindrome(s string) bool {
	var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	for i := range letters {
		if letters[i] != letters[len(letters) - 1 - i] {
			return false
		}
	}
	return true
}

func IsPalindrome2(s string) bool {
	var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	n := len(letters)/2
	for i := 0; i < n; i++ {
		if letters[i] != letters[len(letters) - 1 - i] {
			return false
		}
	}
	return true
}

func IsPalindrome3(s string) bool {
	var letters []rune = make([]rune, 0, len(s))
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	n := len(letters)/2
	for i := 0; i < n; i++ {
		if letters[i] != letters[len(letters) - 1 - i] {
			return false
		}
	}
	return true
}