package main

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"strings"
	"time"
)

//please uncomment desired function to execute
func main() {
	//circle()
	//palindrome()
	//evenOdd()
}


//functions

func circle() {

	var rad float64
	const PI float64 = 3.14
	var area float64
	var ci float64

	fmt.Print("Enter radius of Circle \n")
	fmt.Scanln(&rad)

	area = math.Pow(rad, 2) * PI
	fmt.Println("The area is:" , area)

	ci = rad * (2 * PI)
	fmt.Println("The circumference is:", ci)

}

func evenOdd() {

	var choice string
	var result string

	//different random number each time
	rand.Seed(time.Now().UnixNano())

	var number int = rand.Intn(100)

	fmt.Print("Choose even or odd \n")

	fmt.Scanln(&choice)

	if number % 2 == 0 && choice == "even" {
		fmt.Println("number ", number  ,"is even", "You win your bet")
		 result = "even"
	} else {
		fmt.Println("number ", number, "is odd")
		 result = "odd"
	}

	if(result == choice) {
		fmt.Println("You win user")
	} else {
		fmt.Println("You lose user")

	}
}

func palindrome() {
	var input string = ""
	for {
		if (input == "") {			
			fmt.Println("input a word to check if it is palindrome")
			fmt.Scanln(&input)
		} else {
			break
		}
	}

	//lower case 
	toLower := strings.ToLower(input)

	//split string input
	var splitted = strings.Split(toLower , "")

	//reverse splitted string 
	var reversed []string

	//loop through each word and append to new slice 
	for i := len(splitted) - 1 ; i >= 0 ; i-- {
		fmt.Println(splitted[i])
		reversed = append(reversed,splitted[i])
	}
	fmt.Println(reversed, splitted)

	//check if reversed string is equal to original
	if reflect.DeepEqual(reversed, splitted ) {
		fmt.Println("Yes PALINDROME")
	} else {
		fmt.Println("NOT PALINDROME")
	}

}