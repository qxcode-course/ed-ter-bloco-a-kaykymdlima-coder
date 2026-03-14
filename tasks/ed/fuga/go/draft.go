package main

import "fmt"

func main() {
	var H, P, F, D int
	fmt.Scan(&H, &P, &F, &D)
	if D == 1 {
		for F <= 16 {
			F = F + 1
			if F == 16 {
				F = 0
			}
			if F == P {
				fmt.Print("N\n")
				break
			} else if F == H {
				fmt.Print("S\n")
				break
			}
		}
	} else if D == -1 {
		for F < 16 {
			F = F - 1
			if F == -1 {
				F = 15
			}
			if F == P {
				fmt.Print("N\n")
				break
			} else if F == H {
				fmt.Print("S\n")
				break
			}
		}
	}
}
