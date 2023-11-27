package main

import (
	"fmt"
)

var firstSums float64
var secondNumber float64
var Sign string
var Operation string
var Options string

func main() {
	//Start of App
	fmt.Println("Hello to The MicroRocket Calculator-App!!!")
	fmt.Println("If you want to continue please type and enter 'go' or if you want to close the app it's 'end'. ")
	fmt.Scan(&Options)
	//

	//Continue or End App

	if Options == "end" {
		fmt.Println("Thanks for using the app!!!")
		return
	} else if Options == "go" {
		fmt.Println("Ok lets go!!!!")
	}
	//

	for Options == "go" {

		fmt.Println("Please Tell me are adding,subtracting,multiplying or dividing?")
		fmt.Scan(&Operation)
		fmt.Println("Tell me the the first number")
		fmt.Scan(&firstSums)
		fmt.Println("Tell me the the second number")
		fmt.Scan(&secondNumber)

		if Operation == "adding" {
			Answer := firstSums + secondNumber
			fmt.Printf("Answer = %v \n ", Answer)
		} else if Operation == "subtracting" {
			Answer := firstSums - secondNumber
			fmt.Printf("Answer = %v \n ", Answer)
		} else if Operation == "multiplying" {
			Answer := firstSums * secondNumber
			fmt.Printf("Answer = %v \n ", Answer)
		} else if Operation == "dividing" {
			Answer := firstSums / secondNumber
			fmt.Printf("Answer = %v \n ", Answer)
		} else {
			fmt.Println("Looks Like you enter a wrong operation or you did not use numbers.")
			fmt.Println("Please try agian.")
		}

	}

}
