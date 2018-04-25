package main

import . "fmt"

type Cr struct {
	radius float64
}



func main(){
	var c1 Cr
	c1.radius = 5
	Println(c1.getA())
	Println(g)
}

var g = 60

func (c Cr) getA() float64 {
	return 3.14 * c.radius
}

