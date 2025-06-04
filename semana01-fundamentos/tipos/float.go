package tipos

import "fmt"

func Float() {

	//OBS: O compilador não compila numeros com o tamanho de bits diferentes: int8 + int64//

	var floatNumber float32 = 1.1
	var float float64 = 3.5

	fmt.Printf("floatNumber - float32: %f", floatNumber)
	fmt.Printf("float - float64: %f", float)
}
