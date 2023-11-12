package main

import (
	"fmt"
	"unicode/utf8"
	"sort"
)

type Item struct {
	symbol string
	num int
}

type MyNode struct {
	data Item
	leftChild *MyNode
	rightChild *MyNode
}

func mainК() {
	var s string
	fmt.Scan(&s)

	dictionary := make(map[string]int)
	dictionaryEncoded := make(map[string]string)
	order := new([]MyNode)
	*order = make([]MyNode, 0)
	
	for i := 0; i < len(s); i++ {
		_, ok := dictionary[string(s[i])]
		if ok {
			dictionary[string(s[i])] += 1
			dictionaryEncoded[string(s[i])] = ""
		} else {
			dictionary[string(s[i])] = 1
			dictionaryEncoded[string(s[i])] = ""
		}
	}

	for k, v := range dictionary {
		*order = append(*order, MyNode{data: Item{symbol: k, num: v}, leftChild: nil, rightChild: nil})
	}

	p := new(MyNode)

	for {
		if len(dictionary) == 1 {
			dictionaryEncoded[string(s[0])] = "0"
			break
		}
		if len(*order) < 2 {
			break
		}
		fll := extractMin(order)
		frl := extractMin(order)
		p = makeTree(&fll, &frl)
		insertNode(order, *p)
	}

	encodedStr := "";

	traverseRecursive(p, dictionaryEncoded)

	for i := 0; i < len(s); i++ {
		code := dictionaryEncoded[string(s[i])]
		encodedStr += code
	}

	// RESULT
	fmt.Println(len(dictionaryEncoded), len(encodedStr))
	for key, val := range dictionaryEncoded {
		fmt.Printf("%s: %s\n", key, val)
	}
	fmt.Println(encodedStr)
}
// abcabcaacdddddd

func extractMin(slice *[]MyNode) (MyNode) {
	minVal := (*slice)[0].data.num
	minIndex := 0
	for i:= 0; i < len(*slice); i++ {
		if (*slice)[i].data.num < minVal {
			minVal = (*slice)[i].data.num
			minIndex = i
		}
	}
	node := (*slice)[minIndex]
	*slice = append((*slice)[:minIndex], (*slice)[minIndex+1:]...)
	return node
}

func insertNode(slice *[]MyNode, node MyNode) {
	*slice = append(*slice, node)
}

func makeTree(LefttChild *MyNode, RightChild *MyNode) *MyNode {
	p := new(MyNode)
	*p = MyNode{data: Item{symbol: LefttChild.data.symbol + RightChild.data.symbol, num: LefttChild.data.num + RightChild.data.num}, leftChild: LefttChild, rightChild: RightChild}
	return p
}

func traverseRecursive(node *MyNode, dictionary map[string]string) {
  if (node != nil) {

		if node.rightChild != nil {
			for i:= 0; i < len(node.rightChild.data.symbol); i++ {
				dictionary[string(node.rightChild.data.symbol[i])] += "1"
			}
		}

		if node.leftChild != nil {
			for i:= 0; i < len(node.leftChild.data.symbol); i++ {
				dictionary[string(node.leftChild.data.symbol[i])] += "0"
			}
		}
		
		traverseRecursive(node.leftChild, dictionary);
		traverseRecursive(node.rightChild, dictionary);
	}
}

type Data struct {
	symbol string
	code string
}

type DecodeNode struct {
	data Data
	leftChild *DecodeNode
	rightChild *DecodeNode
}

func main() {
	var k, l int
	var key, val string
	var encdedStr string
	dictionary := make(map[string]string)
	codes := make([]string, 0)

	fmt.Scanf("%d %d\n", &k, &l)
	for i := 0; i < k; i++ {
		n, err := fmt.Scanf("%s %s", &key, &val)
		if n != 2 || err != nil {
			break
		}
		dictionary[string(key[0])] = val
		codes = append(codes, val)
	}
	fmt.Scanf("%s\n", &encdedStr)

	sort.SliceStable(codes, func(i, j int) bool {
		return len(codes[i]) < len(codes[j])
	})

	p := new(DecodeNode)
	p = &DecodeNode{data: Data{symbol: "", code: ""}}

	for k, v := range dictionary {
		makeTree2(k, v, p)
	}

	resStr := new(string)

	traverseRecursive2(p, encdedStr, resStr, p)

	fmt.Println(*resStr)
}

func makeTree2(key, val string, p *DecodeNode) {
		if len(val) == 0 {
			return
		}

		tempKey := ""
		if len(val) == 1 {
			tempKey = key
		}

		if string(val[0]) == "1" {
			if p.rightChild == nil {
				p.rightChild = &DecodeNode{data: Data{symbol: tempKey, code: p.data.code + "1"}}
			}
			_, i := utf8.DecodeRuneInString(val)
			makeTree2(key, string(val[i:]), p.rightChild)
		}

		if string(val[0]) == "0" {
			if p.leftChild == nil {
				p.leftChild = &DecodeNode{data: Data{symbol: tempKey, code: p.data.code + "0"}}
			}
			_, i := utf8.DecodeRuneInString(val)
			makeTree2(key, string(val[i:]), p.leftChild)
		}
}

func traverseRecursive2(node *DecodeNode, inputStr string, resStr *string, rootNode *DecodeNode) {
  if (node != nil && len(inputStr) != 0) {
		if node.data.symbol == "" {
			if len(inputStr) == 1 {
				inputStr += "c"
			}
			firstChar := string(inputStr[0])
			_, i := utf8.DecodeRuneInString(inputStr)
			newStr := string(inputStr[i:])
			if firstChar == "0" {
				traverseRecursive2(node.leftChild, newStr, resStr, rootNode)
			} else if string(inputStr[0]) == "1" {
				traverseRecursive2(node.rightChild, newStr, resStr, rootNode)
			}
		} else {
			*resStr += node.data.symbol
			traverseRecursive2(rootNode, inputStr, resStr, rootNode)
		}
	}
}

// 4 14
// a: 0

// 4 14
// a: 0
// b: 10
// c: 110
// d: 111

// for a
// 1 1
// a: 1
// 1

// for aa
// 1 2
// a: 0
// 00

// for ab
// 2 2
// a: 0
// b: 1
// 01

// for abc
// 3 5
// a: 0
// b: 1
// c: 110
// 01110
