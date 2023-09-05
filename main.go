package main

import (
	"fmt"

	"github.com/rajvinoth227/mypackage/math"
	"github.com/rajvinoth227/mypackage/strings"
	"github.com/rajvinoth227/mypackage/utils"
)

func main() {
	// Math Package
	result1 := math.Add(5, 3)
	result2 := math.Subtract(10, 4)
	fmt.Printf("Math Add: %d\n", result1)
	fmt.Printf("Math Subtract: %d\n", result2)

	// Strings Package
	concatResult := strings.Concat("Hello, ", "world!")
	reverseResult := strings.Reverse("Hello, world!")
	fmt.Printf("Strings Concat: %s\n", concatResult)
	fmt.Printf("Strings Reverse: %s\n", reverseResult)

	// Utils Package
	isEven := utils.IsEven(6)
	isValidEmail := utils.IsValidEmail("example@email.com")
	fmt.Printf("IsEven: %v\n", isEven)
	fmt.Printf("IsValidEmail: %v\n", isValidEmail)
}
