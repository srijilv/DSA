package main

import (
	"fmt"
	"regexp"
)

func main() {

	a := []byte("amazing")
	pattern := "010"
	l := len(pattern)

	combinations := []string{}

	for i := 0; i < len(a) && len(a) >= i+l; i++ {
		end := i + l
		aa := a[i:end]
		combinations = append(combinations, string(aa))
	}

	fmt.Println(combinations)
	counter := 0
	for _, val := range combinations {

		counter += checkPaternMatch(val, pattern)
	}
	fmt.Printf("counter: %v\n", counter)

}
func checkPaternMatch(str, pat string) int {
	vowels := "[aeiou]"
	rr, err := regexp.Compile(vowels)
	if err != nil {
		panic(err)
	}

	for i, val := range pat {
		if string(val) == "0" && !rr.MatchString(string(str[i])) {
			return 0
		} else if string(val) == "1" && rr.MatchString(string(str[i])) {
			return 0
		}
	}
	return 1
}
