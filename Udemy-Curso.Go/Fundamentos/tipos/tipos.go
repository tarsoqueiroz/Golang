package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// números primos
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(3200))

	// sem sinal (só positivos)... uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	// com sinal... int8 int16 int32 int64
	i8, i16, i32, i64 := math.MaxInt8, math.MaxInt16, math.MaxInt32, math.MaxInt64
	fmt.Println("O valor máximo do int8  é", i8)
	fmt.Println("O tipo         do int8  é", reflect.TypeOf(i8))
	fmt.Println("O valor máximo do int16 é", i16)
	fmt.Println("O tipo         do int16 é", reflect.TypeOf(i16))
	fmt.Println("O valor máximo do int32 é", i32)
	fmt.Println("O tipo         do int32 é", reflect.TypeOf(i32))
	fmt.Println("O valor máximo do int64 é", i64)
	fmt.Println("O tipo         do int64 é", reflect.TypeOf(i64))

	var i2 rune = 'a' // representa um mapeamento da tabela Unicode (int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	// números reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99))

	// boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string
	s1 := "Olá meu nome é Zô"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string s1 é", len(s1))

	// string com múltiplas linhas
	s2 := `Olá
	meu
	nome
	é
	Zô`
	fmt.Println("O tamanho da string s2 é", len(s2))

	// char???
	char := 'a'
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
	fmt.Println(char)

}
