package main

import "fmt"

func main() {
	var talb, tfig int
	fmt.Scan(&talb, &tfig)
	album := make([]int, talb+1)
	repet := make([]int, tfig)
	var fig int
	figant := -1
	qrepet := 0

	for i := 0; i < tfig; i++ {
		fmt.Scan(&fig)
		album[fig] = 1
		if fig == figant {
			repet[qrepet] = fig
			qrepet++
		}
		figant = fig
	}

	if qrepet == 0 {
		fmt.Println("N")
	} else {
		for i := 0; i < qrepet; i++ {
			fmt.Printf("%d", repet[i])
			if i < qrepet-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
	falta := 0
	pri := 0
	for i := 1; i <= talb; i++ {
		if album[i] == 0 {
			if pri == 1 {
				fmt.Print(" ")
			}
			fmt.Print(i)
			falta = 1
			pri = 1
		}
	}

	if falta == 0 {
		fmt.Println("N")
	} else {
		fmt.Println()
	}
}
