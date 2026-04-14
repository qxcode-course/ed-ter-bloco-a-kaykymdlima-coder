package main

import "fmt"

func main() {
	var quantini int
	fmt.Scan(&quantini)
	fila := make([]int, quantini)
	identificador := 0
	for i := 0; i < quantini; i++ {
		fmt.Scan(&identificador)
		fila[i] = identificador
	}
	var quantsairam int
	fmt.Scan(&quantsairam)
	/*esqueci a identação do map, :(, ent vou fzer com vetores)
	sairam := map[int]int
	*/
	sairam := make([]int, quantsairam)
	saiu := 0
	for i := 0; i < quantsairam; i++ {
		fmt.Scan(&saiu)
		sairam[i] = saiu
	}
	quantfinal := quantini - quantsairam
	filafinal := make([]int, quantfinal)

	k := 0
	for i := 0; i < quantini; i++ {
		for j := 0; j < quantsairam; j++ {
			if fila[i] == sairam[j] {
				filafinal[k] = fila[i]
				k++
			}
		}
	}
	for i := 0; i < quantfinal; i++ {
		fmt.Printf("%d", filafinal[i])
	}

}
