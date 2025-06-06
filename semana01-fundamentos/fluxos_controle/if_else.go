package fluxos

import "fmt"

func IfElse() {

	var nota int

	fmt.Print("Digite sua nota:")
	fmt.Scanf("%d", &nota)

	if nota >= 90 {
		fmt.Println("Aprovado com distinção!")

	} else if nota >= 70 {
		fmt.Println("Aprovado")

	} else {
		fmt.Println("Reprovado")
	}

}
