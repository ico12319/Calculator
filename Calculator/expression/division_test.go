package expression

import (
	"Calculator/expression/mocks"
	"math"
	"testing"
)

func TestDivision_Eval(t *testing.T) {
	tests := []struct {
		testName        string
		leftVal         float64
		rightVal        float64
		want            float64
		expectedInfinty bool
	}{
		{"12.0 / 4.0 = 3.0", 12.0, 4.0, 3.0, false},
		{"6.0 / 3.0 = 2.0", 6.0, 3.0, 2.0, false},
		{"65.54 / 13.32 = 4.92", 65.0, 10.0, 6.5, false},
		{"1.0 / 0.0 = error", 1.0, 0.0, math.Inf(0), true},
		{"3.0 / 1.0 = 3.0", 3.0, 1.0, 3.0, false},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			mockExpressionLeft := new(mocks.Expression)
			mockExpressionRight := new(mocks.Expression)

			mockExpressionLeft.EXPECT().Eval().Return(test.leftVal).Once()
			mockExpressionRight.EXPECT().Eval().Return(test.rightVal).Once()

			divisionObj := NewDivision(mockExpressionLeft, mockExpressionRight)
			want := test.want
			got := divisionObj.Eval()

			assertInfinity(t, got, test.expectedInfinty)
			assertInequality(t, got, want)

			mockExpressionLeft.AssertExpectations(t)
			mockExpressionRight.AssertExpectations(t)
		})
	}
}

func assertInfinity(t testing.TB, got float64, expectedInfinity bool) {
	t.Helper()
	if math.IsInf(got, 0) != expectedInfinity {
		t.Errorf("expected inf but received none")
	}
}

func assertInequality(t testing.TB, got float64, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("expected %.2f got %.2f", want, got)
	}
}
