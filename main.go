package main

import (
	"fmt"
	"math/rand"
)

const (
	bound = 10000
)

func generator(size, divisor int64) []int64 {
	rnd := rand.New(rand.NewSource(size + 42))
	prev := rnd.Int63n(size)
	// rands := make([]int64, size)
	rands := []int64{}
	for i := int64(0); len(rands) < int(size); i++ {
		if d := ((prev + 1) + rnd.Int63n(bound)); d&int64(1) == divisor {
			rands = append(rands, d)
			prev = d
		}
	}

	return rands
}

func main() {
	fmt.Println(generator(10, 0))
	fmt.Println(generator(10, 1))

}
