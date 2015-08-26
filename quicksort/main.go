package main

import (
	"fmt"
	"os"
	"strconv"
)

var (
	rodada = 1
)

func main() {
	entrada := os.Args[1:]
	numeros := make([]int, len(entrada))
	for i, n := range entrada {
		numero, err := strconv.Atoi(n)
		if err != nil {
			fmt.Printf("%s não é um número válido\n", n)
		}
		numeros[i] = numero
	}
	fmt.Println(quicksort(numeros))
}

func quicksort(numeros []int) []int {
	fmt.Printf("############### Rodada %d #################\n", rodada)
	fmt.Printf("numeros iniciais: %v\n", numeros)
	rodada++
	if len(numeros) <= 1 {
		return numeros
	}

	n := make([]int, len(numeros))
	copy(n, numeros)
	indicePivo := len(n) / 2
	fmt.Printf("Indice Pivo: %d\n", indicePivo)
	pivo := n[indicePivo]
	fmt.Printf("Pivo: %d\n", pivo)
	n = append(n[:indicePivo], n[indicePivo+1:]...)
	fmt.Printf("Slice sem pivo: %v\n", n)
	menores, maiores := particionar(n, pivo)
	fmt.Printf("menores: %v\n", menores)
	fmt.Printf("maiores: %v\n", maiores)
	return append(
		append(quicksort(menores), pivo), quicksort(maiores)...)
}

func particionar(numeros []int, pivo int) (menores []int, maiores []int) {
	for _, n := range numeros {
		if n <= pivo {
			menores = append(menores, n)
		} else {
			maiores = append(maiores, n)
		}
	}
	return
}
