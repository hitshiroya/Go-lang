package main

import "fmt"


func countSum (nums ...int) int {
	count := 0

	for _,num := range nums {
		count += num
	}

	return count
}
func countSumOfArray (nums ...int) int {
	count := 0

	for _,num := range nums {
		count += num
	}

	return count
}


func main(){
	
	nums := []int {1,2,3,4,5,6,7,8,9,10}
	fmt.Println(countSum(1,2,3,4,5,6,7,8,9,10))
	fmt.Println(countSumOfArray(nums...))
}

// hitshiroya@hshiroya-mac varadic_functions % go run variadic.go
// 55
