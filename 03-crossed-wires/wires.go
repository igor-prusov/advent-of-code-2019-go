package main

import (
	"fmt"
	"image"
	"os"
)

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func abs(x int) int {
	if x > 0 {
		return x
	} else {
		return -x
	}
}

type direction int

const (
	_ = iota
	UP
	DOWN
	LEFT
	RIGHT
)

type movement struct {
	length    int
	direction direction
}

type field struct {
	matrix [][]int
	start  image.Point
}

func makeField(area image.Rectangle, start image.Point) field {
	var f = field{}
	f.matrix = make([][]int, area.Dy()+1)
	for i := range f.matrix {
		f.matrix[i] = make([]int, area.Dx()+1)
	}
	f.start = start
	return f
}

func doMovement(mv movement, p image.Point) image.Point {
	switch mv.direction {
	case UP:
		return image.Point{p.X, p.Y + mv.length}
	case DOWN:
		return image.Point{p.X, p.Y - mv.length}
	case LEFT:
		return image.Point{p.X - mv.length, p.Y}
	case RIGHT:
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
func runWire(f field, a image.Point, b image.Point, wire_id int) int {
	minDistance := len(f.matrix) + len(f.matrix[0])
	if a.X == b.X {
		for i := 0; i < b.Y-a.Y; i++ {
			p := image.Point{a.X, a.Y + i}
			dist := abs(p.X-f.start.X) + abs(p.Y-f.start.Y)
			v := f.matrix[p.Y][p.X]
			if v != wire_id && v != 0 && dist < minDistance && !p.Eq(f.start) {
				fmt.Println("Intersection at ", image.Point{p.X, p.Y}, "dist = ", dist)
				minDistance = dist
			}
			f.matrix[p.Y][p.X] = wire_id
		}
	} else if a.Y == b.Y {
		for i := 0; i < b.X-a.X; i++ {
			p := image.Point{a.X + i, a.Y}
			dist := abs(p.X-f.start.X) + abs(p.Y-f.start.Y)
			v := f.matrix[p.Y][p.X]
			if v != wire_id && v != 0 && dist < minDistance && !p.Eq(f.start) {
				fmt.Println("Intersection at ", image.Point{p.X, p.Y}, "dist = ", dist)
				minDistance = dist
			}
			f.matrix[p.Y][p.X] = wire_id
		}
	} else {
		panic("REEEE")
	}
	return minDistance

}

func runSimulation(area image.Rectangle, start image.Point, wire1 []movement, wire2 []movement) int {
	f := makeField(area, start)
	p := f.start
	for _, op := range wire1 {
		new_p := doMovement(op, p)
		r := image.Rectangle{p, new_p}.Canon()
		runWire(f, r.Min, r.Max, 1)

		p = new_p
	}

	p = f.start
	minDistance := len(f.matrix) + len(f.matrix[0])
	for _, op := range wire2 {
		new_p := doMovement(op, p)

		r := image.Rectangle{p, new_p}.Canon()
		d := runWire(f, r.Min, r.Max, 2)
		minDistance = min(d, minDistance)

		p = new_p
	}
	//printMatrix(matrix)
	return minDistance
}

func printMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
}

func main() {
	wire1 := []movement{
		{8, RIGHT},
		{5, UP},
		{5, LEFT},
		{3, DOWN},
	}

	wire2 := []movement{
		{7, UP},
		{6, RIGHT},
		{4, DOWN},
		{4, LEFT},
	}

	start, area := getArea(wire1, wire2)
	fmt.Println(area)
	runSimulation(area, start, wire1, wire2)

}
