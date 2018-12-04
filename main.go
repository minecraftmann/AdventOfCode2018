package main

import (
	"fmt"
	"./solvers/q1"
	"./solvers/q2"
	"time"
)

func main() {

	// Day 1
	t1 := time.Now()
	res1, err1 := q1.Part2()
	if err1 != nil {
		fmt.Println(err1)
	}else{
		fmt.Println(res1)	
	}
	elapsed1 := time.Since(t1);
	fmt.Println("Took ",elapsed1)

	// Day 2
	res2, err2 := q2.Part1()
	if err2 != nil {
		fmt.Println(err2)
	}else{
		fmt.Println(res2)	
	}

}

