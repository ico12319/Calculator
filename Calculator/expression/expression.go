package expression

//go:generate mockery --dir=. --name=Expression --output=./mocks --filename=mock_expression.go --with-expecter=true
type Expression interface {
	Eval() float64
}
