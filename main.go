package main

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func MathChallenge(num int) string {
	// code goes here
	numInFloat := float64(num)

	return strconv.FormatBool(math.Ceil(math.Log2(numInFloat)) == math.Floor(math.Log2(numInFloat)))
}

func StringChallenge(str string) string {
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

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	// fmt.Println(MathChallenge(16))
	// fmt.Println(StringChallenge("Noelseesleon"))
	fmt.Println(ArrayChallenge([]int{5, 2, 2, 1}))
}

func ArrayChallenge(arr []int) int {
	mode := 0
	count2 := 0
	for i := 0; i < len(arr); i++ {
		var count = 0
		for j := 0; j < len(arr); j++ {
			if arr[i] == arr[j] && j != i {
				count++
			}
			if count > count2 {
				mode = arr[i]
				count2 = count
				count = 0
			}
		}
	}
	if count2 == 0 {
		return -1
	}
	return mode
}
