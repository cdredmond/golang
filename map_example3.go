package main

import "fmt"

func main(){
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

	fmt.Print("Enter element symbol: ")
	var input string
	fmt.Scanf("%s", &input)

	if name, ok := elements[input]; ok {
		fmt.Println(name, ok)
	}
}

