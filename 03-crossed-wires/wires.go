package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
	"strconv"
	"strings"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

type direction int

const (
	_ = iota
	goUp
	goDown
	goLeft
	goRight
)

type movement struct {
	length    int
	direction direction
}

type field struct {
	matrix [][]int
	wire1  [][]int
	wire2  [][]int
	time1  int
	time2  int
	start  image.Point
}

func makeField(area image.Rectangle, start image.Point) field {
	var f = field{}
	f.matrix = make([][]int, area.Dy()+1)
	f.wire1 = make([][]int, area.Dy()+1)
	f.wire2 = make([][]int, area.Dy()+1)
	for i := range f.matrix {
		f.matrix[i] = make([]int, area.Dx()+1)
		f.wire1[i] = make([]int, area.Dx()+1)
		f.wire2[i] = make([]int, area.Dx()+1)
	}
	f.start = start
	f.time1 = 0
	f.time2 = 0
	return f
}

func doMovement(mv movement, p image.Point) image.Point {
	switch mv.direction {
	case goUp:
		return image.Point{p.X, p.Y + mv.length}
	case goDown:
		return image.Point{p.X, p.Y - mv.length}
	case goLeft:
		return image.Point{p.X - mv.length, p.Y}
	case goRight:
		return image.Point{p.X + mv.length, p.Y}

	}
	fmt.Fprintln(os.Stderr, "Can't move to unknown direction")
	return image.Point{}
}

func addPoint(area image.Rectangle, p image.Point) image.Rectangle {
	area.Min.X = min(area.Min.X, p.X)
	area.Min.Y = min(area.Min.Y, p.Y)

	area.Max.X = max(area.Max.X, p.X)
	area.Max.Y = max(area.Max.Y, p.Y)

	return area
}

func updateArea(area image.Rectangle, ops []movement) image.Rectangle {
	currentPosition := image.Point{0, 0}
	for _, op := range ops {
		currentPosition = doMovement(op, currentPosition)
		area = addPoint(area, currentPosition)

	}
	return area

}

func getArea(wire1 []movement, wire2 []movement) (image.Point, image.Rectangle) {
	area := image.Rectangle{}
	area = updateArea(area, wire1)
	area = updateArea(area, wire2)
	return image.Pt(0, 0).Sub(area.Min), area.Sub(area.Min)
}

// return minimal point of intersection
func runWire(f *field, a image.Point, b image.Point, wireID int) (int, int) {
	minDistance := len(f.matrix) + len(f.matrix[0])
	minTime := len(f.matrix) * len(f.matrix[0])

	var wire [][]int
	var time *int
	if wireID == 1 {
		wire = f.wire1
		time = &f.time1
	} else if wireID == 2 {
		wire = f.wire2
		time = &f.time2
	} else {
		panic("Incorrect wire")
	}

	for p := a; !p.Eq(b); {

		dist := abs(p.X-f.start.X) + abs(p.Y-f.start.Y)
		v := f.matrix[p.Y][p.X]
		if wire[p.Y][p.X] == 0 {
			wire[p.Y][p.X] = *time
		}
		f.matrix[p.Y][p.X] = wireID
		if v != wireID && v != 0 && !p.Eq(f.start) {
			minDistance = min(dist, minDistance)

			t := f.wire1[p.Y][p.X] + f.wire2[p.Y][p.X]
			minTime = min(t, minTime)
		}
		*time++
		if a.X == b.X && a.Y < b.Y {
			p.Y++
		} else if a.X == b.X && a.Y > b.Y {
			p.Y--
		} else if a.Y == b.Y && a.X < b.X {
			p.X++
		} else if a.Y == b.Y && a.X > b.X {
			p.X--
		} else {
			panic("Can't run wire of zero length")
		}
	}
	return minDistance, minTime

}

func runSimulation(area image.Rectangle, start image.Point, wire1 []movement, wire2 []movement) (int, int) {
	f := makeField(area, start)
	p := f.start
	for _, op := range wire1 {
		newP := doMovement(op, p)
		runWire(&f, p, newP, 1)

		p = newP
	}

	p = f.start
	minDistance := len(f.matrix) + len(f.matrix[0])
	minTime := len(f.matrix) * len(f.matrix[0])
	for _, op := range wire2 {
		newP := doMovement(op, p)

		d, t := runWire(&f, p, newP, 2)
		minDistance = min(d, minDistance)
		minTime = min(t, minTime)

		p = newP
	}
	//printMatrix(matrix)
	return minDistance, minTime
}

func printMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
}

func parseMovements(input string) []movement {
	var result []movement
	for _, s := range strings.Split(input, ",") {
		var m movement

		switch s[0] {
		case 'R':
			m.direction = goRight
		case 'L':
			m.direction = goLeft
		case 'U':
			m.direction = goUp
		case 'D':
			m.direction = goDown
		default:
			fmt.Fprintln(os.Stderr, "Can't parse direction:", s[0])
			panic(1)
		}

		v, err := strconv.Atoi(s[1:])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Can't parse length:", s[1:], err)
		}
		m.length = v

		result = append(result, m)

	}
	return result

}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cant' open input.txt")
	}

	scanner := bufio.NewScanner(file)
	var input []string

	for scanner.Scan() {
		input = append(input, scanner.Text())
		if len(input) > 2 {
			fmt.Fprintf(os.Stderr, "Too many lines")
			panic(1)
		}
	}
	if len(input) != 2 {
		fmt.Fprintf(os.Stderr, "Need two lines in input")
		panic(1)
	}

	file.Close()

	wire1 := parseMovements(input[0])
	wire2 := parseMovements(input[1])

	start, area := getArea(wire1, wire2)
	res1, res2 := runSimulation(area, start, wire1, wire2)
	fmt.Println("Result =", res1)
	fmt.Println("Result =", res2)
}
