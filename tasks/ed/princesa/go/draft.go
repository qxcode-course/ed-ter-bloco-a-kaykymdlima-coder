package main

import "fmt"

func mata(vivos []int, morto int) []int {
	for i := morto; i < len(vivos)-1; i++ {
		vivos[i] = vivos[i+1]
	}
	return vivos[:len(vivos)-1]
}
func printVivos(vivos []int, espada int) {
	fmt.Print("[ ")
	for i := 0; i < len(vivos); i++ {
		if i == espada {
			fmt.Printf("%d> ", vivos[i])
		} else {
			fmt.Printf("%d ", vivos[i])
		}
	}
	fmt.Println("]")
}
func main() {
	var n, e int
	if _, err := fmt.Scan(&n, &e); err != nil {
		return
	}

	vivos := make([]int, n)
	for i := 0; i < n; i++ {
		vivos[i] = i + 1
	}

	espada := e - 1

	printVivos(vivos, espada)

	for len(vivos) > 1 {
		proxmort := (espada + 1) % len(vivos)

		vivos = mata(vivos, proxmort)

		espada = proxmort % len(vivos)

		printVivos(vivos, espada)
	}
}
