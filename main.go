package main

import (
	"encoding/json"
	"fmt"

	"github.com/Sycri/DatZM014-MPD/bf_post_validate"
	"github.com/Sycri/DatZM014-MPD/bf_pre_validate"
	"github.com/Sycri/DatZM014-MPD/models"
	"github.com/Sycri/DatZM014-MPD/utils"
)

func main() {
	problem, err := utils.GetProblemFromFile("./input.json")
	if err != nil {
		panic(err)
	}

	jsonBytes, err := json.Marshal(problem)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Problem: %s\n", string(jsonBytes))

	solvers := map[models.Solver]string{
		&bf_pre_validate.Solver{}:  "Bruteforce pre-validate",
		&bf_post_validate.Solver{}: "Bruteforce post-validate",
	}

	for solver, name := range solvers {
		solution := solver.Solve(problem)

		jsonBytes, err = json.Marshal(solution)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s solver solution: %s\n", name, string(jsonBytes))
	}
}
