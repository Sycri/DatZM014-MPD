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
	filePathIterations := map[string]int{
		"./testdata/01_input.json": 100,
		"./testdata/02_input.json": 100,
		"./testdata/03_input.json": 100,
		"./testdata/04_input.json": 5,
	}

	for filePath, iterations := range filePathIterations {
		fmt.Printf("File: %s\n", filePath)

		problem, err := utils.GetObjectFromFile[models.Problem](filePath)
		if err != nil {
			panic(err)
		}

		jsonBytes, err := json.Marshal(problem)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Problem: %s\n", string(jsonBytes))

		solvers := map[models.Solver]string{
			&bruteforce_powerset.Solver{}: "Bruteforce powerset",
			&bruteforce_prevalid.Solver{}: "Bruteforce pre-validation",
			&simulated_annealing.Solver{}: "Simulated annealing",
		}

		for solver, name := range solvers {
			averageTime := time.Duration(0)

			for i := 0; i < iterations; i++ {
				startTime := time.Now()
				solution := solver.Solve(problem)
				elapsedTime := time.Since(startTime)
				averageTime += elapsedTime

				solution.Combination.FillNames(&problem.Basket.Products, &problem.Stores)

				jsonBytes, err = json.Marshal(solution)
				if err != nil {
					panic(err)
				}

				fmt.Printf("%s solver solution #%d (%s): %s\n", name, i+1, elapsedTime, string(jsonBytes))
			}

			fmt.Printf("%s solver average time: %s\n", name, (averageTime / time.Duration(iterations)))
		}
	}
}
