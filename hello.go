package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
)

func main() {
	exibirNomes()
	for {
		exibirIntroducao()
		comando := lerComando()
		condicionalIfElse(comando)

		if comando == 0 {
			os.Exit(0)
		}
	}
}

func exibirIntroducao() {
	fmt.Println("Menu: ")
	fmt.Println("1 - Iniciar Monitoramento.")
	fmt.Println("2 - Exibir Logs.")
	fmt.Println("0 - Sair do Programa.")
}

func lerComando() int {
	comando := 0

	fmt.Scan(&comando)

	println("O comando digitado foi: ", comando)
	println("O ponteiro, o endereço da variável comando é: ", &comando)

	return comando
}

func condicionalIfElse(comando int) {
	comandoDesistencia := false
	if comando != 0 {
		if comando == 1 {
			iniciarMonitoramento(200)
		} else if comando == 2 {
			fmt.Println("Exibindo Logs...")
		} else if comando == 0 {
			fmt.Println("Saindo do Programa...")
			comandoDesistencia = true
			os.Exit(0)
		} else {
			fmt.Println("Opção Invalida.")
			comandoDesistencia = true
			os.Exit(-1)
		}
	} else {
		if comandoDesistencia == true {
			os.Exit(0)
		}
		fmt.Println("Opção Invalida.")
	}
}

func condicionalSwitch(comando int) {
	switch comando {
	case 1:
		iniciarMonitoramento(200)
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do Programa...")
	default:
		fmt.Println("Opção Invalida.")
	}
}

func iniciarMonitoramento(statusCode int) {
	if statusCode == 0 {
		statusCode = 200
	}
	sites := [...]string{fmt.Sprintf("https://httpbin.org/status/%d", statusCode), "https://www.alura.com.br/", "https://www.xvideos.com/", "https://www.uol.com.br/"}
	fmt.Println("Iniciando Monitoramento...")

	for i := 0; i < len(sites); i++ {
		println(sites[i])
	}

	site := sites[0]

	response, err := http.Get(site)

	if response.StatusCode == 200 {
		fmt.Printf("O status do site: %d", response.StatusCode)
	} else {
		fmt.Printf("O status do site: %d, o error: %s", response.StatusCode, err)
	}
	fmt.Println("O status: ", response.Status)
}

func exibirNomes() {
	nomesArray := [3]string{"Ramon", "Milla", "Valeriano"}

	nomesSlice := []string{"Ramon", "Milla", "Gabriela", "Valeriano"}

	fmt.Printf("\nNomes Arrays: %v - Quantidade de elementos: %d - Capacidade: %d", nomesArray, len(nomesArray), cap(nomesArray))

	fmt.Printf("\nNomes Slices: %v - Quantidade de elementos: %d - Capacidade: %d", nomesSlice, len(nomesSlice), cap(nomesSlice))

	nomesSlice = append(nomesSlice, "Gael")

	fmt.Printf("\nNomes Slices: %v - Quantidade de elementos: %d - Capacidade: %d", nomesSlice, len(nomesSlice), cap(nomesSlice))

	nomesSlice = append(nomesSlice, "Dante", "Rayan")

	fmt.Printf("\nNomes Slices: %v - Quantidade de elementos: %d - Capacidade: %d\n", nomesSlice, len(nomesSlice), cap(nomesSlice))

	for indice, valor := range nomesSlice {
		fmt.Printf("Inidice: %d - Valor: %s\n", indice, valor)
	}

	for _, value := range nomesArray {
		fmt.Printf("Value: %s\n", value)
	}
}

func textosIniciaisDeAprendizado() {
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

	// Em muitos casos você não precisa declarar qual o tipo, o Go saberá interpretar qual o tipo de cada variável.
	var nome_usuario = "Valeriano"
	var idade_nova = 35
	var versao_nova = 1.2

	fmt.Println("O seu nome é: ", nome_usuario, "Tem a idade: ", idade_nova)
	fmt.Println("A versão: ", versao_nova)

	// Para saber quais o tipos de variáveis.
	fmt.Println("Tipo do nome: ", reflect.TypeOf(nome_usuario))
	fmt.Println("Tipo da idade: ", reflect.TypeOf(idade_nova))
	fmt.Println("Tipo de versão: ", reflect.TypeOf(versao_nova))

	// Atribuição de variaveis curto, no lugar de declarar "var"

	nome_dela := "Milla"
	idade_dela := 36

	fmt.Println("O nome dela: ", nome_dela, "A idade dela: ", idade_dela)

	fmt.Println("Olá Sr. ", nome_do_usuario)
	fmt.Println("O programa está na versão: ", versao_float)
}
