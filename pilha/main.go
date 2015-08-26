package main

import (
	"errors"
	"fmt"
)

// Pilha is a struct with valores, the type of valores is []interface{}
type Pilha struct {
	valores []interface{}
}

// Tamanho is a method to verify the size of the one Pilha
func (p Pilha) Tamanho() int {
	return len(p.valores)
}

// Vazia is a method to verify if one Pilha is empty
func (p Pilha) Vazia() bool {
	return p.Tamanho() == 0
}

// Empilhar insert values on the Pilha
func (p *Pilha) Empilhar(valor interface{}) {
	p.valores = append(p.valores, valor)
}

// Desempilhar remove valores from Pilha
func (p *Pilha) Desempilhar() (interface{}, error) {
	if p.Vazia() {
		return nil, errors.New("Pilha vazia")
	}
	valor := p.valores[p.Tamanho()-1]
	p.valores = p.valores[:p.Tamanho()-1]
	return valor, nil
}

func main() {
	pilha := Pilha{}
	fmt.Println("Pilha criada com tamanho", pilha.Tamanho())
	fmt.Println("Vazia?", pilha.Vazia())
	pilha.Empilhar("Go")
	pilha.Empilhar(2009)
	pilha.Empilhar(3.14)
	pilha.Empilhar("Fim")
	fmt.Println("Tamanho da pilha após empilhar quatro valores:", pilha.Tamanho())
	fmt.Println("Vazia?", pilha.Vazia())

	for !pilha.Vazia() {
		v, _ := pilha.Desempilhar()
		fmt.Println("Desempilhando ", v)
		fmt.Println("Tamanho:", pilha.Tamanho())
		fmt.Println("Vazia?", pilha.Vazia())
	}

	_, err := pilha.Desempilhar()
	if err != nil {
		fmt.Println("Erro:", err)
	}
}
