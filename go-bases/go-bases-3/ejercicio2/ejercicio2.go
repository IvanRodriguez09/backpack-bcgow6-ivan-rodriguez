package main

import (
	"fmt"
	"errors"
)

func average(notes ... float64) (float64, error){
	var averageRes float64
	var err error
	for _, note := range notes {
		if note < 0 {
			err = errors.New("the notes don't be negative")
		}else {
			averageRes += note
		}
	}
	return (averageRes / float64(len(notes))), err
}

func main()  {
	averageRes, err := average(3,5,4,3,5,4)
	
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(averageRes)
	}
	
}