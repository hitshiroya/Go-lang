package main

import "fmt"

func hello(){
	fmt.Println("hello there!")
}

func greetMultipleTimes(words string,count int){
	for i:=0;i<count;i++{
		fmt.Println(words)
	}
}
func getSum(num1 , num2  int) int {
	return num1+num2
}
func getSums (num1,num2 int) (int,error){
	total := num1+num2

	if total < 30 {
		return total, fmt.Errorf("less than 30")
	}

	return total,nil
}

func main () {
	hello()
	word := "Hello there!"
	greetMultipleTimes(word,5)
ans := getSum(10,15)
fmt.Println(ans)
val, err := getSums(10,5)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(val)
}

