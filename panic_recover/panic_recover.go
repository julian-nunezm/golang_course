package main

import (
	"fmt"
)

func connectDB() (bool, string) {
	errorMsg := "Error in connectDB: It cannot connect to the DB"
	return false, errorMsg
}

func emailBlockfreight(user string) {
	fmt.Println(user+"@blockfreight.com")
}

func main() {
	emailBlockfreight("julian.nunez")
	db, err := connectDB()
	if(!db) {
		defer func() {
			errorMsg := recover()
			fmt.Println(errorMsg)
		}()
		panic(err)
	}
}