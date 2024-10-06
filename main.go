package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Sycri/DatZM014-MPD/bruteforce"
	"github.com/Sycri/DatZM014-MPD/bruteforce_prevalid"
	"github.com/Sycri/DatZM014-MPD/models"
	"github.com/Sycri/DatZM014-MPD/simulated_annealing"
	"github.com/Sycri/DatZM014-MPD/utils"
)

func main() {
	problem, err := utils.GetProblemFromFile("./testdata/01_input.json")
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
		&bruteforce.Solver{}:          "Bruteforce no pre-validation",
		&simulated_annealing.Solver{}: "Simulated annealing",
	}

	for solver, name := range solvers {
		startTime := time.Now()
		solution := solver.Solve(problem)
		elapsedTime := time.Since(startTime)

		jsonBytes, err = json.Marshal(solution)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s solver solution (%s): %s\n", name, elapsedTime, string(jsonBytes))
	}
}
