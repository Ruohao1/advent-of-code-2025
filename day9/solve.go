package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

type pointPair struct {
	area int
	a, b point
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func area(a, b point) int {
	return (1 + abs(a.x-b.x)) * (1 + abs(a.y-b.y))
}

func part1(lines []string) int {
	pointPairs := []pointPair{}
	for i, line := range lines {
		split := strings.Split(line, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		a := point{x, y}
		for j, line2 := range lines {
			if i < j {
				continue
			}

			split := strings.Split(line2, ",")
			x, _ := strconv.Atoi(split[0])
			y, _ := strconv.Atoi(split[1])
			b := point{x, y}
			pointPairs = append(pointPairs, pointPair{area(a, b), a, b})
		}
	}

	slices.SortFunc(pointPairs, func(a, b pointPair) int {
		if a.area > b.area {
			return -1
		} else if a.area == b.area {
			return 0
		} else {
			return 1
		}
	})

	// fmt.Println(pointPairs)
	return pointPairs[0].area
}

type edge struct {
	a, b point
}

// properCross returns true if the two rectilinear segments cross in the
// *interior* of both (not just touch at endpoints or overlap).
func properCross(e1, e2 edge) bool {
	// Orientation: horizontal if y constant
	h1 := (e1.a.y == e1.b.y)
	h2 := (e2.a.y == e2.b.y)

	// Only vertical vs horizontal can cross
	if h1 == h2 {
		return false
	}

	var h, v edge
	if h1 {
		h = e1
		v = e2
	} else {
		h = e2
		v = e1
	}

	Hy := h.a.y
	Hx1 := minInt(h.a.x, h.b.x)
	Hx2 := maxInt(h.a.x, h.b.x)

	Vx := v.a.x
	Vy1 := minInt(v.a.y, v.b.y)
	Vy2 := maxInt(v.a.y, v.b.y)

	// intersection point is (Vx, Hy)
	// strict inequalities to avoid counting endpoints
	if Vx > Hx1 && Vx < Hx2 && Hy > Vy1 && Hy < Vy2 {
		return true
	}
	return false
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func polygonEdges(poly []point) []edge {
	n := len(poly)
	res := make([]edge, 0, n)
	for i := 0; i < n; i++ {
		j := (i + 1) % n
		res = append(res, edge{poly[i], poly[j]})
	}
	return res
}

func rectangleEdges(rect []point) []edge {
	// rect has 4 points in order from your rectangle() function
	n := len(rect)
	res := make([]edge, 0, n)
	for i := 0; i < n; i++ {
		j := (i + 1) % n
		res = append(res, edge{rect[i], rect[j]})
	}
	return res
}

func rectangle(a, c point) []point {
	minX, maxX := a.x, c.x
	if minX > maxX {
		minX, maxX = maxX, minX
	}
	minY, maxY := a.y, c.y
	if minY > maxY {
		minY, maxY = maxY, minY
	}

	p1 := point{minX, minY}
	p2 := point{maxX, minY}
	p3 := point{maxX, maxY}
	p4 := point{minX, maxY}
	poly := []point{p1, p2, p3, p4}

	return poly
}

func pointOnSegment(px, py float64, a, b point) bool {
	ax, ay := float64(a.x), float64(a.y)
	bx, by := float64(b.x), float64(b.y)

	// collinearity via cross product == 0
	cross := (bx-ax)*(py-ay) - (by-ay)*(px-ax)
	const eps = 1e-9
	if math.Abs(cross) > eps {
		return false
	}

	// within bounding box
	if px < math.Min(ax, bx)-eps || px > math.Max(ax, bx)+eps {
		return false
	}
	if py < math.Min(ay, by)-eps || py > math.Max(ay, by)+eps {
		return false
	}

	return true
}

// pointInPolygon returns true if (px,py) is strictly inside poly.
// poly is a slice of vertices in order; last edge is (n-1)->0.
func pointInPolygon(px, py float64, poly []point) bool {
	n := len(poly)
	if n < 3 {
		return false
	}

	inside := false

	// standard pattern: j = previous vertex
	for i, j := 0, n-1; i < n; j, i = i, i+1 {
		xi := float64(poly[i].x)
		yi := float64(poly[i].y)
		xj := float64(poly[j].x)
		yj := float64(poly[j].y)

		// does edge (j->i) straddle the horizontal line y = py?
		intersectsY := ((yi > py) != (yj > py))
		if !intersectsY {
			continue
		}

		// x coordinate of intersection of edge with horizontal ray at y=py
		xIntersect := xj + (py-yj)*(xi-xj)/(yi-yj)

		// if intersection is to the right of the point, we flip inside/outside
		if xIntersect > px {
			inside = !inside
		}
	}

	return inside
}

func pointInOrOnPolygon(px, py float64, poly []point) bool {
	n := len(poly)
	if n < 3 {
		return false
	}

	// boundary check
	for i, j := 0, n-1; i < n; j, i = i, i+1 {
		if pointOnSegment(px, py, poly[j], poly[i]) {
			return true
		}
	}

	// strict inside check
	return pointInPolygon(px, py, poly)
}

func isInside(rect []point, poly []point) bool {
	// 1) All rectangle vertices must be inside or on boundary
	for _, p := range rect {
		if !pointInOrOnPolygon(float64(p.x), float64(p.y), poly) {
			return false
		}
	}

	// 2) No polygon edge may strictly cross any rectangle edge
	pEdges := polygonEdges(poly)
	rEdges := rectangleEdges(rect)

	for _, pe := range pEdges {
		for _, re := range rEdges {
			if properCross(pe, re) {
				return false
			}
		}
	}

	return true
}

func part2(lines []string) int {
	polygon := []point{}
	for _, line := range lines {
		split := strings.Split(line, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		polygon = append(polygon, point{x, y})
	}

	res := 0
	for i := range len(polygon) {
		for j := i + 1; j < len(polygon); j++ {
			rect := rectangle(polygon[i], polygon[j])
			// fmt.Println(rect, isInside(rect, polygon))
			if isInside(rect, polygon) {
				rectArea := area(polygon[i], polygon[j])
				if rectArea > res {
					res = area(polygon[i], polygon[j])
				}
			}
		}
	}

	return res
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run solve.go <inputName> <part>")
		return
	}
	inputName := os.Args[1]
	part, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Error parsing part:", err)
		return
	}

	inputFile, err := os.OpenFile(inputName, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	lines := []string{}
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	result := 0

	switch part {
	case 1:
		result = part1(lines)
	case 2:
		result = part2(lines)
	default:
		fmt.Println("Unknown part:", part)
		return
	}

	fmt.Printf("Part %d: %d\n", part, result)
}
