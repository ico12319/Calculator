package expression

import (
	"Calculator/expression/mocks"
	//"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddition_Eval(t *testing.T) {
	tests := []struct {
		testName string
		leftVal  float64
		rightVal float64
		want     float64
	}{
		{"5 + 12 = 17", 5.0, 12.0, 17.0},
		{"1 + 3 = 4", 1.0, 3.0, 4.0},
		{"17.54 + 3.45 = 20.99", 17.54, 3.45, 20.99},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			mockExpressionLeft := mocks.NewExpression(t)
			mockExpressionRight := mocks.NewExpression(t)

			mockExpressionLeft.On("Eval").Return(test.leftVal).Once()
			mockExpressionRight.On("Eval").Return(test.rightVal).Once()
			additionObj := NewAddition(mockExpressionLeft, mockExpressionRight)

			wanted := test.want
			got := additionObj.Eval()
			if wanted != got {
				t.Errorf("got %.2f expected %.2f", got, wanted)
			}

			mockExpressionLeft.AssertExpectations(t)
			mockExpressionRight.AssertExpectations(t)
		})
	}
}
