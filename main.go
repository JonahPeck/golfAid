package main

import (
	"fmt"
)

func main() {
	fmt.Println("What is the Yardage?")
	var yardage int
	var club string
	// var hazards string
	var lie string
	var lieSuggestion string

	fmt.Scanln(&yardage)

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
	// fmt.Println("You should hit your", club)
	fmt.Println("How is the lie? (if you are teeing off simply respond with 'good' for a reminder of the fundamentals of a proper swing)")
	fmt.Println("Responses can be: tight, flier, jumpy, uphill, downhill, below feet, above feet, plugged, burried or good")
	fmt.Scanln(&lie)

	if lie == "tight" {
		lieSuggestion = "Focus on hitting down on the ball and focus on getting contact with the ball first then making a small divot upon followthrough"
	} else if lie == "flyer" || lie == "jumpy" {
		lieSuggestion = "To counteract the 'flier' take less club, choke down, and make a three-quarter swing."
	} else if lie == "uphill" {
		lieSuggestion = "To hit an uphill lie, take more club than you normally would (8 iron instead of 9 iron). Aim right of the target as the ball tends to travel to the left. Don't plan for much roll as the higher trajectory will change that. Place the clubhead behind the ball and match your shoulders with the slope of the lie. Swing with the slope of the lie."
	} else if lie == "downhill" {
		lieSuggestion = "downhill"
	} else if lie == "below feet" {
		lieSuggestion = "below feet"
	} else if lie == "above feet" {
		lieSuggestion = "above feet"
	} else if lie == "plugged" || lie == "burried" {
		lieSuggestion = "plugged or burried"
	} else if lie == "good" {
		lieSuggestion = "good"
	} else {
		fmt.Println("Please select a lie that is listed: ")
	}

	//what hazards are ahead, what lie is it in, how is your confidence right now
	fmt.Println("Based on your Yardage, we suggest you hit your", club)
	fmt.Println(lieSuggestion)
}
