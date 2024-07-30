package main

import "fmt"

func incrementNumber () func() int {
	var count int = 0;

	return func() int {
		count ++
		return count
	}
}

func main(){
	increment := incrementNumber()
	fmt.Println(increment())
	fmt.Println(increment())
}


// hitshiroya@hshiroya-mac closures % go run closures.go
// 1
// 2