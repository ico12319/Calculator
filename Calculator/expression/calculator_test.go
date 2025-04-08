package expression

import (
	"math"
	"testing"
)

func TestCalculator_Evaluate(t *testing.T) {
	tests := []struct {
		testName         string
		expressionString string
		want             float64
		expectPanic      bool
	}{
		{"test with lots of operations", "(5 + 12 * (4 + (4 ^ (6 + 1))))", 196661.0, false},
		{"small test", "((1 + 14) * 2)", 30.0, false},
		{"test with whitespaces", "(3          + 5)", 8.0, false},
		{"test case with unary minus", "((-12 + 8) * (-2))", 8.0, false},
		{"test division by zero", "1/0", math.Inf(0), true},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			if test.expectPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("expected panic but recived none")
					}
				}()
			}
			calc := NewCalculator(test.expressionString)
			want := test.want
			got := calc.Evaluate()
			assertInequality(t, got, want)
		})
	}
}
