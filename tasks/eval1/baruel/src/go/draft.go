package main

import "fmt"

func main() {
	var talbum, tfigu int
	fmt.Scan(&talbum, &tfigu)

	baruel := make([]int, tfigu)
	repetidas := make([]int, tfigu)
	quantrepet := 0
	var atualfig int
	anterfig := 0
	for i := 0; i < tfigu; i++ {
		fmt.Scan(atualfig)
		if anterfig == atualfig {
			repetidas[quantrepet] = atualfig
			quantrepet++
		}
		baruel[i] = atualfig
		anterfig = atualfig
	}
	if quantrepet == 0 {
		fmt.Println("N")
	} else {
		for i := 0; i < quantrepet; i++ {
			fmt.Printf("%d ", repetidas[i])
		}
	}
	fmt.Println("")

	temnoalbum := 1
	falta := 0
	for i := 0; i < tfigu; i++ {
		if baruel[i] == temnoalbum {
			temnoalbum++
			falta = 1
			continue
		} else {
			fmt.Printf("%d ", baruel[i])
			temnoalbum++
		}
	}
	if falta == 0 {
		fmt.Println("N")
	}
}
