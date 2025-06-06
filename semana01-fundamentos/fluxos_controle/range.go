package fluxos

import "fmt"

func Range() {
	// utilizado para acessar indices de tipos compostos, como slicess

	nums := []string{"Ivine", "Maria", "Oliveira"}

	for key, value := range nums {
		fmt.Println(key, value)
	}

	user := map[string]string{
		"Nome": "Ivine", "Idade": "30",
	}

	for key, value := range user {
		fmt.Println(key, value)
		fmt.Println(key)
	}
}
