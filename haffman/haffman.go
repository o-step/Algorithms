package main

import (
	"fmt"
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

func main() {
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
