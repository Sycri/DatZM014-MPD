package utils

import (
	"encoding/json"
	"os"
)

func GetObjectFromFile[T any](filePath string) (*T, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	object := new(T)
	if err := json.NewDecoder(file).Decode(object); err != nil {
		return nil, err
	}

	return object, nil
}

func PutObjectToFile(filePath string, object any) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(object); err != nil {
		return err
	}

	return nil
}
