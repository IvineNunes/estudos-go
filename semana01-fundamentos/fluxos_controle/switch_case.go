package fluxos

import (
	"fmt"
	"time"
)

func SwitchCase() {

	fmt.Println("Quando é segunda?")

	today := time.Now().Weekday()

	switch time.Monday {
	case today + 0:
		fmt.Println("É hoje")
	case today + 1:
		fmt.Println("É amanhã")
	case today + 2:
		fmt.Println("É em dois dias")
	default:
		fmt.Println("Está longe ainda.")
	}
}
