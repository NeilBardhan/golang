package main

import (
	"fmt"
	"math"
)

func main() {
	card0 := "AS"
	fmt.Println(card0)
	card1 := newCard()
	fmt.Println(card1)
	newCard2()
	fmt.Println("wut!")
	card3 := newCard3() //go assigns integer returned by 'newCard3()' to this var and makes it an int
	fmt.Println(card3)
	a := 2
	b := 3
	c := newCard4(a, b) //call func with args
	fmt.Println(c)
}

//func newCard() {} implies we are not returning anything
func newCard() string { //implies we are returning a string
	return "5D"
}

func newCard2() { //doesn't return anything, only logs output to console
	fmt.Println("Psych!")
}

func newCard3() int { //returns an integer
	return 22
}

func newCard4(a int, b int) float64 { //2 integers as arguments, returns int
	return math.Sqrt(float64(a + b)) //a,b or any combination of n integers must be type cast to float 64 before returning
}

/*
func <funcName>(<args>) <return type>
func newCard() string takes 0 arguments returns a string
func newCard() takes 0 arguments returns nothing
func newCard(a int, b int) float64 takes 2 int arguments, returns a float64
*/
