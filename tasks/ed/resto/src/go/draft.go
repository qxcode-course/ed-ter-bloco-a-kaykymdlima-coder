package main

import "fmt"

func restoeprint(num int) {
	if num == 0 {
		return
	}
	resu := num / 2
	rest := num % 2
	restoeprint(resu)
	fmt.Printf("%d %d\n", resu, rest)
}
func main() {
	var num int
	fmt.Scan(&num)
	restoeprint(num)
}
