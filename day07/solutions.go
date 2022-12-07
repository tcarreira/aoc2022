package day07

import (
	"fmt"
	"strconv"
	"strings"
)

type Puzzle struct{}

type FSNode struct {
	parent   *FSNode
	name     string
	size     int
	children map[string]*FSNode
}

func parseInput(raw string) FSNode {
	filesystem := FSNode{
		name:     "/",
		children: make(map[string]*FSNode),
	}
	fsContext := &filesystem

	lines := strings.Split(strings.TrimSpace(raw), "\n")
	for _, line := range lines {
		resultLine := strings.Split(line, " ")

		if resultLine[0] == "$" {
			// 0 = $    | $
			// 1 = cd   | ls
			// 2 = args | Panic
			if resultLine[1] == "ls" {
				continue
			}

			// / == root
			// .. == parent
			// abc == children
			arg := resultLine[2]

			switch arg {
			case "/":
				fsContext = &filesystem
			case "..":
				fsContext = fsContext.parent
			default:
				fsContext = fsContext.children[arg]
			}
			continue
		}

		newNode := fsContext.createNode(resultLine)
		fsContext.children[newNode.name] = &newNode

		// cd()
	}

	return filesystem
}

func (fsContext *FSNode) createNode(resultLine []string) FSNode {
	if resultLine[0] == "dir" {
		return FSNode{
			name:     resultLine[1],
			parent:   fsContext,
			children: make(map[string]*FSNode),
		}
	}

	size, err := strconv.Atoi(resultLine[0])
	if err != nil {
		panic(fmt.Errorf("not parsable: %w", err))
	}
	return FSNode{
		name:     resultLine[1],
		parent:   fsContext,
		size:     size,
		children: nil,
	}
}

func (node *FSNode) computeSize() int {
	for _, child := range node.children {
		if len(child.children) == 0 {
			node.size += child.size
		} else {
			node.size += child.computeSize()
		}
	}
	return node.size
}
func (node *FSNode) getSumLessThan100000() int {
	sum := 0
	for _, child := range node.children {
		if len(child.children) == 0 {
			continue
		} else {
			sum += child.getSumLessThan100000()
			if child.size <= 100000 {
				sum += child.size
			}
		}
	}
	return sum
}

func (*Puzzle) Part1(input string) string {
	// filesystem := parseInput(input) -> FSNode
	//     readCommand()
	//     executeCommand()
	//        fileNode := readFileNode(line) -> FSNode
	//        filesystem.append(fileNode)
	// dirSizes := computeDirSizes(filesystem)
	// somar dir com tamanho <= 1000000
	filesystem := parseInput(input)
	filesystem.computeSize()
	result := filesystem.getSumLessThan100000()

	return fmt.Sprint(result)
}

func (*Puzzle) Part2(input string) string {
	return "-"
}

func (*Puzzle) Notes() string {
	return ""
}
