package main

import (
	"fmt"
)

var energyTests = []struct {
	testName    string
	skip        bool
	energyValue interface{}
	description string
}{
	// Проверки нижней границы
	{
		testName:    "Energy -500",
		description: "Класс (-∞:-1] типичный представитель",
		energyValue: -500,
	},
	{
		testName:    "Energy -2",
		description: "Класс (-∞:-1] шаг в границу",
		energyValue: -2,
	},
	{
		testName:    "Energy -1",
		description: "Класс (-∞:-1] граничное значение и шаг за границу для класса [0:1000)",
		energyValue: -1,
	},
	{
		testName:    "Energy 0",
		description: "Класс (-∞:-1] шаг за границу и граничное значение для класса [0:1000)",
		energyValue: 0,
	},
	{
		testName:    "Energy 1",
		description: "Шаг за границу для класса [0:1000)",
		energyValue: 1,
	},
	{
		testName:    "Energy 500",
		description: "Класс [0:1000) типичный представитель",
		energyValue: 500,
	},
	// Граница конвертации в тонны
	{
		testName:    "Energy 998",
		description: "Класс [0:1000) шаг в границу ",
		energyValue: 998,
	},
	{
		testName:    "Energy 999",
		description: "Класс [0:1000) граничное значение и шаг за границу для класса [1000:1000000)",
		energyValue: 999,
	},
	{
		testName:    "Energy 1000",
		description: "Класс [0:1000) шаг за границу и граничное значение для класса [1000:1000000)",
		energyValue: 1000,
	},
	{
		testName:    "Energy 1001",
		description: "Шаг за границу для класса [1000:1000000)",
		energyValue: 1001,
	},
	{
		testName:    "Energy 550000",
		description: "Класс [1000:1000000) типичный представитель",
		energyValue: 550000,
	},
	// Проверки конвертации в тысячи тонн
	{
		testName:    "Energy 999998",
		description: "Класс [1000:1000000) шаг в границу ",
		energyValue: 999998,
	},
	{
		testName:    "Energy 999999",
		description: "Класс [1000:1000000) граничное значение и шаг за границу для класса [1000000:1000000000)",
		energyValue: 999999,
	},
	{
		testName:    "Energy 1000000",
		description: "Класс [1000:1000000) шаг за границу и граничное значение для класса [1000000:1000000000)",
		energyValue: 1000000,
	},
	{
		testName:    "Energy 1000001",
		description: "Шаг за границу для класса [1000000:1000000000)",
		energyValue: 1000001,
	},
	{
		testName:    "Energy 550000000",
		description: "Класс [1000000:1000000000) типичный представитель",
		energyValue: 550000000,
	},
	// Проверки конвертации в миллионы тонн
	{
		testName:    "Energy 999999998",
		description: "Класс [1000000:1000000000) шаг в границу ",
		energyValue: 999999998,
	},
	{
		testName:    "Energy 999999999",
		description: "Класс [1000000:1000000000) граничное значение и шаг за границу для класса [1000000000:1000000000000)",
		energyValue: 999999999,
	},
	{
		testName:    "Energy 1000000000",
		description: "Класс [1000000:1000000000) шаг за границу и граничное значение для класса [1000000000:1000000000000)",
		energyValue: 1000000000,
	},
	{
		testName:    "Energy 1000000001",
		description: "Шаг за границу для класса [1000000000:1000000000000)",
		energyValue: 1000000001,
	},
	{
		testName:    "Energy 550000000000",
		description: "Класс [1000000000:1000000000000) типичный представитель",
		energyValue: 550000000000,
	},
	// Проверки конвертации в миллиарды тонн
	{
		testName:    "Energy 999999999998",
		description: "Класс [1000000000:1000000000000) шаг в границу ",
		energyValue: 999999999998,
	},
	{
		testName:    "Energy 999999999999",
		description: "Класс [1000000000:1000000000000) граничное значение и шаг за границу для класса [1000000000000:1000000000000000)",
		energyValue: 999999999999,
	},
	{
		testName:    "Energy 1000000000000",
		description: "Класс [1000000000:1000000000000) шаг за границу и граничное значение для класса [1000000000000:1000000000000000)",
		energyValue: 1000000000000,
	},
	{
		testName:    "Energy 1000000000001",
		description: "Шаг за границу для класса [1000000000000:1000000000000000)",
		energyValue: 1000000000001,
	},
	{
		testName:    "Energy 550000000000000",
		description: "Класс [1000000000000:1000000000000000)типичный представитель",
		energyValue: 550000000000000,
	},
	// Проверки конвертации в миллиарды тонн
	{
		testName:    "Energy 999999999999998",
		description: "Класс [1000000000000:1000000000000000) шаг в границу ",
		energyValue: 999999999999998,
	},
	{
		testName:    "Energy 999999999999999",
		description: "Класс [1000000000000:1000000000000000) граничное значение и шаг за границу для класса [1000000000000000:1000000000000000000)",
		energyValue: 999999999999999,
	},
	{
		testName:    "Energy 1000000000000000",
		description: "Класс [1000000000000:1000000000000000) шаг за границу и граничное значение для класса [1000000000000000:1000000000000000000)",
		energyValue: 1000000000000000,
	},
	{
		testName:    "Energy 1000000000000001",
		description: "Шаг за границу для класса [1000000000000000:1000000000000000000)",
		energyValue: 1000000000000001,
	},
	// Тесты на другие классы эквивалентности
	{
		testName:    "Energy число 576 строкой",
		description: "Проверка на класс эквивалентности число строкой",
		energyValue: "576",
	},
	{
		testName:    "Energy дробное число 11.6",
		description: "Проверка на класс эквивалентности число строкой",
		energyValue: 11.6,
	},
	{
		testName:    "Energy null вместо числа",
		description: "Проверка на класс эквивалентности null",
		energyValue: nil,
	},
	{
		testName:    "Energy @ вместо числа",
		description: "Проверка на класс эквивалентности спецсимвол",
		energyValue: "@",
	},
	{
		testName:    "Energy строка abc вместо числа",
		description: "Проверка на класс эквивалентности буквы",
		energyValue: "abc",
	},
}

func RunEnergyTests(tcId int64) int64 {

	for _, td := range energyTests {
		if td.skip {
			fmt.Printf("Skipped test [%s].\n", td.testName)
			tcId++
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
		fmt.Printf("Running test [%s].\n", td.testName)
		err := RunTest(req, fmt.Sprintf("output/%d.png", tcId))
		if err != nil {
			fmt.Printf("Error on making screenshot in test [%s]. Error: %s", td.testName, err.Error())
		}
		tcId++
	}
	return tcId
}
