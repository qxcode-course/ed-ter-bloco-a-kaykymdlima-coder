package main

import (
	"fmt"
	"math/rand"
)

func ri(inf, sup int) float64 {
	return float64(rand.Intn(sup-inf+1) + inf)
}
func circulo(pen *Pen, raio float64) {
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
}

func arvore(pen *Pen, dista float64) {
	if dista < 10 {
		if ri(0, 200) == 0 {
			pen.SetRGB(255, 0, 0)
			pen.FillCircle(10)
		}
		return
	}

	angDI := ri(20, 26)
	angES := ri(20, 26)
	pen.SetLineWidth(dista / 10)
	pen.SetRGB(101, 67, 33)
	pen.Walk(dista)
	pen.Right(angDI)
	pen.SetRGB(0, 67, 0)
	arvore(pen, dista*(ri(80, 85)/100))
	pen.Right(-(angDI + angES))
	arvore(pen, dista*(ri(80, 85)/100))
	pen.Right(angES)
	pen.Walk(-dista)

}

func fractal(pen *Pen, rad float64) {
	if rad < 10 {
		return
	}
}

func main() {
	pen := NewPen(500, 500)   // cria um canvas de 500 de largura por 500 de altura
	pen.SetHeading(90)        // coloca o pincel apontando para cima
	pen.SetPosition(250, 250) // move o pincel para x 250, y 500

	dista := 50.0
	//embua(pen, dista)
	fractal(pen, dista)
	pen.SavePNG("floconeve.png")
	fmt.Println("PNG file created successfully.")
}
