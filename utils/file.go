package utils

import (
	"encoding/json"
	"os"

	"github.com/Sycri/DatZM014-MPD/models"
)

func GetProblemFromFile(filePath string) (*models.Problem, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	problem := &models.Problem{}
	if err := json.NewDecoder(file).Decode(problem); err != nil {
		return nil, err
	}

	return problem, nil
}

func GetSolutionFromFile(filePath string) (*models.Solution, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	solution := &models.Solution{}
	if err := json.NewDecoder(file).Decode(solution); err != nil {
		return nil, err
	}

	return solution, nil
}
