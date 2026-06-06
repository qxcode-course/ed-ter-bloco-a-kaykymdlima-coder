package main
import "fmt"
func quadrante(matriz [][]rune, lin, col int) []rune {
    dim := len(matriz)
    l := (lin // dim) * dim
    c := (col // dim) * dim

    if dim == 4 {
        return []rune{
            matriz[l+0][c], matriz[l+0][c+1],
            matriz[l+1][c], matriz[l+1][c+1],
        }
    }

    if dim == 9 {
        return []rune{
            matriz[l+0][c], matriz[l+0][c+1], matriz[l+0][c+2],
            matriz[l+1][c], matriz[l+1][c+1], matriz[l+1][c+2],
            matriz[l+2][c], matriz[l+2][c+1], matriz[l+2][c+2],
        }
    }
    return nil
}
```

```go
func resolver(matriz [][]rune, index int) bool {
    nl := len(matriz)
    l := index / nl
    c := index % nl
    if index == nl * nl {
        return True
    }
    // se não for ponto, continue
    // para todos os números de [1 a N]
    //     se o número não estiver na linha, coluna e quadrante
    //         coloque o número na matriz
    //         se resolver(matriz, index + 1):
    //             return True
    //         matriz[l][c] = '.' // desfaz a tentativa
    // return False
}
```
