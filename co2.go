package main

import (
	"fmt"
)

var co2Tests = []struct {
	testName    string
	skip        bool
	co2Value    interface{}
	description string
}{
	// Проверки нижней границы
	{
		testName:    "Co2 -500",
		description: "Класс (-∞:-1] типичный представитель",
		co2Value:    -500,
	},
	{
		testName:    "Co2 -2",
		description: "Класс (-∞:-1] шаг в границу",
		co2Value:    -2,
	},
	{
		testName:    "Co2 -1",
		description: "Класс (-∞:-1] граничное значение и шаг за границу для класса [0:1000)",
		co2Value:    -1,
	},
	{
		testName:    "Co2 0",
		description: "Класс (-∞:-1] шаг за границу и граничное значение для класса [0:1000)",
		co2Value:    0,
	},
	{
		testName:    "Co2 1",
		description: "Шаг за границу для класса [0:1000)",
		co2Value:    1,
	},
	{
		testName:    "Co2 500",
		description: "Класс [0:1000) типичный представитель",
		co2Value:    500,
	},
	// Граница конвертации в тонны
	{
		testName:    "Co2 998",
		description: "Класс [0:1000) шаг в границу ",
		co2Value:    998,
	},
	{
		testName:    "Co2 999",
		description: "Класс [0:1000) граничное значение и шаг за границу для класса [1000:1000000)",
		co2Value:    999,
	},
	{
		testName:    "Co2 1000",
		description: "Класс [0:1000) шаг за границу и граничное значение для класса [1000:1000000)",
		co2Value:    1000,
	},
	{
		testName:    "Co2 1001",
		description: "Шаг за границу для класса [1000:1000000)",
		co2Value:    1001,
	},
	{
		testName:    "Co2 550000",
		description: "Класс [1000:1000000) типичный представитель",
		co2Value:    550000,
	},
	// Проверки конвертации в тысячи тонн
	{
		testName:    "Co2 999998",
		description: "Класс [1000:1000000) шаг в границу ",
		co2Value:    999998,
	},
	{
		testName:    "Co2 999999",
		description: "Класс [1000:1000000) граничное значение и шаг за границу для класса [1000000:1000000000)",
		co2Value:    999999,
	},
	{
		testName:    "Co2 1000000",
		description: "Класс [1000:1000000) шаг за границу и граничное значение для класса [1000000:1000000000)",
		co2Value:    1000000,
	},
	{
		testName:    "Co2 1000001",
		description: "Шаг за границу для класса [1000000:1000000000)",
		co2Value:    1000001,
	},
	{
		testName:    "Co2 550000000",
		description: "Класс [1000000:1000000000) типичный представитель",
		co2Value:    550000000,
	},
	// Проверки конвертации в миллионы тонн
	{
		testName:    "Co2 999999998",
		description: "Класс [1000000:1000000000) шаг в границу ",
		co2Value:    999999998,
	},
	{
		testName:    "Co2 999999999",
		description: "Класс [1000000:1000000000) граничное значение и шаг за границу для класса [1000000000:1000000000000)",
		co2Value:    999999999,
	},
	{
		testName:    "Co2 1000000000",
		description: "Класс [1000000:1000000000) шаг за границу и граничное значение для класса [1000000000:1000000000000)",
		co2Value:    1000000000,
	},
	{
		testName:    "Co2 1000000001",
		description: "Шаг за границу для класса [1000000000:1000000000000)",
		co2Value:    1000000001,
	},
	{
		testName:    "Co2 550000000000",
		description: "Класс [1000000000:1000000000000) типичный представитель",
		co2Value:    550000000000,
	},
	// Проверки конвертации в миллиарды тонн
	{
		testName:    "Co2 999999999998",
		description: "Класс [1000000000:1000000000000) шаг в границу ",
		co2Value:    999999999998,
	},
	{
		testName:    "Co2 999999999999",
		description: "Класс [1000000000:1000000000000) граничное значение и шаг за границу для класса [1000000000000:1000000000000000)",
		co2Value:    999999999999,
	},
	{
		testName:    "Co2 1000000000000",
		description: "Класс [1000000000:1000000000000) шаг за границу и граничное значение для класса [1000000000000:1000000000000000)",
		co2Value:    1000000000000,
	},
	{
		testName:    "Co2 1000000000001",
		description: "Шаг за границу для класса [1000000000000:1000000000000000)",
		co2Value:    1000000000001,
	},
	{
		testName:    "Co2 550000000000000",
		description: "Класс [1000000000000:1000000000000000)типичный представитель",
		co2Value:    550000000000000,
	},
	// Проверки конвертации в миллиарды тонн
	{
		testName:    "Co2 999999999999998",
		description: "Класс [1000000000000:1000000000000000) шаг в границу ",
		co2Value:    999999999999998,
	},
	{
		testName:    "Co2 999999999999999",
		description: "Класс [1000000000000:1000000000000000) граничное значение и шаг за границу для класса [1000000000000000:1000000000000000000)",
		co2Value:    999999999999999,
	},
	{
		testName:    "Co2 1000000000000000",
		description: "Класс [1000000000000:1000000000000000) шаг за границу и граничное значение для класса [1000000000000000:1000000000000000000)",
		co2Value:    1000000000000000,
	},
	{
		testName:    "Co2 1000000000000001",
		description: "Шаг за границу для класса [1000000000000000:1000000000000000000)",
		co2Value:    1000000000000001,
	},
	// Тесты на другие классы эквивалентности
	{
		testName:    "Co2 число 576 строкой",
		description: "Проверка на класс эквивалентности число строкой",
		co2Value:    "576",
	},
	{
		testName:    "Co2 дробное число 11.6",
		description: "Проверка на класс эквивалентности число строкой",
		co2Value:    11.6,
	},
	{
		testName:    "Co2 null вместо числа",
		description: "Проверка на класс эквивалентности null",
		co2Value:    nil,
	},
	{
		testName:    "Co2 @ вместо числа",
		description: "Проверка на класс эквивалентности спецсимвол",
		co2Value:    "@",
	},
	{
		testName:    "Co2 строка abc вместо числа",
		description: "Проверка на класс эквивалентности буквы",
		co2Value:    "abc",
	},
}

func RunCO2Tests(tcId int64) int64 {
	for _, td := range co2Tests {
		if td.skip {
			fmt.Printf("Skipped test [%s].\n", td.testName)
			tcId++
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
		fmt.Printf("Running test [%s].\n", td.testName)
		err := RunTest(req, fmt.Sprintf("output/%d.png", tcId))
		if err != nil {
			fmt.Printf("Error on making screenshot in test [%s]. Error: %s", td.testName, err.Error())
		}
		tcId++
	}
	return tcId
}
