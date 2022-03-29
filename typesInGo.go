package main

import (
	"fmt"
)

func tellMeTypes() {

	//Insert variables declarations here based on activity
	text := "The following is the account information."
	firstName := "Luke"
	lastName := "Skywalkter"
	age := 20
	weight := 73.0
	height := 1.72
	credits := 123.55
	accountName := "admin"
	accountPw := "password"
	subscribed := true

	//insert code here to print out types of variables
	fmt.Printf("%T %T %T %T %T %T %T %T %T %T\n", text, firstName, lastName, age, weight, height, credits, accountName, accountPw, subscribed)
}

func main() {
	tellMeTypes()
}
