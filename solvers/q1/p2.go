package q1

import (
	"io/ioutil"
	"strings"
	"strconv"
	"fmt"
)

func Part2() (int, error) {
	content, err := ioutil.ReadFile("./input/q1/one.txt")
	if err != nil {
		return 0, err
	}else{

		total , err:= Part1() //477
		if err != nil {
			return 0, err
		}else{
			current := 0
			counter := 0

			m := make(map[int][]int) // map of modulo relations

			lines := strings.Split(string(content), "\n")
			
			for i := 0; i < len(lines) ; i++ {
				if lines[i][len(lines[i])-1] == 13{
					lines[i] = lines[i][:len(lines[i])-1]
				}
				intParsed, err := strconv.ParseInt(lines[i],10,64)
				if err != nil {
					return 0, err
				}else{
					current += int(intParsed)
					mapIndex := current%total
					if mapIndex<0{
						mapIndex += total;
					}
					m[mapIndex] = append(append(m[mapIndex], counter),current)
					counter = counter + 1
				}
			}

			fmt.Println(m)
			fmt.Println(m[15])

			for a := 0; a < len(m); a++ {
				if len(m[a]) == 4{
					t := (m[a][1]-m[a][3])/total 
					m[a] = []int{t}
				}
				if len(m[a]) == 6{
					t1 := (m[a][1]-m[a][3])/total
					t2 := (m[a][3]-m[a][5])/total
					t3 := (m[a][1]-m[a][5])/total 
					m[a] = []int{t1,t2,t3}
				}
				if len(m[a]) == 8{
					t1 := (m[a][1]-m[a][3])/total
					t2 := (m[a][1]-m[a][5])/total
					t3 := (m[a][1]-m[a][7])/total
					t4 := (m[a][3]-m[a][5])/total
					t5 := (m[a][3]-m[a][7])/total
					t6 := (m[a][5]-m[a][7])/total 
					m[a] = []int{t1,t2,t3,t4,t5,t6}
				}
			}


			fmt.Println("  ")

			fmt.Println("  ")

			fmt.Println(m[15])



			return 1, nil
		}	
	}

}
