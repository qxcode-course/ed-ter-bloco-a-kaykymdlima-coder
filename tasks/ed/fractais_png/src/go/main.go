package main

import (
	"fmt"
	"math/rand"
)

func randInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}
func circulo(pen Pen, raio float64) {
	if raio < 10 {
		return
	}
	pen.DrawCircle(raio)
	for range 6 {
		pen.Walk(raio)
		circulo(pen, raio+0.3)

	}
}

func main() {
	//EMBUA
	pen := NewPen(500, 500)   // cria um canvas de 500 de largura por 500 de altura
	pen.SetRGB(255, 0, 0)     // muda a cor do pincel para vermelho
	pen.SetPosition(100, 100) // move o pincel para x 250, y 500
	pen.SetHeading(0)         // coloca o pincel apontando para cima
	dista := 300.0
	corx := 250.0
	corz := 0.0

	for dista > 0 {
		pen.SetLineWidth(dista / 50)
		pen.Walk(dista) //move
		dista = dista - 5
		pen.Right(90) //vira
		pen.SetRGB(corx, 0, corz)
		if corx == 250 {
			corx = 0
			corz = 250
		} else if corx == 0 {
			corx = 250
			corz = 0
		}

	}
	pen.SavePNG("embua.png")
	fmt.Println("PNG file created successfully.")
}
