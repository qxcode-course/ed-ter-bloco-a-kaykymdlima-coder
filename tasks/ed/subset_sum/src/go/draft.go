package main

import "fmt"

func dahsoma(list []int, alvo int) bool {
	var busca func(int, int) bool
	busca = func(index, soma int) bool {
		if soma == alvo {
			return true
		} else if index == len(list) || soma > alvo {
			return false
		}
		return busca(index+1, soma+list[index]) || busca(index+1, soma)
	}
	return busca(0, 0)
}
func main() {
	var n, k int
	fmt.Scan(&n, &k)
	list := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&list[i])
	}
	if dahsoma(list, k) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
