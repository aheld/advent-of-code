// way too much copy form https://github.com/nlowe/aoc2021/blob/master/challenge/day18/snailfish.go
package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

type snailfishNumber struct {
	parent *snailfishNumber

	left  *snailfishNumber
	right *snailfishNumber

	value int
}

func Parse(line string) (*snailfishNumber, string) {
	if line[0] == '[' {
		line = line[1:]
		result := &snailfishNumber{}
		result.left, line = Parse(line)
		result.left.parent = result

		if line[0] == ',' {
			line = line[1:]
		}

		result.right, line = Parse(line)
		result.right.parent = result

		return result, line[1:]
	}

	//Not nested

	idx := strings.IndexAny(line, "],")
	parts := strings.SplitN(line, string(line[idx]), 2)

	if line[idx] == ']' {
		line = "]" + parts[1]
	} else {
		line = parts[1]
	}
	//fmt.Println("I am here ", parts[0])
	value, _ := strconv.Atoi(parts[0])

	return &snailfishNumber{value: value}, line
}
func (s *snailfishNumber) Split() bool {

	target := findSplitTarget(s)
	if target == nil {
		return false
	}

	left := int(math.Floor(float64(target.value) / 2.0))
	right := int(math.Ceil(float64(target.value) / 2.0))

	target.value = 0
	target.left = &snailfishNumber{parent: target, value: left}
	target.right = &snailfishNumber{parent: target, value: right}

	return true
}

func findSplitTarget(s *snailfishNumber) *snailfishNumber {
	if s.left == nil && s.right == nil {
		if s.value >= 10 {
			return s
		}
		return nil
	}

	if target := findSplitTarget(s.left); target != nil {
		return target
	}

	return findSplitTarget(s.right)
}

func (s *snailfishNumber) Explode() bool {
	target := findExplodeTarget(0, s)
	//fmt.Println("Target =", target)
	if target == nil {
		return false
	}
	leftTarget := target
	for leftTarget.parent != nil {
		old := leftTarget
		leftTarget = leftTarget.parent
		if leftTarget.left != old {
			leftTarget = leftTarget.left
			break
		}
	}

	if leftTarget != s {
		for leftTarget.right != nil {
			leftTarget = leftTarget.right
		}
	}

	rightTarget := target
	for rightTarget.parent != nil {
		old := rightTarget
		rightTarget = rightTarget.parent

		if rightTarget.right != old {
			rightTarget = rightTarget.right
			break
		}
	}

	if rightTarget != s {
		for rightTarget.left != nil {
			rightTarget = rightTarget.left
		}
	}

	if leftTarget != nil {
		leftTarget.value += target.left.value
	}

	if rightTarget != nil {
		rightTarget.value += target.right.value
	}

	target.left = nil
	target.right = nil
	target.value = 0

	return true
}

func (s *snailfishNumber) add(other *snailfishNumber) *snailfishNumber {
	result := &snailfishNumber{
		left:  s,
		right: other,
	}

	s.parent = result
	other.parent = result

	result.reduce()
	return result
}

func (s *snailfishNumber) reduce() {
	for {
		if s.Explode() {
			continue
		}
		if s.Split() {
			continue
		}
		return
	}
}

func findExplodeTarget(n int, s *snailfishNumber) *snailfishNumber {
	if s.left == nil && s.right == nil {
		return nil
	}

	if n == 4 && s.left.left == nil && s.right.right == nil {
		return s
	}

	if target := findExplodeTarget(n+1, s.left); target != nil {
		return target
	}

	return findExplodeTarget(n+1, s.right)
}

func (s *snailfishNumber) String() string {
	if s == nil {
		return ""
	}

	if s.left == nil && s.right == nil {
		return strconv.Itoa(s.value)
	}

	if s.left == nil || s.right == nil {
		panic(fmt.Errorf("bad snailfish %p: left=%p, right=%p, v=%d", s, s.left, s.right, s.value))
	}

	line := strings.Builder{}
	line.WriteRune('[')
	line.WriteString(s.left.String())
	line.WriteRune(',')
	line.WriteString(s.right.String())
	line.WriteRune(']')

	return line.String()
}

func AddAll(input string) *snailfishNumber {
	lines := strings.Split(input, "\n")
	var s *snailfishNumber
	for _, line := range lines {
		parsed, left := Parse(line)
		if left != "" {
			panic(fmt.Errorf("failed to parse '%s': leftover: %s", line, left))
		}

		if s == nil {
			s = parsed
		} else {
			s = s.add(parsed)
		}
	}

	return s
}

func (s *snailfishNumber) isLiteral() bool {
	return s.left == nil && s.right == nil
}
func (s *snailfishNumber) magnitude() int {
	if s.isLiteral() {
		return s.value
	}
	return 3*s.left.magnitude() + 2*s.right.magnitude()
}

func part1(filename string) int {
	data := loadFile(filename)
	sn := AddAll(data)
	return sn.magnitude()
}

func part2(filename string) int {
	largest := 0
	data := strings.Split(loadFile(filename), "\n")
	for _, left := range data {
		for _, right := range data {
			if left == right {
				continue
			}
			l, _ := Parse(left)
			r, _ := Parse(right)
			mag := l.add(r).magnitude()
			if mag > largest {
				largest = mag
			}
			r, _ = Parse(left)
			l, _ = Parse(right)
			mag = l.add(r).magnitude()
			if mag > largest {
				largest = mag
			}
		}
	}
	return largest
}

func loadFile(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return string(data)
}

func main() {
	fmt.Println("Part1 ", part1("input.txt"))
	fmt.Println("Part2 ", part2("input.txt"))
}
