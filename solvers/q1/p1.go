package q1

import (
	"io/ioutil"
	"strings"
	"strconv"
	"fmt"
)

func Part1() (int, error) {
	content, err := ioutil.ReadFile("./input/q1/one.txt")
	if err != nil {
		return 0, err
	}else{
		total := 0
		lines := strings.Split(string(content), "\n")
		totals1 := []int{}
		totals2 := []int{}
		for i := 0; i < len(lines) ; i++ {
			if lines[i][len(lines[i])-1] == 13{
				lines[i] = lines[i][:len(lines[i])-1]
			}
			intParsed, err := strconv.ParseInt(lines[i],10,64)
			if err != nil {
				return 0, err
			}else{
				total += int(intParsed)
				totals1 = append(totals1, total)
				totals2 = append(totals2, total+137*477)
			}
		}
		flag := true;
		for l := 0;flag && l < len(totals2); l++ {
			for i := 0;flag && i < len(totals1) && flag; i++ {
				if totals2[l] == totals1[i]{
					fmt.Println("Part 2: ",totals2[l])
					flag = false;
				}
			}
		}
		return total, nil
	}
}