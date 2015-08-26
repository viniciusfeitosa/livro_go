package main

import (
	"fmt"
)

type ListaDeCompras []string

func imprimir(slice []string) {
	fmt.Println("Slice:", slice)
}

func imprimirLista(lista ListaDeCompras) {
	fmt.Println("Lista de Compras:", lista)
}

func main() {
	lista := ListaDeCompras{"Alface", "Atum", "Azeite"}
	slice := []string{"Alface", "Atum", "Azeite"}
	imprimir([]string(lista))
	imprimirLista(ListaDeCompras(slice))
}
