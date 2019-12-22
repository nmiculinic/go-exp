package main

import (
	"fmt"
	"github.com/nmiculinic/go-exp/advent_of_code/22/deck"
)

func main() {
	d := deck.New(10)
	fmt.Println(d.ToInts())
	d.DealWithIncrementN(3)
	fmt.Println(d.ToInts())
}
