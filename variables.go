package main

/*  Notes:

Variable declaration
- var foo int
- var foo int = 42
- foo := 42 -> (Usually used)

- Cant redeclare variables, but can shadow them (from package level)
- All variables must be used (will get an error on compile time)

- Visibilty
	- lower case first letter for package scope (accessible in the file)
	- upper case first letter to export (accessible to other files)
	- no private scope

- Naming conventions
	- Pascal or camelCase
		- Capitalize accronoyms (HTTP, URL)
	- As short as reasonable
		- longer names for longer lives (descript or used on function name)

- Type conversions
	- destinationType(variable) -> convert as method
	- use strconv pacakge for strings
*/

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
