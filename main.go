package main

import "fmt"

// this is a comment

func main(){
	fmt.Println("Hello World")

	fmt.Println("1 + 1 =", 1 + 1)

	fmt.Print("Enter a number: ")
	
	var input float64

	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)

	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}
}

