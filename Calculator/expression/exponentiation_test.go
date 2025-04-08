package expression

import (
	"Calculator/expression/mocks"
	"testing"
)

func TestExponentiation_Eval(t *testing.T) {
	tests := []struct {
		testName string
		leftVal  float64
		rightVal float64
		want     float64
	}{
		{"2.0 ^ 3.0 = 8.0", 2.0, 3.0, 8.0},
		{"2.0 ^ -2.0 = 0.25", 2.0, -2.0, 0.25},
		{"100.0 ^ 0.0 = 1.0", 100.0, 0.0, 1.0},
		{"2.5 ^ 2.5 = 9.82", 2.5, 2.5, 9.882117688026186},
		{"2.0 ^ 1.0 = 2.0", 2.0, 1.0, 2.0},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			mockExpressionLeft := new(mocks.Expression)
			mockExpressionRight := new(mocks.Expression)

			mockExpressionLeft.EXPECT().Eval().Return(test.leftVal).Once()
			mockExpressionRight.EXPECT().Eval().Return(test.rightVal).Once()

			expoObject := NewExponentiation(mockExpressionLeft, mockExpressionRight)

			want := test.want
			got := expoObject.Eval()

			assertInequality(t, got, want)

			mockExpressionLeft.AssertExpectations(t)
			mockExpressionRight.AssertExpectations(t)
		})
	}
}
