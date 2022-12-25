package day25

import "strings"

type Puzzle struct{}

func parseInput(raw string) []string {
	return strings.Split(strings.TrimSpace(raw), "\n")
}

func exp(n, exp int) int {
	res := 1
	for i := 0; i < exp; i++ {
		res *= n
	}
	return res
}

func snafuToDecimal(snafu string) int {
	number := 0
	for i := 0; i < len(snafu); i++ {
		var digit int
		switch string(snafu[len(snafu)-1-i]) {
		case "2":
			digit += 2
		case "1":
			digit += 1
		case "0":
			digit += 0
		case "-":
			digit -= 1
		case "=":
			digit -= 2
		}
		number += exp(5, i) * digit
	}
	return number
}

func decimalToSnafu(num int) string {
	solution := ""
	for num > 0 {
		switch num % 5 {
		case 4:
			solution = "-" + solution
			num += 2
		case 3:
			solution = "=" + solution
			num += 3
		case 2:
			solution = "2" + solution
		case 1:
			solution = "1" + solution
		case 0:
			solution = "0" + solution
		}
		num /= 5
	}

	return solution
}

func (*Puzzle) Part1(input string) string {
	snafuList := parseInput(input)

	sum := 0
	for _, snafu := range snafuList {
		x := snafuToDecimal(snafu)
		sum += x
	}

	sol := decimalToSnafu(sum)
	return sol
}

func (*Puzzle) Part2(input string) string {
	return "-"
}

func (*Puzzle) Notes() string {
	return "converter número base estranha"
}
