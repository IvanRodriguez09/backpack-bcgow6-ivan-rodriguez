package ordering

import "fmt"

func ordering(numbers ...int) {
	fmt.Println(numbers)
	for i := range numbers {
		for j := range numbers {
			if i > j {
				temp := numbers[i]
				numbers[i] = numbers[j]
				numbers[j] = temp
			}
		}
	}
	fmt.Println(numbers)
}
