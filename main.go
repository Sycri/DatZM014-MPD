package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Sycri/DatZM014-MPD/bruteforce_powerset"
	"github.com/Sycri/DatZM014-MPD/bruteforce_prevalid"
	"github.com/Sycri/DatZM014-MPD/models"
	"github.com/Sycri/DatZM014-MPD/simulated_annealing"
	"github.com/Sycri/DatZM014-MPD/utils"
)

func main() {
	problem, err := utils.GetObjectFromFile[models.Problem]("./testdata/03_input.json")
	if err != nil {
		panic(err)
	}

	jsonBytes, err := json.Marshal(problem)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Problem: %s\n", string(jsonBytes))

	solvers := map[models.Solver]string{
		&bruteforce_prevalid.Solver{}: "Bruteforce pre-validation",
		&bruteforce_powerset.Solver{}: "Bruteforce powerset",
		&simulated_annealing.Solver{}: "Simulated annealing",
	}

	for solver, name := range solvers {
		startTime := time.Now()
		solution := solver.Solve(problem)
		elapsedTime := time.Since(startTime)

		solution.Combination.FillNames(&problem.Basket.Products, &problem.Stores)

		jsonBytes, err = json.Marshal(solution)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s solver solution (%s): %s\n", name, elapsedTime, string(jsonBytes))
	}
}
