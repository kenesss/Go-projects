package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nSimple CLI Calculator")
		fmt.Println("Choose an operation: +  -  *  /  sqrt  pow")
		fmt.Print("Enter operation: ")
		opInput, _ := reader.ReadString('\n')
		opInput = strings.TrimSpace(opInput)

		var num1, num2 float64
		var err error

		if opInput != "sqrt" {
			fmt.Print("Enter first number: ")
			num1, err = getFloatInput(reader)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number.")
				continue
			}

			fmt.Print("Enter second number: ")
			num2, err = getFloatInput(reader)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number.")
				continue
			}
		} else {
			fmt.Print("Enter a number: ")
			num1, err = getFloatInput(reader)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number.")
				continue
			}
		}

		switch opInput {
		case "+":
			fmt.Printf("Result: %.2f\n", num1+num2)
		case "-":
			fmt.Printf("Result: %.2f\n", num1-num2)
		case "*":
			fmt.Printf("Result: %.2f\n", num1*num2)
		case "/":
			if num2 == 0 {
				fmt.Println("Error: Division by zero is not allowed.")
			} else {
				fmt.Printf("Result: %.2f\n", num1/num2)
			}
		case "sqrt":
			if num1 < 0 {
				fmt.Println("Error: Cannot calculate square root of a negative number.")
			} else {
				fmt.Printf("Result: %.2f\n", math.Sqrt(num1))
			}
		case "pow":
			fmt.Printf("Result: %.2f\n", math.Pow(num1, num2))
		default:
			fmt.Println("Invalid operation. Please choose +, -, *, /, sqrt, or pow.")
		}

		fmt.Print("Do you want to perform another calculation? (y/n): ")
		again, _ := reader.ReadString('\n')
		again = strings.TrimSpace(again)
		if strings.ToLower(again) != "y" {
			break
		}
	}
}

func getFloatInput(reader *bufio.Reader) (float64, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	return strconv.ParseFloat(input, 64)
}
