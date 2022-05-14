package problems

import "fmt"

func GreatestCommonFactorSolved() {
	fmt.Println("---- GreatestCommonFactor problem -----")
	fmt.Println("input: num1=36, num2=54")
	fmt.Println("result:", greatestCommonFactor(54, 36))
}

func greatestCommonFactor(num1 int, num2 int) int {
	for i := num1; i > 0; i-- {
		if num1%i == 0 && num2%i == 0 {
			return i
		}
	}

	return 1
}
