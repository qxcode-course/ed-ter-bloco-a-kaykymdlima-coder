package main

import "fmt"

func main() {
	var T, R int
	fmt.Scan(&T, &R)

	vetori := make([]int, T)
	for i := 0; i < T; i++ {
		fmt.Scan(&vetori[i])
	}
	vetorf := make([]int, T)
	for i := 0; i < T; i++ {
		alterna := (i + R) % T
		vetorf[alterna] = vetori[i]
	}
	fmt.Print("[ ")
	for _, vetorfp := range vetorf {
		fmt.Printf("%d ", vetorfp)
	}
	fmt.Println("]")
}
