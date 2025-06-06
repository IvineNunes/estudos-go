package compostos

import "fmt"

func Array() {
	var gavetas [2]string

	gavetas[0] = "copos"
	gavetas[1] = "pratos"

	fmt.Println(gavetas)
	fmt.Println(gavetas[1])
}
