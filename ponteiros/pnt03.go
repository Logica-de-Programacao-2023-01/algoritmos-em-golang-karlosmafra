package main

import "fmt"

type Produto struct {
	nome  string
	preco float64
	quant int
}

func main() {
	p := Produto{
		nome:  "Shampoo",
		preco: 20.5,
		quant: 40,
	}

	var ptr *Produto = &p

	fmt.Println(p.nome)
	fmt.Println("Preço:", p.preco)
	fmt.Println("Quantidade:", p.quant)

	var prc float64
	fmt.Print("Novo preço: ")
	fmt.Scan(&prc)

	atualizar(ptr, prc)
	fmt.Println("Novo preço:", p.preco)
}

func atualizar(ptr *Produto, prc float64) {
	ptr.preco = prc
}
