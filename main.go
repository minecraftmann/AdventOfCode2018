package main

import (
	"fmt"
	"./solvers/q1"
)

func main() {
	p2 , err := q1.Part2()
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("Part 1 : ",p2)	
	}
	
}
