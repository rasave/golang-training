package main

import "fmt"

func main() {
	// declaring the variable which is an ASCII value of A
	var startingASCIINumber int = 64

	fmt.Println("Printing the alphabets from A to Z using a loop.")

	// printing the Alphabet from A to Z using for loop and
	// ASCII value concept
	for i := 1; i <= 26; i++ {
		fmt.Print(i)
		fmt.Print(i + 1)
		fmt.Print(string(rune(startingASCIINumber + i)))
		i++
		fmt.Print(string(rune(startingASCIINumber + i)))

		if i == 26 {
			fmt.Print(2728)
		}
	}
}
