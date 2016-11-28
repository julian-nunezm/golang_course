package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sum(a int, b int) int{
	return a+b
}

func subs(a, b int) (r int){
	r = a-b
	return
}

func f1(){
	fmt.Println("Start f1")
	f2()
	fmt.Println("End f1")
}

func f2(){
	fmt.Println("Start f2")
	f3()
	fmt.Println("End f2")
}

func f3(){
	fmt.Println("Start f3")
	fmt.Println("End f3")
}

func prod(numbers ...int) int{
	result := 1
	for _, value := range numbers {
		result *= value
	}
	return result
}

func conc(chain string, chains ...string) string{
	chain += " - "
	for _, ch := range chains {
		chain += ch+" "
	}
	return chain
}

func printEmail(user string) string{
	return user+"@gmail.com"
}

func callPrintEmail(user string, printEmail func(string) string) string{
	return printEmail(user)+" from callPrintEmail"
}

func concPoints() func() string{
	s := ""
	return func() string {
		s += " . "
		return s
	}
}

func pow(n float64) (a, b, c float64){
	a = n*n
	b = n*n*n
	c = n*n*n*n
	return
}

func main(){
	a := 151
	b := 684
	fmt.Println("Sum of "+strconv.Itoa(a)+" and "+strconv.Itoa(b)+" is "+strconv.Itoa(sum(a,b)))
	fmt.Println("Substraction of "+strconv.Itoa(a)+" and "+strconv.Itoa(b)+" is "+strconv.Itoa(subs(a,b)))

	//Functions's sign
	fmt.Printf("Sum function's sign: %T\n",sum)
	fmt.Printf("Subs function's sign: %T\n",subs)

	//package math
	//func Sin(x float64) float64	//implemented in assembly language

	//Call Stack -> It is solved by LIFO
	f1()

	//Variadic functions
	fmt.Println(prod(2,3,4,5,6,7))
	fmt.Println(prod(2,2,2,2,2))
	fmt.Println(prod(2,1,1))
	numbers := []int{
		25,
		56,
		32,
	}
	fmt.Println(prod(numbers...))
	fmt.Println(conc("Test", "This", "is", "how", "Variadic", "functions", "works."))

	//Closure: Manage functions as values
	user := "julian.nunezm"
	email := printEmail
	fmt.Println(email(user))

	student := "Julian"
	role := func () string{
		return "student"
	}
	fmt.Println(student+" is a "+role())

	fmt.Println(callPrintEmail(user, printEmail))

	var f func()
	if f == nil{
		fmt.Println("f is a nil function")
	}

	//Closure: It saves the state of its variables
	p := concPoints()
	fmt.Println(p())
	fmt.Println(p())
	fmt.Println(p())
	fmt.Println(p())
	fmt.Println(p())

	//Anonymous functions
	str := "@ABYZ6789"
	fmt.Println("Original string:",str)
	str = strings.Map(func (r rune) rune {
			return r+1
		}, str)
	fmt.Println("Modified string:",str)

	//Returning multiple values
	fmt.Println(pow(3.87))
}