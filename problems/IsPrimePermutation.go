package problems

import (
	"fmt"
	"math"
	"strconv"
)

func IsPrimePermutation_Solved() {
	fmt.Println("---- IsPrimePermutation problem -----")
	fmt.Println("input: 910")
	fmt.Println("result:", isPrimePermutation(910))
	fmt.Println("input: 880")
	fmt.Println("result:", isPrimePermutation(880))
}

func isPrimePermutation(num int) int {
	perms := permutations(strconv.Itoa(num))
	for _, strNum := range perms {
		num, err := strconv.Atoi(strNum)
		if err != nil {
			return 0
		}
		if isPrime(num) {
			return 1
		}
	}
	return 0
}

func permutations(testStr string) []string {
	var n func(testStr []rune, p []string) []string
	n = func(testStr []rune, p []string) []string {
		if len(testStr) == 0 {
			return p
		} else {
			result := []string{}
			for _, e := range p {
				result = append(result, join([]rune(e), testStr[0])...)
			}
			return n(testStr[1:], result)
		}
	}

	output := []rune(testStr)
	return n(output[1:], []string{string(output[0])})
}

func join(ins []rune, c rune) (result []string) {
	for i := 0; i <= len(ins); i++ {
		result = append(result, string(ins[:i])+string(c)+string(ins[i:]))
	}
	return
}

func isPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}
