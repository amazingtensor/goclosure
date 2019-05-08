package main

import "fmt"
import "strings"

func stringInSlice(a string, list string) bool {
	if strings.Contains(list, a){
		fmt.Println("yes " + a + " in "  + list)
		return true
	}
	return false
}

func main() {
	fmt.Println("Hello World!")
	var X = "AB"
	var Xp = X
	fmt.Println("X+ starts as " + Xp)
	var fds = map[string]string {
		"AB": "C",
		"BC": "AD",
		"D": "E",
		"CF": "B",
	}

	fmt.Print("Functional dependencies are {")
	for W, Z := range fds {
		fmt.Print(W + " -> " + Z + ", ")
	}
	fmt.Println("}")

	var stillChanging = true

	for stillChanging {
		stillChanging = false
		for W, Z := range fds {
			fmt.Println("Checking dependency " + W + " -> " + Z)
			if stringInSlice(W, Xp) && !stringInSlice(Z, Xp) {
				for _,r := range Z {
					if !strings.Contains(Xp,string(r)) {
						Xp = Xp + string(r)
						stillChanging = true
					}
				}
			}
		}
	}

	fmt.Println("Final X+ is " + Xp)
}
