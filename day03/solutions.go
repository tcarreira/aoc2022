package day03

import (
	"fmt"
	"strings"
)

type Puzzle struct{}

type Ruckpack struct {
	half1 string
	half2 string
}

func parseInput(raw string) []Ruckpack {
	ruckpacks := strings.Split(strings.TrimSpace(raw), "\n")

	result := []Ruckpack{}
	for _, ruckpack := range ruckpacks {
		ruckpack := Ruckpack{
			half1: ruckpack[:len(ruckpack)/2],
			half2: ruckpack[len(ruckpack)/2:],
		}
		result = append(result, ruckpack)
	}

	return result
}

func findRepeatedItem(s1, s2 string) string {
	for _, c1 := range s1 {
		for _, c2 := range s2 {
			if c1 == c2 {
				return string(c1)
			}
		}
	}
	panic("unreachable")
}

func convertItemToPriority(c string) int {
	byteStr := []byte(c)
	cAscii := int(byteStr[0])
	if cAscii >= 97 && cAscii <= 122 {
		cAscii -= 96
	} else if cAscii >= 65 && cAscii <= 90 {
		cAscii -= (65 - 27)
	} else {
		panic("unreachable (not a-z or A-Z)")
	}
	return cAscii
}

func (*Puzzle) Part1(input string) string {
	// lista mochiloes
	//   lista de items
	//     primeira metade | segunda metade
	//
	// identificar caracter repetido nas duas metades
	// a:z == 1:26
	// A:Z == 27:52

	// 1. tratar o input: parseInput
	//   1.2 separar lista em duas metades
	// 2. encontrar o caracter repetido
	// 3. converter o item em prioridade
	// 4. somar as prioridades
	ruckpacks := parseInput(input)

	sum := 0
	for _, ruckPack := range ruckpacks {
		repeatedItem := findRepeatedItem(ruckPack.half1, ruckPack.half2)
		priority := convertItemToPriority(repeatedItem)
		sum += priority
	}

	return fmt.Sprint(sum)
}

func (*Puzzle) Part2(input string) string {
	return "-"
}

func (*Puzzle) Notes() string {
	return ""
}
