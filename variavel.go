package main

import "fmt"
import "reflect"

func main() {

	var precoLeite float32 //declarando variável sem valor, default = 0
	var precoOvo float32 = 3.99
	var precoPao = 0.99 //Se quiser também podemos omitir o tipo da variável que pelo valor atribuído o "Go" entendi o Tipo
	precoFarinha := 2.30

	//retorno de multiplus valores
	precoOvo, precoPao = returnEggValueAndBreadValue()

	//se você quiser ignorar um dos retornos use _
	_, precoPao = returnEggValueAndBreadValue()

	fmt.Println("O valor do Leite é R$", precoLeite)
	fmt.Println("O valor do Ovo é R$", precoOvo)
	fmt.Println("O valor do Pão é R$", reflect.TypeOf(precoPao)) //Para saber o tipo da variável com reflect.TypeOf()
	fmt.Println("O valor da Farinha é R$", precoFarinha)
}

func returnEggValueAndBreadValue() (float32, float64) {
	var precoOvo float32 = 8.99
	var precoPao = 0.87
	return precoOvo, precoPao
}
