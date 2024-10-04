package models

type Solver interface {
	Solve(*Problem) *Solution
}

type Solution struct {
	Combination      Combination
	UsedDayCount     int
	TotalProductCost int64
	TotalCost        int64
}
