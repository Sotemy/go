package main

import "fmt"
// format
func main() {
	fmt.Println("Hello World")
	var x int = 5
	fmt.Println(x)
	y := 5
	fmt.Println(y)
	z := "Hello World"
	fmt.Println(z)
	fmt.Println("Hello " + "World")
	fmt.Println("Hello World"[1])
	// function1()
	// function2()
	forCycle2()
}

var(
	a = 5
	b = 10
	c = 15
)

func function1() {
	fmt.Print("Enter number :")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)
}

func function2() {
	fmt.Print("Enter m:")
	var input float64
	fmt.Scanf("%f", &input)

	output:= input*0.3048
	fmt.Println(output)
}

func forCycle() {
	i:=1
	for i <= 10 {
		fmt.Println(i)
		i=i+1
	}
}

func forCycle2() {
	for i := 1; i<=10; i++ {
		if i % 2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}
}