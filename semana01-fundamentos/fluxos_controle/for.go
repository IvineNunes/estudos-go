package fluxos

import "fmt"

func For() {
	// for inicialização; expressão; fim de iteração

	// sum := 0

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// 	sum += i
	// }
	// fmt.Println(sum)

	// for sum < 20 {
	// 	fmt.Println(sum)
	// 	sum += 2
	// }

	nums := []int{1, 2, 3, 4, 5, 6}

	for j := 0; j < len(nums); j++ {
		fmt.Println(nums[j])
	}
}
