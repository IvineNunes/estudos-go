package tipos

import "fmt"

func Int() {

	//OBS: O compilador não compila numeros com o tamanho de bits diferentes: int8 + int64//

	// - int8: Inteiro de 8 bits (valores entre -128 a 127)
	// - int16: Inteiro de 16 bits (valores entre -32768 a 32767)
	// - int32: Inteiro de 32 bits (valores entre -2147483648 a 2147483647)
	// - int64 ou int: Inteiro de 64 bits (caralhada de numeros)

	var numero int = 22
	fmt.Println(numero)
}
