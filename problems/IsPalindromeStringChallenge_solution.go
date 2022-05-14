package problems

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func IsPalindromeStringChallenge_Solved() {
	fmt.Println("---- IsPalindromeStringChallenge problem -----")
	fmt.Println("input: Noel-sees-leon")
	fmt.Println("result:" + isPalindromeStringChallenge("Noel-sees-leon"))
	fmt.Println("input: keke")
	fmt.Println("result:" + isPalindromeStringChallenge("keke"))
}

func isPalindromeStringChallenge(str string) string {
	// Clear non alphabetic chars from str
	reg, err := regexp.Compile("[^A-Za-z]+")
	if err != nil {
		log.Fatal(err)
		return "false"
	}
	str = strings.ToLower(reg.ReplaceAllString(str, ""))
	fmt.Println(str)

	lenOfStr := len(str)

	for i := 0; i < lenOfStr/2; i++ {
		if str[i] != str[lenOfStr-i-1] {
			return "false"
		}
	}

	return "true"
}
