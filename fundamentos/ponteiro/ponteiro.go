package main

import "fmt"

func main() {
	// Go nao tem aritmetica de ponteiros
	// p++

	i := 1

	var p *int = nil
	p = &i // pegando o endereco da variavel
	*p++   // desreferenciando (pegando o valor)
	i++

	fmt.Println(p, *p, i, &i)
}
