package models

type Solver interface {
	Solve(*Problem) *Solution
}

type Solution struct {
	Combination  Combination
	UsedDayCount int
	ProductCost  int64
	Cost         int64
}
