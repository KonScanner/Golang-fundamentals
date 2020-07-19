package main

import (
	"fmt"
	"math"
)

func main() {
	var f float64

	fmt.Println("Enter a float value : ")
	_, err := fmt.Scanf("%f", &f)

	if err != nil {
		fmt.Println(err) // Error handling
	}

	fmt.Printf("%.0f", math.Trunc(f)) // math package ensures that it does not get rounded up

}
