// Exercises from codingbat.com ported to Golang.
package warmup2

import (
	"strings"
)

// Given an slice of ints, return true if the sequence of numbers 1, 2, 3
// appears in the slice somewhere.
func Array123(s []int) bool {
	for i, _ := range s {
		if s[i] == 1 {
			if i+1 < len(s) && s[i+1] == 2 {
				if i+2 < len(s) && s[i+2] == 3 {
					return true
				}
			}
		}
	}
	return false
}

// Given a slice of ints, return the number of 9's in the slice.
func ArrayCount9(s []int) int {
	ct := 0
	for _, v := range s {
		if v == 9 {
			ct++
		}
	}
	return ct
}

// Given a slice of ints, return true if one of the first 4 elements in the
// slice is a 9. The slice length may be less than 4.
func ArrayFront(s []int) bool {
	idx := 3
	for i, v := range s {
		if i <= idx && v == 9 {
			return true
		}
	}
	return false
}

// Given a string and a non-negative int n, we'll say that the front of the
// string is the first 3 chars, or whatever is there if the string is less
// than length 3. Return n copies of the front
func FrontTimes(s string, n int) string {
	var str strings.Builder
	var idx int
	if len(s) < 3 {
		idx = len(s)
	} else {
		idx = 3
	}
	for i := 0; i < n; i++ {
		str.WriteString(s[0:idx])
	}
	return str.String()
}

// Given a string, return the count of the number of times that a substring length
// 2 appears in the string and also as the last 2 chars of the string, so
// "hixxxhi" yields 1 (we won't count the end substring).
func Last2(s string) int {
	if len(s) < 2 {
		return 0
	}
	//regex doesn't work for overlapping matches
	// rx := regexp.MustCompile(string(s[len(s)-2:]))
	// rx.FindAllString(s[:len(s)-2], -1)
	// return len(rx.FindAllStringSubmatch(s[:len(s)-2], -1))
	match := string(s[len(s)-2:])
	count := -1 //Offset the last match
	for i := range s {
		if (i+len(match) <= len(s)) && string(s[i:i+len(match)]) == match {
			count++
		}
	}
	return count
}

// Given a string, return a new string made of every other char starting with
// the first, so "Hello" yields "Hlo".
func StringBits(s string) string {
	var str strings.Builder
	for i, v := range s {
		if i%2 == 0 {
			str.WriteString(string(v))
		}
	}
	return str.String()
}

// Given 2 strings, a and b, return the number of the positions where they contain
// the same length 2 substring. So "xxcaazz" and "xxbaaz" yields 3, since the
// "xx", "aa", and "az" substrings appear in the same place in both strings.
//"xxbaaz"
//"xxcaazz"
func StringMatch(a, b string) int {
	count := 0
	for i, v := range a {
		if (i < len(b) && string(v) == string(b[i])) && (i+1 < len(a) && i+1 < len(b) && string(a[i+1]) == string(b[i+1])) {
			count++
		}
	}
	return count
}

// Given a non-empty string like "Code" return a string like "CCoCodCode".
func StringSplosion(s string) string {
	var str strings.Builder
	for i, _ := range s {
		str.WriteString(string(s[0 : i+1]))
	}
	return str.String()
}

// Given a string and a non-negative int n, return a larger string that is n
// copies of the original string.
func StringTimes(s string, n int) string {
	var str strings.Builder
	for i := 0; i < n; i++ {
		str.WriteString(s)
	}
	return str.String()
}
