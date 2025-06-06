package compostos

import "fmt"

func Maps() {

	var pessoas = map[string]int{}

	pessoas["Ivine"] = 22
	pessoas["Debora"] = 21

	if idade, ok := pessoas["Ivine"]; ok {
		fmt.Println("Pessoa existe no map", idade, ok)
	} else {
		fmt.Println("Pessoas não existe no map")
	}

	fmt.Println(pessoas)
	delete(pessoas, "Debora")
	fmt.Println(pessoas["Ivine"])
	fmt.Println(pessoas)

}
