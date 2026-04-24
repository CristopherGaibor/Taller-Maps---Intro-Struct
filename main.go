package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Ingrese texto: ")
	scanner.Scan()
	texto := scanner.Text()

	palabras := strings.Fields(strings.ToLower(texto))
	conteo := make(map[string]int)

	for _, p := range palabras {
		conteo[p]++
	}

	var masRepetida string
	max := 0

	fmt.Println("\nResultados:")
	for p, c := range conteo {
		fmt.Printf("%s: %d\n", p, c)
		if c > max {
			max = c
			masRepetida = p
		}
	}

	fmt.Printf("\nPalabra más repetida: %s (%d veces)\n", masRepetida, max)
}
