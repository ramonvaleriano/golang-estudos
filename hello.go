package main

import "fmt"

func main() {
	fmt.Println("Olá mundo, aprendendo Golang.")

	var nome_do_usuario string = "Ramon Valeriano"
	var versao_string string = "1.0"
	var versao_float float32 = 1.1
	var idade int
	// var varia_nao_usade int // Variável não usada, em go não se permite declarar e não usar.

	fmt.Println("Nome do Usuário: ", nome_do_usuario)
	fmt.Println("Estamos usando na versão: ", versao_string)
	fmt.Println("Versão em float: ", versao_float)
	fmt.Println("A idade é: ", idade)
}
