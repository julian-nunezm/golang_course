package main

import(
	"fmt"
	"strconv"
)

func main() {
	var num1, num2, counter int

	header := `
	*******************
	Odd numbers counter
	*******************`
	fmt.Println(header)

	fmt.Println("Type the first number")
	fmt.Scanln(&num1)
	fmt.Println("Type the second number")
	fmt.Scanln(&num2)
	//fmt.Println(num1, num2)

	for i := num1; i <= num2; i++ {
		if i%2 != 0 {
			counter++
			fmt.Println(strconv.Itoa(i)+" is an odd number")
		}
	}
	fmt.Printf("It founded %d odd numbers between %d and %d\n",counter,num1,num2)
}