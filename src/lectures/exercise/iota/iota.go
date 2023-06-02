package main

import "fmt"

const (
	Add = iota
	Subtract
	Multiply
	Divide
)

type Operands struct {
	a, b uint
}

type Operation int

func (op Operation) calculate(operation Operands) uint {
	switch op {
	case Add:
		return operation.a + operation.b
	case Subtract:
		return operation.a - operation.b
	case Multiply:
		return operation.a * operation.b
	case Divide:
		return operation.a / operation.b
	}
	panic("Unhandled error")
}

func main() {
	add := Operation(Add)
	fmt.Println(add.calculate(Operands{2, 2}))

	sub := Operation(Subtract)
	fmt.Println(sub.calculate(Operands{10, 3}))

	mul := Operation(Multiply)
	fmt.Println(mul.calculate(Operands{3, 3}))

	div := Operation(Divide)
	fmt.Println(div.calculate(Operands{150, 2}))
}
