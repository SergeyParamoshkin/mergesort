package main

import (
	"fmt"
	"math/rand"
)

const (
	bound = 10000
)

func generator(size, divisor int64) chan int64 {

	// TODO: migrate to crypto

	c := make(chan int64, size)

	go func() {
		defer close(c)

		rnd := rand.New(rand.NewSource(size + 42))
		prev := rnd.Int63n(size)

		for i := int64(0); len(c) < int(size); i++ {
			if d := prev + rnd.Int63n(bound); d&int64(1) == divisor {
				c <- d
				prev = d
			}
		}
	}()

	return c
}

func main() {
	g1 := generator(10, 0)
	g2 := generator(10, 1)

	v1, _ := <-g1
	v2, _ := <-g2
	ok := true
	for g1 != nil || g2 != nil {
		if g1 == nil {
			fmt.Println(v2)
			v2, ok = <-g2
			if !ok {
				g2 = nil
			}
		} else if g2 == nil {
			fmt.Println(v1)
			v1, ok = <-g1
			if !ok {
				g1 = nil
			}
		} else if v1 < v2 {
			fmt.Println(v1)
			v1, ok = <-g1
			if !ok {
				g1 = nil
			}
		} else {
			fmt.Println(v2)
			v2, ok = <-g2
			if !ok {
				g2 = nil
			}
		}
	}
}
