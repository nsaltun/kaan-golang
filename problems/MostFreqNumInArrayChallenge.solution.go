package problems

import "fmt"

func MostFreqNumInArrayChallenge_Solved() {
	fmt.Println("---- MostFreqNumInArrayChallenge problem -----")
	fmt.Println("input: [5,2,2,1]")
	fmt.Println("result:", mostFreqNumInArrayChallenge([]int{5, 2, 2, 1}))
	fmt.Println("input: [5,2,1]")
	fmt.Println("result:", mostFreqNumInArrayChallenge([]int{5, 2, 1}))
}

func mostFreqNumInArrayChallenge(arr []int) int {
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
