package main

import (
	"fmt"
	"math/rand"
)

func ri(min, max int) int {
	return float64 (rand. intn(

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
func embua(pen *Pen, dista float64) {
	if dista < 0 {
		return
	}
	pen.SetLineWidth(dista / 50)
	pen.SetRGB(255, 0, 0) // muda a cor do pincel para vermelho
	pen.Walk(dista)       //move
	pen.SavePNG("embua.png")
	var dummy rune
	fmt.Scanf("%c", &dummy)
	dista = dista - 5
	pen.Right(90) //vira
	embua(pen, dista)
	/*
			corx := 250.0
			corz := 0.0
		pen.SetRGB(corx, 0, corz)
			if corx == 250 {
				corx = 0
				corz = 250
			} else if corx == 0 {
				corx = 250
				corz = 0
			}

		}
	*/
}

func arvore(pen *Pen, dista float64) {
	if dista < 10 {
		return
	}
	ang := 25.0
	fator := 0.8
	pen.SetLineWidth(dista / 10)
	pen.Walk(dista)
	pen.Right(ang)
	arvore(pen, dista*(ri(80,85)/100))
	pen.Right(-2 * ang)
	arvore(pen, dista*fator)
	pen.Right(ang)
	pen.Walk(-dista)

}

func main() {
	pen := NewPen(500, 500)   // cria um canvas de 500 de largura por 500 de altura
	pen.SetHeading(90)        // coloca o pincel apontando para cima
	pen.SetPosition(250, 500) // move o pincel para x 250, y 500

	dista := 90.0
	//embua(pen, dista)
	arvore(pen, dista)
	pen.SavePNG("tree.png")
	fmt.Println("PNG file created successfully.")
}
