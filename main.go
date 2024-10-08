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
		"./testdata/01_input.json": 10000,
		"./testdata/02_input.json": 1000,
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
			iterationInterval := time.Duration(iterations)

			sumSolveTime := time.Duration(0)
			sumIterationTime := time.Duration(0)

			startLoopTime := time.Now()

			for i := 0; i < iterations; i++ {
				startTime := time.Now()

				solution := solver.Solve(problem)

				elapsedSolveTime := time.Since(startTime)
				sumSolveTime += elapsedSolveTime

				solution.Combination.FillNames(&problem.Basket.Products, &problem.Stores)

				jsonBytes, err = json.Marshal(solution)
				if err != nil {
					panic(err)
				}

				elapsedIterationTime := time.Since(startTime)
				sumIterationTime += elapsedIterationTime

				fmt.Printf(
					"%s solver solution #%d (solve time: %s, iteration time: %s): %s\n",
					name, i+1, elapsedSolveTime, elapsedIterationTime, string(jsonBytes),
				)
			}

			elapsedLoopTime := time.Since(startLoopTime)
			averageLoopTime := elapsedLoopTime / iterationInterval

			averageSolveTime := sumSolveTime / iterationInterval
			averageIterationTime := sumIterationTime / iterationInterval

			fmt.Printf(
				"%s solver average solve time - %s, average iteration time - %s, elapsed loop time - %s, average loop time - %s\n",
				name, averageSolveTime, averageIterationTime, elapsedLoopTime, averageLoopTime,
			)
		}
	}
}
