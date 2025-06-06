package compostos

import "fmt"

func Slices() {
	var gavetas []string

	// adicionando itens
	gavetas = append(gavetas, "pratos", "copos", "panos")

	// remoção de item
	gavetas = gavetas[:2]

	// printando todos e apenas o de indice 1
	fmt.Println(gavetas)
	fmt.Println(gavetas[1])

	// printando o tamanho (len) e recortando (x:x-1)
	fmt.Println(gavetas[1:3])
	fmt.Println(len(gavetas))

}
