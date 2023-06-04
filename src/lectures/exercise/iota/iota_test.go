package main

import "testing"

func TestCalculate(t *testing.T) {
	add := Operation(Add)
	sub := Operation(Subtract)
	mul := Operation(Multiply)
	div := Operation(Divide)

	testCases := []struct {
		name     string
		operands Operands
		op       Operation
		want     uint
	}{
		{
			name:     "Addition",
			operands: Operands{a: 2, b: 2},
			op:       add,
			want:     4,
		},
		{
			name:     "Subtraction",
			operands: Operands{a: 10, b: 3},
			op:       sub,
			want:     7,
		},
		{
			name:     "Multiplication",
			operands: Operands{a: 3, b: 3},
			op:       mul,
			want:     9,
		},
		{
			name:     "Division",
			operands: Operands{a: 150, b: 2},
			op:       div,
			want:     75,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.op.calculate(tc.operands)
			if got != tc.want {
				t.Errorf("Expected %d but got %d", tc.want, got)
			}
		})
	}
}
