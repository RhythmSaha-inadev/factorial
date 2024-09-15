// Write a program to find out the factorial of a number using function and user input

package main

import "fmt"

func factorial(a *int) {
	fact := 1
	for i := 1; i <= *a; i++ {
		fact *= i //fact=fact*i
	}
	fmt.Println("Factorial is", fact)
}

func main() {
	var a int
	fmt.Println("Enter number")
	fmt.Scan(&a)

	factorial(&a)
}
