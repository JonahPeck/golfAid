package main

import (
	"fmt"
)

func main() {
	fmt.Println("What is the Yardage?")
	var yardage int
	var club string

	fmt.Scanln(&yardage)
	// fmt.Println("The yardage is ", yardage)

	if yardage < 30 {
		club = "60°"
	} else if yardage >= 30 && yardage < 50 {
		club = "56°"
	} else if yardage >= 50 && yardage < 90 {
		club = "50°"
	} else if yardage >= 90 && yardage < 115 {
		club = "Pitching Wedge"
	} else if yardage >= 115 && yardage < 125 {
		club = "9 Iron"
	} else if yardage >= 125 && yardage < 135 {
		club = "8 Iron"
	} else if yardage >= 135 && yardage < 145 {
		club = "9 Iron"
	} else if yardage >= 145 && yardage < 155 {
		club = "7 Iron"
	} else if yardage >= 155 && yardage < 165 {
		club = "6 Iron"
	} else if yardage >= 165 && yardage < 185 {
		club = "5 Iron"
	} else if yardage >= 185 && yardage < 200 {
		club = "4 Hybrid"
	} else if yardage >= 200 && yardage < 220 {
		club = "2 Hybrid"
	} else {
		club = "Driver"
	}
	fmt.Println("You should hit your", club)

	//what hazards are ahead, what lie is it in, how is your confidence right now

}
