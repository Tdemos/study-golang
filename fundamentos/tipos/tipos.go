package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// numeros inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro e", reflect.TypeOf(32000))

	// sem sinal (so positivos) ... uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte e", reflect.TypeOf(b))

	//com sinal... int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor maximo do int e", i1)
	fmt.Println("O pito de i1 e", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela Unicode (int32)
	fmt.Println("O rune e", reflect.TypeOf(i2))
	fmt.Println(i2)

	// numeros reais(float32, flaot64)
	var x float32 = 49.99
	fmt.Println("O tipo de x e", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 e", reflect.TypeOf(49.99))

	//boolean
	bo := true
	fmt.Println("O tipo de bo e", reflect.TypeOf(bo))
	fmt.Print(!bo)

	//string
	s1 := "Ola meu nome e ---"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string s1 e", len(s1))

	// string com multiplas linhas
	s2 := `Ola
	meu
	nome
	e 
	---
	`
	fmt.Println("O tamanho da string s2 3", len(s2))
}
