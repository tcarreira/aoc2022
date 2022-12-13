package day13

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Puzzle struct{}

type Pair [][]interface{}

func parsePackage(line string) []interface{} {
	pack := []interface{}{}
	err := json.Unmarshal([]byte(line), &pack)
	if err != nil {
		panic(fmt.Errorf("%q is not unmarshable", line))
	}
	return pack
}

func parseInput(raw string) []Pair {
	output := []Pair{}
	for _, pairStr := range strings.Split(strings.TrimSpace(raw), "\n\n") {
		lines := strings.Split(pairStr, "\n")
		pair := Pair{parsePackage(lines[0]), parsePackage(lines[1])}
		output = append(output, pair)
	}
	return output
}

func toInterfaceList(v interface{}) []interface{} {
	var result []interface{}
	if flist, ok := v.([]float64); ok {
		for _, f := range flist {
			result = append(result, f)
		}
	} else if xlist, ok := v.([]interface{}); ok {
		for _, x := range xlist {
			result = append(result, x)
		}
	}
	return result
}

func isCorrectOrder(e1, e2 interface{}) (correct bool, isDecided bool) {
	num1, isNum1 := e1.(float64)
	num2, isNum2 := e2.(float64)

	// rule 1
	if isNum1 && isNum2 {
		if num1 == num2 {
			return false, false
		}
		return num1 < num2, true
	}

	// rule 3
	if isNum1 && !isNum2 {
		return isCorrectOrder([]float64{num1}, e2)
	}
	if isNum2 && !isNum1 {
		return isCorrectOrder(e1, []float64{num2})
	}

	// rule 2 (both lists)
	list1 := toInterfaceList(e1)
	list2 := toInterfaceList(e2)

	for i := 0; i < len(list1) && i < len(list2); i++ {
		result, isDecided := isCorrectOrder(list1[i], list2[i])
		if isDecided {
			return result, true
		}
	}
	if len(list1) == len(list2) {
		return false, false
	}
	return len(list1) < len(list2), true
}

func (*Puzzle) Part1(input string) string {
	pairs := parseInput(input)

	sum := 0
	for i, pair := range pairs {
		if isCorrect, _ := isCorrectOrder(pair[0], pair[1]); isCorrect {
			sum += (i + 1)
		}
	}

	return fmt.Sprint(sum)
}

func (*Puzzle) Part2(input string) string {
	return "-"
}

func (*Puzzle) Notes() string {
	return ""
}
