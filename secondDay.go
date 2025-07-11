/* Problem Statement:
Create a program that calculates:

The total trip cost (based on fuel, food, and accommodation expenses)

The cost per person (divide total cost by number of travelers)

The percentage each category (fuel, food, stay) contributes to the total cost */

package main

import (
	"fmt"
)

func main() {

	fuel := requirements("Fule cost?_")
	food := requirements("food expenses?_")
	accommodation := requirements("cost of accomodation?_")

	numPerson := requirements("What is the total number of people on the trip?_")

	totalTripCost, costPerPerson, fuelPercentage, foodPercentage, accommodationPercentage := calculations(fuel, food, accommodation, numPerson)

	fmt.Println("The total trip cost is:", totalTripCost)
	fmt.Println("The cost per person is:", costPerPerson)
	fmt.Printf("The total trip cost is: %.2f%%, %.2f%%, %.2f%% respectively", fuelPercentage, foodPercentage, accommodationPercentage)

}

func requirements(askuser string) float64 {

	var userinput float64
	fmt.Print(askuser)
	fmt.Scan(&userinput)
	return userinput

}

func calculations(fuel, food, accommodation, numPerson float64) (float64, float64, float64, float64, float64) {

	totalTripCost := fuel + food + accommodation
	costPerPerson := totalTripCost / numPerson
	fuelPercentage := (fuel / totalTripCost) * 100
	foodPercentage := (food / totalTripCost) * 100
	accommodationPercentage := (accommodation / totalTripCost) * 100

	return totalTripCost, costPerPerson, fuelPercentage, foodPercentage, accommodationPercentage

}
