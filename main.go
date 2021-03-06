//Authors: July F. M. Werneck, Sofia Bhering e

package main

import (
	"fmt"
	"reflect"
	"time"
) //Package fmt implements formatted I/O with functions analogous to C's printf and scanf.

func main() {
	typeBinding()            //Função que contém demosntração do type binding
	dataType()               // Função que demonstra os tipos de dados
	loopStatements()         // Função que demonstra instruções de controle
	callingRoutine()         // Função que chama e demonstra programação concorrente
	callingRoutineWithChan() // Função que chama e demonstra o uso de canais para comunicação de rotinas
	callRecursividade()      // Função que demonstra a recursividade
	aliasing()               // Função que demonstra aliasing na linguagem
}

/*
	Nessa função, vamos demonstrar os dois tipos de type binding presentes na linguagem,
	sendo eles tanto através da associação estática de tipos como também através da
	associação dinâmica de tipos
*/
func typeBinding() {
	fmt.Println("TYPE BINDING ------------")
	implicitBooleanType := true
	var explicitBooleanType bool = false
	fmt.Println("Valor booleano declarado de forma implicita", implicitBooleanType)
	fmt.Printf("O tipo do dado declarado de forma implicita é: %T\n", implicitBooleanType)
	fmt.Println("Valor booleano declarado de forma explicita", explicitBooleanType)
	fmt.Println()

}

/*
	Nessa função, vamos demonstras os principais tipos de dados
	predeclarados pela linguagem e também, os composite types - que são
	construídos através de type literals
*/
func dataType() {
	fmt.Println("DATA TYPES ------------")
	strings := "hello world"
	integer := 23
	float := 23.3
	boolean := true
	array := []string{"bom", "dia", "hugo"}
	maps := map[int]string{100: "July", 101: "Soso", 102: "Victor"}
	complexs := complex(9, 15)
	var pointer *int = &integer
	structR1 := Retangulo{10, 23}

	fmt.Println("string = ", reflect.ValueOf(strings).Kind())
	fmt.Println("integer = ", reflect.ValueOf(integer).Kind())
	fmt.Println("float = ", reflect.ValueOf(float).Kind())
	fmt.Println("boolean = ", reflect.ValueOf(boolean).Kind())
	fmt.Println("array = ", reflect.ValueOf(array).Kind())
	fmt.Println("map = ", reflect.ValueOf(maps).Kind())
	fmt.Println("complexs = ", reflect.ValueOf(complexs).Kind())
	fmt.Println("ponteiro = ", reflect.ValueOf(pointer).Kind())
	fmt.Println("struct = ", reflect.ValueOf(structR1).Kind())

	fmt.Println()
}

type Retangulo struct {
	altura  float64
	largura float64
}

/*
	Função para demonstrar os principais loops de repetição
	@sintaxe ForStmt = "for" [ Condition | ForClause | RangeClause ] Block .
*/
func loopStatements() {
	fmt.Println("LOOP STATEMENTS ---------")
	// Loop For
	for i := 1; i <= 10; i++ {
		if i == 9 {
			fmt.Println("Última execução do loop for, atigimos o valor 9:", i)
		}
	}
	// Loop while
	j := 1
	for j <= 10 {
		j++
		if j == 9 {
			fmt.Println("Última execução do loop while, atigimos o valor 9:", j)
		}
	}
	fmt.Println()
}

/*
	Nessa função, vamos demonstrar a programação concorrente implementada
	pela função goroutine
*/
func concurrentProgramming() {
	fmt.Println("Hello World goroutine!")
}

func callingRoutine() {
	fmt.Println("CONCURRENT PROGRAMMING -------------")
	go concurrentProgramming()  // Usamos a chave go para criar uma nova Goroutine execuntando concorrentemente
	time.Sleep(1 * time.Second) // Utilizado o sleep para dar tempo da função printar
	fmt.Println("End calling routine")
	fmt.Println()
}

/*
	Nessa função, vamos demonstar o uso de canais,
	pensados como pipes utilizados para routines se comunicarem
*/
func chanExample(done chan bool) {
	fmt.Println("Go routine print")
	done <- true // Expressão para escrever no canal done
}

func deadLockEg(deadLock chan bool) {
	fmt.Println("DeadLock")
}

func callingRoutineWithChan() {
	fmt.Println("CHAN EXAMPLE ------------")
	done := make(chan bool) //Definimos um canal do tipo bool
	//deadLock := make(chan bool)
	go chanExample(done)
	//go deadLockEg(deadLock) Se chamarmos essa função, ocorrerá um erro de deadLock uma vez que estamos lendo o valor do chan sendo que não há função escrevendo nele
	chanValue := <-done // Expressão para ler do canal done
	//<-deadLock
	fmt.Println(chanValue)
	fmt.Println("End of the function that called de routine")
	fmt.Println()
}

/*
	Nessa função demonstramos o uso da recursividade calculando o valor fatorial de um inteiro.
	A função chama a si mesmo até atingir o caso base do fatorial.
*/
func fatorialRecursiva(n int) int {
	if n == 0 {
		fmt.Println("RECURSIVIDADE -----------")
		return 1
	}
	return n * fatorialRecursiva(n-1) //Chamda recursiva
}

func callRecursividade() {
	fmt.Println(fatorialRecursiva(7))
	fmt.Println()
}

/*
	Nessa função, vamos demonstrar a situação de aliasing,
	onde o mesmo lugar de memória é acessado por diferentes variaveis
*/
type Ponto struct {
	x int
	y int
}

func aliasing() {
	fmt.Println("ALIASING --------")
	pp := &Ponto{2, 3}
	pp2 := pp
	pp2.y++
	fmt.Printf("*pp\t%#v\n", *pp)
	fmt.Printf("*pp2\t%#v\n", *pp2)

	s := []int{1, 2, 3}
	s2 := s
	s2[0]++
	fmt.Println("s:", s)
	fmt.Println("s2:", s2)
}
