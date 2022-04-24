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
	// forCycle2()
	// arrays3()
	maps3()
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

func arrays1() {
	var x [5]int
	x[4]=100
	fmt.Println(x)
}

func arrays2() {
	var x[5]float64
	x[0]=98
	x[1]=93
	x[2]=77
	x[3]=82
	x[4]=83

	var total float64=0
	for i:=0; i <5 ; i++ {
		total+=x[i]
	}
	fmt.Println(total/float64(len(x)))
}

func arrays3() {
	var x[5]float64
	x[0]=98
	x[1]=93
	x[2]=77
	x[3]=82
	x[4]=83

	x=[5]float64{98, 93, 77, 82, 83,}

	var total float64 =0
	for _, value := range x {
		total += value
	}
	fmt.Println(total/float64(len(x)))
	arr := [5]float64{1,2,3,4,5}
	y := arr[0:5]
	fmt.Println(arr)
	fmt.Println(y)
}

func maps1() {
	// var x map[string]int
	x := make(map[int]int)
	x[1]=10
	x[2]=2
	fmt.Println(x)
	delete(x, 1)
	fmt.Println(x)
}

func maps2() {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"
	elements = map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}
	
	fmt.Println(elements["Li"])

	name, ok := elements["Un"]
	fmt.Println(name, ok)
}

func maps3() {
	elements := map[string]map[string]string{
		"H": map[string]string{
		  "name":"Hydrogen",
		  "state":"gas",
		},
		"He": map[string]string{
		  "name":"Helium",
		  "state":"gas",
		},
		"Li": map[string]string{
		  "name":"Lithium",
		  "state":"solid",
		},
		"Be": map[string]string{
		  "name":"Beryllium",
		  "state":"solid",
		},
		"B":  map[string]string{
		  "name":"Boron",
		  "state":"solid",
		},
		"C":  map[string]string{
		  "name":"Carbon",
		  "state":"solid",
		},
		"N":  map[string]string{
		  "name":"Nitrogen",
		  "state":"gas",
		},
		"O":  map[string]string{
		  "name":"Oxygen",
		  "state":"gas",
		},
		"F":  map[string]string{
		  "name":"Fluorine",
		  "state":"gas",
		},
		"Ne":  map[string]string{
		  "name":"Neon",
		  "state":"gas",
		},
	}

	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}