package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"unicode"
)

func main() {
	// Part 1 (c)
	fmt.Println("Hello World!")
	m := 2
	n := &m
	*n = *n * 8
	fmt.Printf("m is % d and n is %d\n", m, *n)
	const (
		DaysInYear     = 365
		DaysInLeapYear = 366
	)
	fmt.Printf("Days in Year %d\nDays in Leap Year %d\n", DaysInYear, DaysInLeapYear)

	// Part 1(d)
	first := "My name is Suraj Shrestha."
	second := "I am from Nepal."
	third := "I love playing soccer."
	fifth := "I study in OleMiss."
	fmt.Printf("Lower: %s\n", strings.ToLower(first))
	fmt.Printf("Upper: %s\n", strings.ToUpper(first))
	fmt.Printf("First Word: %s\n", strings.Title(second))
	fmt.Printf("Trim: %s\n", strings.TrimSpace(third))
	fmt.Printf("Length: %d\n", len(fifth))
	fmt.Printf("Replace Shrestha with Stha: %s\n", strings.Replace(first, "Shrestha", "Stha", -1))
	fmt.Printf("Substring: %s\n", second[:9])
	fmt.Printf("Contains Nepal: %t\n", strings.Contains(second, "Nepal"))
	fmt.Printf("Repeat soccer x 3: %s\n", strings.Repeat(third[15:21], 3))
	fmt.Printf("Study is at index: %d\n", strings.Index(fifth, "study"))
	fmt.Printf("Compare last word of fifth sentence with OleMiss: %d\n", strings.Compare(fifth[11:], "OleMiss"))

	// Part 1 (e)
	fmt.Println("\n--------------------\nMath Functions")
	fmt.Printf("Absolute Value: %f\n", math.Abs(-20))
	fmt.Printf("10-Cubed: %f\n", math.Pow(10, 3))
	fmt.Printf("Square root of 25: %f\n", math.Sqrt(25))
	fmt.Printf("Floor of 20.8: %f\n", math.Floor(20.8))
	fmt.Printf("Ceil of 16.8: %f\n", math.Ceil(16.8))
	fmt.Printf("Round of 19.98: %f\n", math.Round(19.98))
	fmt.Printf("20 mod 3 is %f.\n", math.Mod(20, 3))
	fmt.Printf("Random integer 0 to 99: %d\n", rand.Int31n(100))
	if math.IsNaN(math.Sqrt(-2)) {
		fmt.Println("Not a number")
	}
	if unicode.IsDigit('5') {
		fmt.Println("5 is a number")
	}
	if unicode.IsNumber('!') {
		fmt.Println("5 is a number")
	} else {
		fmt.Println("! is NOT a number")
	}
}
