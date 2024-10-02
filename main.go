package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonData := `{
	"Products": [
		{
			"ID": 1,
			"Name": "Product 1"
		},
		{
			"ID": 2,
			"Name": "Product 2"
		},
		{
			"ID": 3,
			"Name": "Product 3"
		}
	],
	"Stores": [
		{
			"ID": 1,
			"Name": "Store 1",
			"DayOfferings": [
				{
					"Day": 1,
					"Products": [
						{
							"ID": 1,
							"Price": 10
						},
						{
							"ID": 2,
							"Price": 20
						}
					]
				},
				{
					"Day": 2,
					"Products": [
						{
							"ID": 3,
							"Price": 30
						}
					]
				},
				{
					"Day": 3,
					"Products": [
						{
							"ID": 1,
							"Price": 5
						}
					]
				}
			]
		}
	],
	"Basket": {
		"Products": [
			{
				"ID": 1,
				"Quantity": 2
			},
			{
				"ID": 3,
				"Quantity": 1
			}
		],
		"MaxDays": 3
	}
}`

	problem := Problem{}
	if err := json.Unmarshal([]byte(jsonData), &problem); err != nil {
		panic(err)
	}

	fmt.Printf("Problem: %+v\n", problem)

	bruteSolution := problem.BruteSolve()

	jsonBytes, err := json.Marshal(bruteSolution)
	if err != nil {
		panic(err)
	}

	jsonData = string(jsonBytes)

	fmt.Printf("Bruteforce solution: %s\n", jsonData)
}
