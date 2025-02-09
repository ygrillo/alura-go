package main

import "fmt"

func main() {
	nome := "Yuri"
	versao := 1.1
	fmt.Println("Olá, Sr.", nome)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir os logs")
	fmt.Println("0- Sair do programa")

	var comando int
	fmt.Scan(&comando)
	fmt.Println("O endereço da variável comando é", &comando)
	fmt.Println("O comando escolhi foi", comando)
}
