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
	time   int
}

var cases = []testcase{
	{
		[]movement{
			{8, goRight},
			{5, goUp},
			{5, goLeft},
			{3, goDown},
		},
		[]movement{
			{7, goUp},
			{6, goRight},
			{4, goDown},
			{4, goLeft},
		},
		image.Rectangle{image.Point{0, 0}, image.Point{8, 7}},
		image.Pt(0, 0),
		6,
		30,
	},
	{
		[]movement{
			{75, goRight},
			{30, goDown},
			{83, goRight},
			{83, goUp},
			{12, goLeft},
			{49, goDown},
			{71, goRight},
			{7, goUp},
			{72, goLeft},
		},
		[]movement{
			{62, goUp},
			{66, goRight},
			{55, goUp},
			{34, goRight},
			{71, goDown},
			{55, goRight},
			{58, goDown},
			{83, goRight},
		},
		image.Rectangle{image.Point{0, 0}, image.Point{238, 147}},
		image.Pt(0, 30),
		159,
		610,
	},
	{
		[]movement{
			{98, goRight},
			{47, goUp},
			{26, goRight},
			{63, goDown},
			{33, goRight},
			{87, goUp},
			{62, goLeft},
			{20, goDown},
			{33, goRight},
			{53, goUp},
			{51, goRight},
		},
		[]movement{
			{98, goUp},
			{91, goRight},
			{20, goDown},
			{16, goRight},
			{67, goDown},
			{40, goRight},
			{7, goUp},
			{15, goRight},
			{6, goUp},
			{7, goRight},
		},
		image.Rectangle{image.Point{0, 0}, image.Point{179, 120}},
		image.Point{0, 16},
		135,
		410,
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
		res1, res2 := runSimulation(c.area, c.start, c.wire1, c.wire2)
		if res1 != c.result {
			t.Errorf("TestSimulation failed(1): got %v, expected %v", res1, c.result)
		}
		if res2 != c.time {
			t.Errorf("TestSimulation failed(2): got %v, expected %v", res2, c.time)
		}
	}
}

func TestParseMovements(t *testing.T) {
	cases := []struct {
		input    string
		expected []movement
	}{
		{
			"R8,U5,L5,D3",
			[]movement{
				{8, goRight},
				{5, goUp},
				{5, goLeft},
				{3, goDown},
			},
		},
		{
			"U7,R6,D4,L4",
			[]movement{
				{7, goUp},
				{6, goRight},
				{4, goDown},
				{4, goLeft},
			},
		},
	}

	for _, c := range cases {
		result := parseMovements(c.input)
		if len(result) != len(c.expected) {
			t.Errorf("TestParseMovements failed (len): got %v, expected %v", result, c.expected)
		}
		for i := range result {
			if result[i] != c.expected[i] {
				t.Errorf("TestParseMovements failed (eq): got %v, expected %v", result, c.expected)
			}
		}
	}

}
