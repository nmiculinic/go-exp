package deferexp

import (
	"fmt"
	"math/rand"
	"testing"
)


func done() {
	if rand.Float32() < 1e-19 { // Practically never
		fmt.Print("lottery")
	}
}

func withoutDefer() {
	done()
}

func withDefer() {
	defer done()

}

func BenchmarkWithoutDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		withoutDefer()
	}
}

func BenchmarkWitDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		withDefer()
	}
}
