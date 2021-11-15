package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main()  {
	firstvalue := getInput("first value")

	secondvalue := getInput("second vaule")

	var output float64

	switch  calculate := getOperation(); calculate {
		case "+":
			output = addValues(firstvalue, secondvalue)
		case "*":
			output = multiplication(firstvalue,secondvalue)
		case "-":
			output = substraction(firstvalue,secondvalue)
		case "/":
			output = division(firstvalue,secondvalue)
		default:
			panic("Invalid operation")	
	}
	fmt.Printf("The result is =  %v\n\n", output)
}

func division(firstvalue, secondvalue float64) float64 {
   return firstvalue / secondvalue
}

func substraction(firstvalue, secondvalue float64) float64{
	return firstvalue - secondvalue
}

func multiplication(firstvalue, secondvalue float64) float64{
	return firstvalue* secondvalue
}

func addValues(firstvalue, secondvalue float64) float64{
	return firstvalue + secondvalue
}

func getInput(prompt string) float64 {
	fmt.Printf("%v: ", prompt)
	input,_ := reader.ReadString('\n')

	value, err := strconv.ParseFloat(strings.TrimSpace(input),64)

	if err != nil {
		message := fmt.Sprintf("%v must be a number", prompt)
		panic(message)
	}
	return value
}

func getOperation() string {
	fmt.Printf("select an operation (+ - * /): ")
	input,_ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
