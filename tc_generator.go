package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const (
	testCaseFileName = "TESTCASES.md"
)

func GenerateTestCases() {
	example, _ := os.ReadFile("tc_example.txt")
	resultTC := "# TestCases для скриншотного тестирования\n\n"
	id := 1
	for _, td := range co2Tests {
		if td.skip {
			fmt.Printf("Skipped test [%s].\n", td.testName)
			continue
		}
		req := InitRequest{
			Result: Result{
				Blocks: Blocks{PersonalImpact: PersonalImpact{Data: Data{
					Co2:    td.co2Value,
					Water:  0,
					Energy: 0,
				}}},
				IsAuthorized: true,
			},
			Status: "ok",
		}
		body, _ := json.Marshal(req)
		resultTC += fmt.Sprintf(string(example), id, td.testName, td.description, string(body), "Co2", strings.ReplaceAll(td.testName, " ", ""))
		id++
	}
	for _, td := range waterTests {
		if td.skip {
			fmt.Printf("Skipped test [%s].\n", td.testName)
			continue
		}
		req := InitRequest{
			Result: Result{
				Blocks: Blocks{PersonalImpact: PersonalImpact{Data: Data{
					Water:  td.waterValue,
					Co2:    0,
					Energy: 0,
				}}},
				IsAuthorized: true,
			},
			Status: "ok",
		}
		body, _ := json.Marshal(req)
		resultTC += fmt.Sprintf(string(example), id, td.testName, td.description, string(body), "Water", strings.ReplaceAll(td.testName, " ", ""))
		id++
	}
	for _, td := range energyTests {
		if td.skip {
			fmt.Printf("Skipped test [%s].\n", td.testName)
			continue
		}
		req := InitRequest{
			Result: Result{
				Blocks: Blocks{PersonalImpact: PersonalImpact{Data: Data{
					Energy: td.energyValue,
					Water:  0,
					Co2:    0,
				}}},
				IsAuthorized: true,
			},
			Status: "ok",
		}
		body, _ := json.Marshal(req)
		resultTC += fmt.Sprintf(string(example), id, td.testName, td.description, string(body), "Energy", strings.ReplaceAll(td.testName, " ", ""))
		id++
	}

	os.WriteFile(testCaseFileName, []byte(resultTC), 0644)
}
