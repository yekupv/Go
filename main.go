package main

import "fmt"

func main() {
	// Feel free to use the main function for testing your functions
	fmt.Println("Hello, à¤¦à¥à¤¨à¤¿à¤¯à¤¾!")
	fmt.Println(Fizzbuzz(15))
	fmt.Println(IsPrime(2))
	fmt.Println(IsPalindrome("kek"))

}

// Fizzbuzz is a classic introductory programming problem.
// If n is divisible by 3, return "Fizz"
// If n is divisible by 5, return "Buzz"
// If n is divisible by 3 and 5, return "FizzBuzz"
// Otherwise, return the empty string
func Fizzbuzz(n int) string {
	// TODO
	if n%3 == 0 && n%5 == 0 {
		return "FizzBuzz"
	} else if n%5 == 0 {
		return "Buzz"
	} else if n%3 == 0 {
		return "Fizz"
	}
	return ""
}

// IsPrime checks if the number is prime.
// You may use any prime algorithm, but you may NOT use the standard library.
func IsPrime(n int) bool {
	// TODO
	if n < 2 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// IsPalindrome checks if the string is a palindrome.
// A palindrome is a string that reads the same backward as forward.
func IsPalindrome(s string) bool {
	arr := []rune(s)
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return string(arr) == s
}
