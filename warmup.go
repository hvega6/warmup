package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"
	"unicode"
)

// 1. Palindrome Checker
func isPalindrome(s string) bool {
	// Remove non-alphanumeric characters and convert to lowercase
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	clean := strings.ToLower(reg.ReplaceAllString(s, ""))
	
	// Compare the string with its reverse
	for i := 0; i < len(clean)/2; i++ {
		if clean[i] != clean[len(clean)-1-i] {
			return false
		}
	}
	return true
}

// 2. Convert Celsius to Fahrenheit
func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

// 3. Count Vowels in a String
func countVowels(s string) int {
	vowels := map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	count := 0
	for _, char := range strings.ToLower(s) {
		if vowels[char] {
			count++
		}
	}
	return count
}

// 4. Check if a Number is Prime
func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// 5. Capitalize First Letter of Every Word
func capitalizeWords(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		if len(word) > 0 {
			r := []rune(word)
			r[0] = unicode.ToUpper(r[0])
			words[i] = string(r)
		}
	}
	return strings.Join(words, " ")
}

// 6. Find the Minimum in a Slice
func findMinimum(arr []int) int {
	if len(arr) == 0 {
		return 0 // or panic, depending on how you want to handle empty slices
	}
	min := arr[0]
	for _, num := range arr {
		if num < min {
			min = num
		}
	}
	return min
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama")) // true
	fmt.Println(celsiusToFahrenheit(25))                        // 77
	fmt.Println(countVowels("Hello, World!"))                   // 3
	fmt.Println(isPrime(17))                                    // true
	fmt.Println(capitalizeWords("hello world"))                 // "Hello World"
	fmt.Println(findMinimum([]int{3, 1, 4, 1, 5, 9, 2, 6}))     // 1
}