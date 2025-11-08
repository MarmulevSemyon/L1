package main

import (
	"fmt"
)

// -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// -20:{-25.4, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20:{24.5}, 30:{32.5}.
func main() {

	seq := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	groups := make(map[int][]float64)
	for i := range seq {

		key := int(seq[i]) - int(seq[i])%10
		groups[key] = append(groups[key], seq[i])
	}

	fmt.Println(groups)
}
