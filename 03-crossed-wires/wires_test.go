package main

import (
	"image"
	"testing"
)

type testcase struct {
	wire1  []movement
	wire2  []movement
	area   image.Rectangle
	start  image.Point
	result int
}

var cases = []testcase{
	{
		[]movement{
			{8, RIGHT},
			{5, UP},
			{5, LEFT},
			{3, DOWN},
		},
		[]movement{
			{7, UP},
			{6, RIGHT},
			{4, DOWN},
			{4, LEFT},
		},
		image.Rectangle{image.Point{0, 0}, image.Point{8, 7}},
		image.Pt(0, 0),
		6,
	},
	{
		[]movement{
			{75, RIGHT},
			{30, DOWN},
			{83, RIGHT},
			{83, UP},
			{12, LEFT},
			{49, DOWN},
			{71, RIGHT},
			{7, UP},
			{72, LEFT},
		},
		[]movement{
			{62, UP},
			{66, RIGHT},
			{55, UP},
			{34, RIGHT},
			{71, DOWN},
			{55, RIGHT},
			{58, DOWN},
			{83, RIGHT},
		},
		image.Rectangle{image.Point{0, 0}, image.Point{238, 147}},
		image.Pt(0, 30),
		159,
	},
	{
		[]movement{
			{98, RIGHT},
			{47, UP},
			{26, RIGHT},
			{63, DOWN},
			{33, RIGHT},
			{87, UP},
			{62, LEFT},
			{20, DOWN},
			{33, RIGHT},
			{53, UP},
			{51, RIGHT},
		},
		[]movement{
			{98, UP},
			{91, RIGHT},
			{20, DOWN},
			{16, RIGHT},
			{67, DOWN},
			{40, RIGHT},
			{7, UP},
			{15, RIGHT},
			{6, UP},
			{7, RIGHT},
		},
		image.Rectangle{image.Point{0, 0}, image.Point{179, 120}},
		image.Point{0, 16},
		135,
	},
}

func TestGetArea(t *testing.T) {
	for _, c := range cases {
		start, area := getArea(c.wire1, c.wire2)
		if area != c.area {
			t.Errorf("TestGetArea failed(area): got %v, expected %v", area, c.area)
		}
		if start != c.start {
			t.Errorf("TestGetArea failed(start): got %v, expected %v", start, c.start)
		}
	}

}

func TestSimulation(t *testing.T) {
	for _, c := range cases {
		result := runSimulation(c.area, c.start, c.wire1, c.wire2)
		if result != c.result {
			t.Errorf("TestSimulation failed: got %v, expected %v", result, c.result)
		}
	}
}
