package string2

import (
	"math/rand"
	"regexp"
	"strings"
	"time"
)

// Given a string, return a string where for every char in the original, there are two chars.
func DoubleChar(s string) string {
	var stringBuilder strings.Builder
	for _, r := range s {
		stringBuilder.WriteString(string(r))
		stringBuilder.WriteString(string(r))
	}
	return stringBuilder.String()
}

// Return the number of times that the string "code" appears anywhere in the
// given string, except we'll accept any _letter_ for the "d", so "cope" and "cooe" count.
// Try not to use regex
func IndexInArray[T any](arr []T, idx int) bool {
	return idx < len(arr)
}

func CountCode(s string) int {
	count := 0
	//var count int
	//count = 0
	for i, _ := range s {
		//finding `co`
		cIndex := i
		oIndex := i + 1
		wcIndex := i + 2
		eIndex := i + 3
		if s[cIndex] == 'c' && IndexInArray([]rune(s), oIndex) && s[oIndex] == 'o' {
			legal := false
			for _, v := range "abcdefghijklmnopqrstuvwxyz" {
				if IndexInArray([]rune(s), wcIndex) && v == rune(s[wcIndex]) {
					legal = true
				}
			}
			if legal {
				if IndexInArray([]rune(s), eIndex) && s[eIndex] == 'e' {
					count++
				}
			}
		}
	}

	return count
}

// Return the number of times that the string "hi" appears anywhere in the given string.
func CountHi(s string) int {
	//Cause why not. CHoose a branch. Same functionality eitehr way
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(2) == 1 {
		return len(regexp.MustCompile("hi").FindAllStringIndex(s, -1))
	} else {
		count := 0
		for i, _ := range s {
			if i+1 < len(s) && s[i:i+2] == "hi" {
				count++
			}
		}
		return count
	}
}

// Given two strings, return true if either of the strings appears at the very end
// of the other string, ignoring upper/lower case differences (in other words, the
//computation should not be "case sensitive"). Note: s.lower() returns the
// lowercase version of a string.
func EndOther(a, b string) bool {
	aLower := strings.ToLower(a)
	bLower := strings.ToLower(b)
	if len(bLower) > len(aLower) {
		if bLower[len(bLower)-len(aLower):] == aLower {
			return true
		}
	} else {
		if aLower[len(aLower)-len(bLower):] == bLower {
			return true
		}
	}
	return false
}

// Return true if the string "cat" and "dog" appear the same number of times
// in the given string.
func CatDog(s string) bool {
	cat := regexp.MustCompile("cat")
	dog := regexp.MustCompile("dog")
	return len(cat.FindAllStringIndex(s, -1)) == len(dog.FindAllStringIndex(s, -1))
}

// Return true if the given string contains an appearance of "xyz" where the xyz
// is not directly preceeded by a period (.). So "xxyz" counts but "x.xyz" does not.
func XyzThere(s string) bool {
	for i := range s {
		if ((0 < i-1 && s[i-1] != '.') || (i == 0)) && s[i] == 'x' && i+1 < len(s) && s[i+1] == 'y' && i+2 < len(s) && s[i+2] == 'z' {
			return true
		}
	}
	return false
}
