package expression

import "testing"

func TestNumberSingleton_Eval(t *testing.T) {
	tests := []struct {
		testName string
		value    float64
		want     float64
	}{
		{"13.0", 13.0, 13.0},
		{"255.67", 256.67, 256.67},
		{"1.0", 1.0, 1.0},
		{"22222222.2222", 22222222.2222, 22222222.2222},
		{"231241.666565", 231241.666565, 231241.666565},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			singleton := NewNumberSingleton(test.value)
			want := test.want
			got := singleton.Eval()

			assertInequality(t, got, want)
		})
	}
}
