package expression

import (
	"Calculator/expression/mocks"
	"testing"
)

func TestSubtraction_Eval(t *testing.T) {
	tests := []struct {
		testName string
		leftVal  float64
		rightVal float64
		want     float64
	}{
		{"4.5 - 3.5 = 1.0", 4.5, 3.5, 1.0},
		{"66.0 - 66.0 = 0.0", 66.0, 66.0, 0.0},
		{"43.5 - 23.5 = 20.0", 43.5, 23.5, 20.0},
		{"15.0 - 17.0 = -2.0", 15.0, 17.0, -2.0},
		{"100.0 - 500.0 = -400.0", 100.0, 500.0, -400.0},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			mockExpressionLeft := new(mocks.Expression)
			mockExpressionRight := new(mocks.Expression)

			mockExpressionLeft.EXPECT().Eval().Return(test.leftVal).Once()
			mockExpressionRight.EXPECT().Eval().Return(test.rightVal).Once()
			subtractionObject := NewSubtraction(mockExpressionLeft, mockExpressionRight)
			want := test.want
			got := subtractionObject.Eval()
			assertInequality(t, got, want)

			mockExpressionLeft.AssertExpectations(t)
			mockExpressionRight.AssertExpectations(t)
		})
	}
}
