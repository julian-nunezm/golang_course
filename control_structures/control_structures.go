package main

import(
	"fmt"
)

func main(){
	//for_loop()
	//if_structure()
	//switch_structure()
	range_loop()
}

func for_loop(){
	fmt.Println("\nFor Loop:")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func if_structure(){
	fmt.Println("\nIf Structure:")
	a, b := 8, 7
	fmt.Printf("a=%d, b=%d\n",a,b)
	if a < b {
		fmt.Println("a is less than b")
	} else if a == b {
		fmt.Println("a is equal to b")
	} else {
		fmt.Println("a is more than b")
	}
}

func switch_structure() {
	fmt.Println("\nSwitch Structure (way 1):")
	var day int
	fmt.Println("Type the number of one day in week...")
	fmt.Scanln(&day)

	switch day {
	case 1:
		fmt.Println("You chose Monday")
	case 2:
		fmt.Println("You chose Tuesday")
	case 3:
		fmt.Println("You chose Wednesday")
	case 4:
		fmt.Println("You chose Thursday")
	case 5:
		fmt.Println("You chose Friday")
	case 6:
		fmt.Println("You chose Saturday")
	case 7:
		fmt.Println("You chose Sunday")
	default:
		fmt.Println("You chose an invalid day")
	}

	fmt.Println("\nSwitch Structure (way 2):")
	var number int
	fmt.Println("Type a number...")
	fmt.Scanln(&number)

	switch {
	case number < 0:
		fmt.Printf("%d is negative\n", number)
		fallthrough
	case number < -5:
		fmt.Printf("%d is less than -5\n", number)
		fallthrough
	case number >= 0:
		fmt.Printf("%d is more or equal than zero\n", number)
		fallthrough
	case number >= 5:
		fmt.Printf("%d is more than 5\n", number)
		fallthrough
	default:
		fmt.Printf("%d is a simple number\n", number)
	}
}

func range_loop(){
	//It is the same for each loop
	fmt.Println("\nRange loop:")
	names := []string{"Batman", "Superman", "Hulk", "Spiderman"}
	for index, value := range names {
		fmt.Printf("%q is in index %d\n",value,index)
	}
	//If we don't need index variable, replace it with _
	for _, value := range names {
		fmt.Println(value)
	}
	phrase := "I am learning Golang"
	for index, value := range phrase {
		fmt.Printf("Char %q is in index %d\n",value,index)
	}
}