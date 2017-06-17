// Author: Cauê Garcia Polimanti
// Date: 17/06/2017

package main

import "fmt"

// Show usage instructions
func showUsageInstructions() {
	fmt.Println("Modo de usar:\n\n\t tweetstorm.exe [texto] \n\n")
	showHelp()
}

// Show help instructions
func showHelp() {
	fmt.Println("Digite o texto que deseja converter como primeiro parâmetro.\n\nExemplo:\n\t tweetstorm.exe \"Este é um texto exemplo de entrada que será convertido para Tweets separados com até 140 caracteres cada um. Esta parte é só para ultrapassar os 140 caracteres\"")
}
