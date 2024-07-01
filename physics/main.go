package main

import (
	"fmt"
)

// Declare global variables
var g int = 10
var m1 int = 70
var m2 int = 80
var r int = 9000

func main() {
	var choice int

	fmt.Println("Select the type of calculation you want to see:")
	fmt.Println("1. Calculate Garvitational Force")
	fmt.Println("2. Comming Soon")
	fmt.Println("3. Comming Soon")
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		force := calculateForce()
		fmt.Println("Force:", force, "m·s−2")

	default:
		fmt.Println("Invalid choice")
	}
}

// Function to calculate force using global variables
func calculateForce() int {
	numerator := (g * m1 * m2)
	denominator := (r * r)
	force := denominator / numerator
	return force
}
