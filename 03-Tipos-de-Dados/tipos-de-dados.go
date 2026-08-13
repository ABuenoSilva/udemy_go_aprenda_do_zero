package main

import (
	"errors"
	"fmt"
)

func main() {
	// Tipos inteiros: int8, int16, int32, int64
	var numero int16 = 100
	// Tipo int utiliza a arquitetura do computador, e também na inferência de tipos
	numero1 := -100000000

	fmt.Println(numero, numero1)

	// Tipo inteiro sem sinal (unsigned): uint
	var numero2 uint32 = 10000
	fmt.Println(numero2)

	// Existem "aliases" para alguns tipo. rune = int32, byte = uint8
	// Geralmente usados para armazenar valores ASCII
	var numero3 rune = -10000
	var numero4 byte = 65

	fmt.Println(numero3, numero4)

	// Tipos reais = float --> float32 e float64
	// Inferência também fica como somente float (não dá pra declarar com esse tipo) e obedece a arquitetura do computador
	var numero5 float32 = 1234.56
	var numero6 float64 = 123123123.7789

	fmt.Println(numero5, numero6)

	// Tipo string
	var str string = "texto2"
	fmt.Println(str)

	// Tipo char não existe, ele é convertido para o número da tabela ASCII (tipo int)
	char := 'B'
	fmt.Println(char)

	// Valor 0 (específico do go) -> Variáveis declaradas mas não inicializadas recebem branco (string) ou zero(números)
	var texto string
	var number int8
	fmt.Println(texto, number)

	// Tipo boolean, o tipo zero é false
	var boolean1 bool = false
	boolean2 := true
	var boolean3 bool
	fmt.Println(boolean1, boolean2, boolean3)

	// Tipo erro (erro em Go é um tipo), tipo zero é nulo <nil>
	var erro error
	var error2 error = errors.New("Erro interno")
	fmt.Println(erro)
	fmt.Println(error2)
}
