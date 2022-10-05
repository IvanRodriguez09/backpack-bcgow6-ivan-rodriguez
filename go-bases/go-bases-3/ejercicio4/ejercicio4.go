package main

import (
	"fmt"
	"errors"
)

const (
	Minimum = "minimum"
	Average = "average"
	Maximum = "maximum"
)

func minimumFunc(nums ... float64) (min float64){
	for i, num := range nums {
		if i == 0 {
			min = num
		} else if (min > num) {
			min = num
		}
	}
	return min
}

func averageFunc(nums ... float64) (av float64){
	for _, num := range nums {
		av += num
	}
	return av / float64(len(nums))
}

func maximumFunc(nums ... float64) (max float64) {
	for i, num := range nums {
		if i == 0 {
			max = num
		} else if (max < num) {
			max = num
		}
	}
	return max
}

func operationManager(operation string)  (func(nums ... float64) float64, error){
	switch operation {
	case Minimum:
		return minimumFunc, nil
	case Average:
		return averageFunc, nil
	case Maximum:
		return maximumFunc, nil
	}
	return nil, errors.New("No existe una operación con ese nombre")
}

func test(operationType string, nums ... float64)  {
	operation, err := operationManager(operationType)
	if (err != nil) {
		fmt.Println(err)
	} else {
		fmt.Println(operation(nums ...))
	}
}

func main()  {
	test(Maximum, 2, 3, 3, 4, 4, 2, 4, 5)
	test(Average, 2, 3, 3, 4, 4, 2, 4, 5)
	test(Minimum, 2, 3, 3, 4, 4, 2, 4, 5)
}