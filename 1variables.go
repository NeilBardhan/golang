package main

import "fmt"

var deckSize int //Variable can be declared outside func main
//b := 5 This is an error; this cannot be done outside func main, valid within func main

func main(){
  var x string = "This throws a warning" //because go can infer datatype from rhs, 'string' is redundant in var declaration
  a := "scoopity"
  fmt.Println(a)
  card := "zippity zoppity"
  //card = 7 This is an error; go is a statically typed language => var types cannot be changed after declaration
	fmt.Println(card)
	card2 := 36.57 //Initialization of 'card2' NOTE : ':=' is only used when a var is declared and initialized 
	card2 = 358.75 //Reassignment of 'card2 note: reassignment is '=' operator
	fmt.Println(card2)
	deckSize = 52 //Previously declared, now being initialized. NOTE: '=' is used here because var was declared previously. Now, only value is being assigned
	fmt.Println(deckSize)
 }
