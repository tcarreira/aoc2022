package day20

import (
	"fmt"
	"strings"
)

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

func (n *DLLNode) mix(nodesMap map[int]*DLLNode) {
	value := n.value
	if value >= 0 {
		n.shiftForward(value)
	} else {
		n.shiftBackward(-value)
	}
}

func buildNodesMap(nodes []*DLLNode) map[int]*DLLNode {
	myMap := map[int]*DLLNode{}
	for _, node := range nodes {
		myMap[node.value] = node
	}
	return myMap
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

func (*Puzzle) Part1(input string) string {
	numberList := parseInput(input)
	nodesMap := buildNodesMap(numberList)
	// fmt.Println(nodesMap)
	for _, node := range numberList {
		// fmt.Printf("%p: %+v\n", node, *node)
		// printList(nodesMap[0])
		node.mix(nodesMap)
	}

	idx := []int{1000 % len(numberList), 2000 % len(numberList), 3000 % len(numberList)}
	solution := 0
	curr := nodesMap[0]
	for i := 0; i < len(numberList); i++ {
		if i == idx[0] || i == idx[1] || i == idx[2] {
			solution += curr.value
		}
		curr = curr.next
	}
	return fmt.Sprint(solution)
}

func (*Puzzle) Part2(input string) string {
	return "-"
}

func (*Puzzle) Notes() string {
	return ""
}
