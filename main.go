//Authors: July F. M. Werneck, Sofia Bhering e

package main

import (
	"fmt"
	"reflect"
) //Package fmt implements formatted I/O with functions analogous to C's printf and scanf.

func main() {
	typeBinding()    //Função que contém demosntração do type binding
	dataType()       // Função que demonstra os tipos de dados
	loopStatements() // Função que demonstra instruções de controle
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
}
