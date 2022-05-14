package problems

import (
	"fmt"
	"math"
	"strconv"
)

func PowerOfTwoChecker_Solved() {
	fmt.Println("---- PowerOfTwoChecker problem -----")
	fmt.Println("Input:124")
	fmt.Println("Result:", powerOfTwoChecker(124))
}

func powerOfTwoChecker(num int) string {
	// code goes here
	numInFloat := float64(num)

	return strconv.FormatBool(math.Ceil(math.Log2(numInFloat)) == math.Floor(math.Log2(numInFloat)))
}
