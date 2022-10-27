package main

import (
	"fmt"
	"strconv"
)

// Package Level
var (
	actorName = "Romeo Enso"
	companion = "Joy Enso"
	age       = 28
)

func main() {
	// 1. Type on declaration
	// var i int -> initialize later
	// i =42 -> initialize
	// var j int = 27 -> declare and assign
	// i := 42 -> normal used declaration
	// i := 42.2
	// fmt.Println("This is first", i)
	fmt.Println(actorName)

	// 2.Converting int to string
	var i int = 42
	fmt.Printf("%v, %T\n", i, i)

	var j string
	j = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", j, j)
}
