package main

import (
	"fmt"
	"sort"
	"strings"
	"strconv"
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
	fmt.Scanln(&n, &m)

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


func getPointsOnSegments() {
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

	var size int
	fmt.Scan(&size)
	ss := make([]int, size)
	f := make([]int, size)
	var result []string
	for i := range s{
		fmt.Scan(&s[i], &f[i])
	}
	for i:=1; i<size; i++{
		for j:=i; j!=0 && f[j]<f[j-1]; j-- {
			ss[j-1], ss[j] = ss[j], ss[j-1]
			f[j-1], f[j] = f[j], f[j-1]
		}
	}
	count, item := 1, f[0]
	result = append(result, strconv.Itoa(f[0]))
	for i:=0; i<size; i++{
		if item < ss[i] {
			result = append(result, strconv.Itoa(f[i]))
			item = f[i]
			count++
		}
	}
	fmt.Println(count)
	fmt.Println(strings.Join(result, " "))
}

type Item struct {
	price float64
	volume float64
}

func main() {
	var n int
	var price, volume, totalPrice, W float64
	fmt.Scan(&n, &W)
	items := make([]Item, 0)

	for i := 0; i < n; i++ {
		fmt.Scan(&price, &volume)
		items = append(items, Item{price: price, volume: volume})
	}
	
	sort.SliceStable(items, func(i, j int) bool {
		return items[i].price / items[i].volume > items[j].price / items[j].volume
	})

		for i := 0; i < n; i++ {
			if (W == 0.000) {
				break;
			}
			if (items[i].volume <= W) {
				W = W - items[i].volume
				totalPrice += items[i].price
			} else {
				pricePerVol := items[i].price / items[i].volume
				totalPrice += pricePerVol * W
				W = 0.000;
			}
		}

		fmt.Printf("%.3f", totalPrice)
}


func mainR() {
	var n int
	fmt.Scan(&n)
	items := make([]string, 0)

	if n == 1 || n == 2 {
		fmt.Printf("%d \n %d", 1, n)
		return
	}

	for i := 1; i <= n; i++ {
		if (n - i) > i {
			n -= i
			items = append(items, strconv.Itoa(i))
		} else {
			items = append(items, strconv.Itoa(n))
			break
		}
	}

	fmt.Printf("%d\n%s", len(items), strings.Join(items, " "))
}
