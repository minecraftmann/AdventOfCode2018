package q2

import (
	"io/ioutil"
	"strings"
	"time"
	"fmt"
)

func Part1() (int, error) {
	inp, err := readInput("./input/q2/one.txt")


	t2 := time.Now()


	if err != nil{
		return 0, err
	}else{
		repeats2 := 0
		repeats3 := 0
		for a:=0; a<len(inp); a++{
			if(isAnyCharRepeatedNTimes(2,inp[a])){
				repeats2++
			}
			if(isAnyCharRepeatedNTimes(3,inp[a])){
				repeats3++
			}
		}
		checksum := repeats3 * repeats2
		elapsed2 := time.Since(t2);
		fmt.Println("Took ",elapsed2)
		return checksum, nil
	}
}

func readInput(filepath string) ([]string, error) {
	content, err := ioutil.ReadFile(filepath)
	if err != nil{
		return nil, nil
	}else{
		lines := strings.Split(string(content), "\n")
		for a:= 0; a < len(lines); a++ {
			if lines[a][len(lines[a])-1] == 13{
				lines[a] = lines[a][:len(lines[a])-1]
			}
		}
		return lines, nil
	}
}

func isAnyCharRepeatedNTimes(n int, inp string) (bool) {
	m := make(map[int]int)
	for a := 0; a < len(inp); a++ {
		ascii := int(inp[a])
		if val, ok := m[ascii]; ok {
			val = val
			m[ascii] = m[ascii] + 1
		}else{
			m[ascii] = 1
		} 
	}
	for a := range m {
		if m[a] == n {
			return true
		}
	}
	return false
}
