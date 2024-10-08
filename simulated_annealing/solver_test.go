package simulated_annealing_test

import (
	"testing"

	"github.com/Sycri/DatZM014-MPD/models"
	"github.com/Sycri/DatZM014-MPD/simulated_annealing"
	"github.com/Sycri/DatZM014-MPD/utils"
)

func TestSolve01(t *testing.T) {
	problem, err := utils.GetObjectFromFile[models.Problem]("../testdata/01_input.json")
	if err != nil {
		t.Fatal(err)
	}

	solver := simulated_annealing.Solver{}
	actual := solver.Solve(problem)
	actual.Combination.FillNames(&problem.Basket.Products, &problem.Stores)

	expected, err := utils.GetObjectFromFile[models.Solution]("../testdata/01_output.json")
	if err != nil {
		t.Fatal(err)
	}

	utils.CompareSolve(t, expected, actual)
}

func TestSolve02(t *testing.T) {
	problem, err := utils.GetObjectFromFile[models.Problem]("../testdata/02_input.json")
	if err != nil {
		t.Fatal(err)
	}

	solver := simulated_annealing.Solver{}
	actual := solver.Solve(problem)
	actual.Combination.FillNames(&problem.Basket.Products, &problem.Stores)

	expected, err := utils.GetObjectFromFile[models.Solution]("../testdata/02_output.json")
	if err != nil {
		t.Fatal(err)
	}

	utils.CompareSolve(t, expected, actual)
}
