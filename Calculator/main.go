package main

import (
	"Calculator/expression"
	"fmt"
	"math"
)

func main() {
	calc := expression.NewCalculator("(5 + 12 * (4 + (4 ^ (6 + 1))))")
	calc2 := expression.NewCalculator("(12.45 + 4)")
	calc3 := expression.NewCalculator("(6 + (9 * 4))")

	calc5 := expression.NewCalculator("((-12 + 8) * (-2))")

	fmt.Println(calc5.Evaluate())

	fmt.Println(calc.Evaluate())
	fmt.Println(calc2.Evaluate())
	fmt.Println(calc3.Evaluate())

	div := expression.NewExponentiation(expression.NewNumberSingleton(2), expression.NewNumberSingleton(-2))
	fmt.Println(div.Eval())

	fmt.Println(math.Pow(2.5, 2.5))
}
