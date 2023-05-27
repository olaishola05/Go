package main

import "fmt"

type Coordinate struct {
	x, y int
}

type Rectangle struct {
	a Coordinate //top left
	b Coordinate //bottom right
}

func width(rect Rectangle) int {
	return (rect.b.x - rect.a.x)
}

func length(rect Rectangle) int {
	return (rect.a.y - rect.b.x)
}

func areaOfRectangle(rect Rectangle) int {
	return length(rect) * width(rect)
}

func perimeterOfRectangle(rect Rectangle) int {
	return (width(rect) * 2) + (length(rect) * 2)
}

func printInfo(rect Rectangle) {
	fmt.Println("Area of a Rectangle:", areaOfRectangle(rect))
	fmt.Println("perimeter of a rectangle:", perimeterOfRectangle(rect))
}

func main() {
	rectangle := Rectangle{a: Coordinate{0, 30}, b: Coordinate{15, 0}}

	printInfo(rectangle)

	rectangle.a.y *= 2
	rectangle.b.x *= 2

	printInfo(rectangle)
}
