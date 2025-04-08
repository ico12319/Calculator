package expression

import (
	"Calculator/expression/mocks"
	"testing"
)

func TestMultiplication_Eval(t *testing.T) {
	tests := []struct {
		testName string
		leftVal  float64
		rightVal float64
		want     float64
	}{
		{"12.0 * 3.0 = 36.0", 12.0, 3.0, 36.0},
		{"3.1 * 4.0 = 12.40", 3.1, 4.0, 12.40},
		{"5.5 * 5.5 = 30.25", 5.5, 5.5, 30.25},
		{"1.0 * 0.0 = 0.0", 1.0, 0.0, 0.0},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			mockExpressionLeft := new(mocks.Expression)
			mockExpressionRight := new(mocks.Expression)

			mockExpressionLeft.EXPECT().Eval().Return(test.leftVal).Once()
			mockExpressionRight.EXPECT().Eval().Return(test.rightVal).Once()

			multiplicationObject := NewMultiplication(mockExpressionLeft, mockExpressionRight)
			want := test.want
			got := multiplicationObject.Eval()
			assertInequality(t, got, want)

			mockExpressionLeft.AssertExpectations(t)
			mockExpressionRight.AssertExpectations(t)
		})
	}
}
