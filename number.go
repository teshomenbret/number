// Package provides  utlities to do 
// operation on number

package number

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer
	constraints.Float
}

// Add add two intiger and return the sum
//
// It has two parameter: a  and b,the numbers to be added,
// Let`s add this link to  test adding link [https://www.mathsisfun.com/numbers/addition.html]
func Add(a, b Number) Number{
	return a + b
}
