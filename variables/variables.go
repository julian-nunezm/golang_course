package main

import (
	"fmt"
	"strconv"
)

var gender = "Male"	//It doesn't work :=

func main(){
	//general()
	//strings()
	//variableTypes()
	//operators()
	//arrays()
	//slices()
	maps()
}

func general(){
	fmt.Println("\nVariables\n***************")
	var num int	//Static Type in Go
	num = 25
	Num := 101
	fmt.Println(num, Num)
}

func strings(){
	fmt.Println("\nStrings\n***************")
	name := "Julian"	//var name string, name will always be string and := is just to declare new variables
	surname, age := "Nunez", 29
	fmt.Println("Name: " + name + ", Surname: " + surname + ", Age: " + strconv.Itoa(age))
	
	str1, str2 := "First", "String"
	fmt.Println(str1, str2)
	str1, str2 = str2, str1
	fmt.Println(str1, str2)

	var(
		mother = "Teresa"
		father = "Alberto"
	)
	fmt.Println(mother + " y " + father)

	//Scope
	fmt.Println("Gender: " + gender)
	print()

	//String functions
	fmt.Println("Length of name:", len(name))
	fmt.Println("First char:", name[0], "Last char:", name[len(name)-1])	//UTF-8 Format
	fmt.Println("Mother's Nickname:", mother[:4])

	//Special characters
	chain := `
	<html>
		<head>
			<meta charset="utf-8">
			<title></title>
		</head>
		<body>

		</body>
	</html>`
	fmt.Println(chain)
}

func print(){
	fmt.Println("Gender: " + gender)
}

func variableTypes(){
	fmt.Println("\nTypes of variables and default value\n************************************")
	var number int
	var str string
	var flag bool
	var flNum float32
	var variable byte
	fmt.Println("Int:", number, "String:", str, "Boolean:", flag, "Float:", flNum, "Byte:", variable)
	number = 15
	str = "Test"
	flag = true
	flNum = 12.75
	fmt.Println("Int:", number, "String:", str, "Boolean:", flag, "Float:", flNum, "Byte:", variable)
}

func operators(){
	fmt.Println("\nOperators\n***************")
	var(
		a = 1
		b = 3
		c = 5
	)
	fmt.Println("1 + 3 = ", a + b)
	fmt.Println("5 - 3 = ", c - b)
	fmt.Println("3 * 5 = ", b * c)
	fmt.Println("5 / 3 = ", c / b)
	fmt.Println("5 % 3 = ", c % b)
	fmt.Println("**************")
	fmt.Println("5 == 3 = ", c == b)
	fmt.Println("3 == 3 = ", b == b)
	fmt.Println("5 != 3 = ", c != b)
	fmt.Println("3 != 3 = ", b != b)
	fmt.Println("**************")
	fmt.Println("True and True = ", true && true)
	fmt.Println("True and False = ", true && false)
	fmt.Println("False and False = ", false && false)
	fmt.Println("True or True = ", true || true)
	fmt.Println("True or False = ", true || false)
	fmt.Println("False or False = ", false || false)
	fmt.Println("**************")
	fmt.Println("Not True = ", !true)
	fmt.Println("Not False = ", !false)
}

func arrays(){
	fmt.Println("\nArrays\n***************")
	var x [3]int 	//Array
	fmt.Println("x:",x)
	x[1] = 25
	fmt.Println("x:",x)
	fmt.Println("x[1]:",x[1])

	var y [3][2]bool 	//Matrix
	fmt.Println("y:",y)
	y[1][1] = true
	fmt.Println("y:",y)

	z := [...]float32{1.1,2.5,3.4,4.7,5.3}
	fmt.Println("z:",z)
	fmt.Println("Length z:",len(z))
}

func slices(){
	//Like ArrayList: Slice's size can change on execution time
	var i []int
	fmt.Println("Slice i:",i)

	j := []int{8,5,6,7,13}
	fmt.Println("Slice j:",j)
	
	k := make([]int, 6, 10)
	fmt.Println("Slice k:",k)
	fmt.Println("Length k:",len(k))
	fmt.Println("Capacity k:",cap(k))

	a := [10]int{0,1,2,3,4,5,6,7,8,9}	//Array
	fmt.Println("Array a:",a)
	
	var b []int
	b = a[3:6]
	fmt.Println("Slice b:",b)
	fmt.Println("Length b:",len(b))
	fmt.Println("Capacity b:",cap(b))
	
	b[0] = 10
	fmt.Println("Set other numeric value to first index on b...")
	fmt.Println("Slice a:",a)
	fmt.Println("Slice b:",b)

	c := make([]byte, 4, 10)
	fmt.Println("Slice c:",c)
	c = []byte{'H','E','L','L','O'}
	fmt.Println("Slice c:",c)
	fmt.Printf("Slice c: %q\n",c)
	fmt.Println("C's capacity: ",cap(c))
	for i := 0; i < len(c); i++ {
		fmt.Printf("Slice c[%d]: %q\n",i,c[i])
	}
	fmt.Println("C after append...")
	c = append(c,' ','W','O','R','L','D')
	fmt.Printf("Slice c: %q\n",c)
	fmt.Println("New c's capacity: ",cap(c))

	from := []int{1,2,3}
	to := []int{4,5,6}
	fmt.Println(from, to)
	copy(to, from)
	fmt.Println(from, to)
}

func maps(){
	//Dictionaries or Hash Tables, they are dynamic
	//It should be necessary set the capacity for performance if the size is known
	//The map's elements are not variables, so we cannot get its addresses.
	m := make(map[string]string)
	m["name"] = "Julian"
	m["age"] = "29"
	fmt.Println(m)
	fmt.Println("Name:",m["name"])

	ages := map[string]int{
		"Julian": 29,
		"Karen": 25,
		"Teresa": 62,
		"Alberto": 76,
	}
	fmt.Println(ages)
	fmt.Println("\"Alberto\" is deleted")
	delete(ages,"Alberto")
	fmt.Println(ages)
	delete(ages,"Ana")
	fmt.Println("Pedro is",ages["Pedro"])
	ages["Pedro"]++
	ages["Carlos"]+=25
	fmt.Println(ages)
	fmt.Println("\n--Range loop--")
	for index := range ages {
		fmt.Println("Looking for "+index+"...")
		if age, ok := ages[index]; ok {
			if age < 18 {
				fmt.Println(index+"'s age is less than 18.")
			} else {
				fmt.Println(index+"'s age is more than 18.")
			}
		} else {
			fmt.Println(index+" is not in the map.")
		}
	}
	fmt.Println("Looking for Manuel...")
	if age, ok := ages["Manuel"]; ok {
		if age < 18 {
			fmt.Println("His age is less than 18.")
		} else {
			fmt.Println("His age is more than 18.")
		}
	} else {
		fmt.Println("He is not in the map.")
	}
	
	fmt.Println("\nAges has "+strconv.Itoa(len(ages))+" elements.")

	days := map[int]string{
		1: "Monday",
		2: "Tuesday",
		3: "Wednesday",
		4: "Thursday",
		5: "Friday",
		6: "Saturday",
		7: "Sunday",
	}
	//fmt.Println(days)
	for index, value := range days {
		fmt.Printf("%s is the day number %d of the week.\n", value, index)
	}

}