package bf_pre_validate_test

import (
	"testing"

	"github.com/Sycri/DatZM014-MPD/bf_pre_validate"
	"github.com/Sycri/DatZM014-MPD/utils"
)

func TestSolve01(t *testing.T) {
	problem, err := utils.GetProblemFromFile("../test_data/01_input.json")
	if err != nil {
		t.Fatal(err)
	}

	solver := bf_pre_validate.Solver{}
	actual := solver.Solve(problem)

	expected, err := utils.GetSolutionFromFile("../test_data/01_output.json")
	if err != nil {
		t.Fatal(err)
	}

	utils.CompareSolve(t, expected, actual)
}
