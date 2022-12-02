input.txt
main.go
main_test.go
test_input.txtpackage main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
)

type Point struct {
	x, y int
}

type Target struct {
	start Point
	end   Point
}

type Probe struct {
	pos      Point
	maxY     int
	inTarget bool
	velocity Point
}

func part1(input string) (int, int) {
	re := regexp.MustCompile(`-?\d+`)
	points := re.FindAllString(input, -1)
	// target := Target{}
	start := Point{}
	end := Point{}
	start.x, _ = strconv.Atoi(string(points[0]))
	end.x, _ = strconv.Atoi(points[1])
	start.y, _ = strconv.Atoi(points[2])
	end.y, _ = strconv.Atoi(points[3])
	target := Target{start, end}
	fmt.Printf("%v\n", target)
	maxY := 0
	probesInTarget := 0
	for x := 0; x < 1000; x++ {
		for y := -1000; y < 1000; y++ {
			p := fireProbe(Probe{Point{0, 0}, y, false, Point{x, y}}, target)
			if p.inTarget {
				probesInTarget++
				if p.maxY > maxY {
					maxY = p.maxY
				}
			}
		}
	}
	return maxY, probesInTarget
}

func fireProbe(probe Probe, target Target) Probe {
	for x := 0; x < 1000; x++ {

		probe.pos.x = probe.pos.x + probe.velocity.x
		probe.pos.y = probe.pos.y + probe.velocity.y
		if probe.maxY < probe.pos.y {
			probe.maxY = probe.pos.y
		}
		if probe.velocity.x > 0 {
			probe.velocity.x--
		}
		probe.velocity.y--
		if probe.pos.x >= target.start.x && probe.pos.x <= target.end.x &&
			probe.pos.y >= target.start.y && probe.pos.y <= target.end.y {
			probe.inTarget = true
			return probe
		}
	}
	return probe
}

func part2(filename string) int {
	return -1
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
	// fmt.Println("Part1 ", part1("input.txt"))
	// fmt.Println("Part2 ", part2("input.txt"))
}
