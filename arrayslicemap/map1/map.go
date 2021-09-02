package main

import "fmt"

func main() {
	// map devem ser inicializados

	aprovados := make(map[int]string)

	aprovados[123456789789] = "Maria"
	aprovados[234239930203] = "Pedro"
	aprovados[837938233029] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[837938233029])
	delete(aprovados, 837938233029)
	fmt.Println(aprovados[837938233029])
}
