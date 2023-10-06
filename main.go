package main

import (
	"fmt"
	"sort"
)

func getFibNumber() {
	var n int
	fmt.Scan(&n)
	var s []int
	s = make([]int, 3, 6)
	s[0] = 0
	s[1] = 1
	s[2] = 1
	if (n <= 2) {
		fmt.Println(1)
        return
	}
	for i := 3; i <= n; i++ {
		s = append(s, (s[i-1] + s[i-2]) % 10)
	}
	fmt.Println(s[n])
}

func getFibNumber1() {
	var n, m int
	fmt.Scanln(&n,&m)

	var s []int
	s = make([]int, 3, 6)
	s[0] = 0
	s[1] = 1
	s[2] = 1

	if (n <= 2) {
		fmt.Println(1)
        return
	}
	
	for i := 3; i <= n; i++ {
		s = append(s, (s[i-1] + s[i-2]))
	}
	flInt := float64(s[n]) / float64(m)
	fmt.Println(s[n], flInt)
}

func getGCD() {
	var a, b int
    fmt.Scan(&a, &b)
    for {
        if b == 0 {
            fmt.Println(a)
            break
        }
        if b > a {
            a,b = b,a
        }
        a,b = b, a%b
    }
}

type Segment struct {
	left int
	right int
}


func main() {
	var numOfSegments, left, right, numOfPoints int
	s := make([]Segment, 0)
	points := make([]int, 0)
	fmt.Scan(&numOfSegments)

	for i := 0 ;; i++ {
		_, err := fmt.Scanln(&left, &right)
		if err != nil {
			break
		}
		s = append(s, Segment{left: left, right: right})
	}

	sort.SliceStable(s, func(i, j int) bool {
		return s[i].right < s[j].right
	})

	for i := 0 ; i < numOfSegments; i++ {
		if i < numOfSegments - 1 && s[i].right >= s[i + 1].left && (len(points) == 0 || (len(points) > 0 && s[i].left > points[len(points) - 1] )) {
			numOfPoints++
			points = append(points, s[i].right)
		} else if len(points) > 0 && s[i].left < points[len(points) - 1] {
      continue
		} else if len(points) > 0 && s[i].left > points[len(points) - 1] {
			numOfPoints++
			points = append(points, s[i].right)
		} else if i < numOfSegments - 1 && ((i > 0 && s[i].left > s[i - 1].right && s[i].right < s[i + 1].left) || (i == 0 && s[i].right < s[i + 1].left)) {
			numOfPoints++
			points = append(points, s[i].right)
		} 
	}
	fmt.Printf("%d\n", numOfPoints)
	for _, value := range points {
		fmt.Printf("%d ", value)
	}
}

