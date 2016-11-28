package main

import (
	"fmt"
	"os"
)

func main(){
	//os library allows us as interface to get access to OS functions. Some functons don't work depends on the OS.
	//Printing the working directory
	//wd,_ := os.Getwd()
	//fmt.Println("Current work directory: "+wd)
	//Opening file
	//f, err := os.Open(wd+"/go_lang/src/github/julian-nunezm/defer/text.txt")
	f, err := os.Open("VGM.csv")
	//f, err := os.Open("testdisk.log")
	//Looking for errors
	if err != nil {
		panic(err)
	}
	defer f.Close()						//It executes at the end of main
	//Saving the content of the file
	data := make([] byte, 1543)
	c, err := f.Read(data)
	//Looking for errors
	if err != nil {
		panic(err)
	}
	//Printing the content
	fmt.Printf("Read bytes: %d\n Read text: %q\n error: %v\n", c, data, err)
}