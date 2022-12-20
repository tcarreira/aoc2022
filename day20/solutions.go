package day20

import (
	"fmt"
	"strings"
)

const decryptionKey = 811589153

type Puzzle struct{}

type DLLNode struct {
	value int
	next  *DLLNode
	prev  *DLLNode
}

func parseInput(raw string) []*DLLNode {
	numberList := []*DLLNode{}
	var prev *DLLNode = nil
	for _, line := range strings.Split(strings.TrimSpace(raw), "\n") {
		var number int
		fmt.Sscanf(line, "%d", &number)
		node := DLLNode{
			value: number,
			prev:  prev,
		}
		numberList = append(numberList, &node)
		prev = &node
	}

	// fill all *next
	for i, node := range numberList {
		node.next = numberList[(i+1)%len(numberList)]
	}

	// fill first *prev
	numberList[0].prev = numberList[len(numberList)-1]

	return numberList
}

func (n *DLLNode) shiftForward(stopIn int) {
	if stopIn == 0 {
		return
	}

	prev := n.prev
	curr := n
	next := n.next

	curr.next = next.next
	curr.prev = next
	prev.next = next
	next.prev = prev
	next.next.prev = curr
	next.next = curr

	n.shiftForward(stopIn - 1)
}

func (n *DLLNode) shiftBackward(stopIn int) {
	if stopIn == 0 {
		return
	}

	prev := n.prev
	curr := n
	next := n.next

	curr.next = prev
	curr.prev = prev.prev
	prev.next = next
	next.prev = prev
	prev.prev.next = curr
	prev.prev = curr

	n.shiftBackward(stopIn - 1)
}

func (n *DLLNode) mix() {
	value := n.value
	if value >= 0 {
		n.shiftForward(value)
	} else {
		n.shiftBackward(-value)
	}
}

func (n *DLLNode) mixPart2(length int) {
	value := n.value % (length - 1) // when I move, I'm not part of the list
	if value >= 0 {
		n.shiftForward(value)
	} else {
		n.shiftBackward(-value)
	}
}

func printList(node *DLLNode) {
	initial := node.value

	fmt.Print(node.value, " ")
	node = node.next
	for node.value != initial {
		fmt.Print(node.value, " ")
		node = node.next
	}
	fmt.Println()
}

func getSolution(zeroNode *DLLNode, size int) int {
	idx := []int{1000 % size, 2000 % size, 3000 % size}
	solution := 0
	curr := zeroNode
	for i := 0; i < size; i++ {
		if i == idx[0] || i == idx[1] || i == idx[2] {
			solution += curr.value
		}
		curr = curr.next
	}
	return solution
}

func (*Puzzle) Part1(input string) string {
	numberList := parseInput(input)
	var zeroNode *DLLNode
	for _, node := range numberList {
		if node.value == 0 {
			zeroNode = node
		}
	}

	// fmt.Println(nodesMap)
	for _, node := range numberList {
		// fmt.Printf("%p: %+v\n", node, *node)
		// printList(nodesMap[0])
		node.mix()
	}

	solution := getSolution(zeroNode, len(numberList))
	return fmt.Sprint(solution)
}

func (*Puzzle) Part2(input string) string {
	numberList := parseInput(input)

	var zeroNode *DLLNode
	for _, node := range numberList {
		if node.value == 0 {
			zeroNode = node
		}
		node.value *= decryptionKey
	}

	for i := 0; i < 10; i++ {
		for _, node := range numberList {
			node.mixPart2(len(numberList))
		}
	}

	solution := getSolution(zeroNode, len(numberList))
	return fmt.Sprint(solution)
}

func (*Puzzle) Notes() string {
	return ""
}
