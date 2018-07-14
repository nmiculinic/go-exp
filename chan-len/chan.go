package main

import (
	"fmt"
	"reflect"
)

func ilen(x interface{}) {
	fmt.Println(x, reflect.ValueOf(x).Len())
}

func main()  {
	c := make(chan int, 100)
	for i := 0; i < 10; i++{
		ilen(c)
		c <- i
	}
}
