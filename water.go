package main

import (
	"fmt"
)

var waterTests = []struct {
	testName    string
	skip        bool
	waterValue  interface{}
	description string
}{
	// Проверки нижней границы
	{
		testName:    "Water -500",
		description: "Класс (-∞:-1] типичный представитель",
		waterValue:  -500,
	},
	{
		testName:    "Water -2",
		description: "Класс (-∞:-1] шаг в границу",
		waterValue:  -2,
	},
	{
		testName:    "Water -1",
		description: "Класс (-∞:-1] граничное значение и шаг за границу для класса [0:1000)",
		waterValue:  -1,
	},
	{
		testName:    "Water 0",
		description: "Класс (-∞:-1] шаг за границу и граничное значение для класса [0:1000)",
		waterValue:  0,
	},
	{
		testName:    "Water 1",
		description: "Шаг за границу для класса [0:1000)",
		waterValue:  1,
	},
	{
		testName:    "Water 500",
		description: "Класс [0:1000) типичный представитель",
		waterValue:  500,
	},
	// Граница конвертации в тонны
	{
		testName:    "Water 998",
		description: "Класс [0:1000) шаг в границу ",
		waterValue:  998,
	},
	{
		testName:    "Water 999",
		description: "Класс [0:1000) граничное значение и шаг за границу для класса [1000:1000000)",
		waterValue:  999,
	},
	{
		testName:    "Water 1000",
		description: "Класс [0:1000) шаг за границу и граничное значение для класса [1000:1000000)",
		waterValue:  1000,
	},
	{
		testName:    "Water 1001",
		description: "Шаг за границу для класса [1000:1000000)",
		waterValue:  1001,
	},
	{
		testName:    "Water 550000",
		description: "Класс [1000:1000000) типичный представитель",
		waterValue:  550000,
	},
	// Проверки конвертации в тысячи тонн
	{
		testName:    "Water 999998",
		description: "Класс [1000:1000000) шаг в границу ",
		waterValue:  999998,
	},
	{
		testName:    "Water 999999",
		description: "Класс [1000:1000000) граничное значение и шаг за границу для класса [1000000:1000000000)",
		waterValue:  999999,
	},
	{
		testName:    "Water 1000000",
		description: "Класс [1000:1000000) шаг за границу и граничное значение для класса [1000000:1000000000)",
		waterValue:  1000000,
	},
	{
		testName:    "Water 1000001",
		description: "Шаг за границу для класса [1000000:1000000000)",
		waterValue:  1000001,
	},
	{
		testName:    "Water 550000000",
		description: "Класс [1000000:1000000000) типичный представитель",
		waterValue:  550000000,
	},
	// Проверки конвертации в миллионы тонн
	{
		testName:    "Water 999999998",
		description: "Класс [1000000:1000000000) шаг в границу ",
		waterValue:  999999998,
	},
	{
		testName:    "Water 999999999",
		description: "Класс [1000000:1000000000) граничное значение и шаг за границу для класса [1000000000:1000000000000)",
		waterValue:  999999999,
	},
	{
		testName:    "Water 1000000000",
		description: "Класс [1000000:1000000000) шаг за границу и граничное значение для класса [1000000000:1000000000000)",
		waterValue:  1000000000,
	},
	{
		testName:    "Water 1000000001",
		description: "Шаг за границу для класса [1000000000:1000000000000)",
		waterValue:  1000000001,
	},
	{
		testName:    "Water 550000000000",
		description: "Класс [1000000000:1000000000000) типичный представитель",
		waterValue:  550000000000,
	},
	// Проверки конвертации в миллиарды тонн
	{
		testName:    "Water 999999999998",
		description: "Класс [1000000000:1000000000000) шаг в границу ",
		waterValue:  999999999998,
	},
	{
		testName:    "Water 999999999999",
		description: "Класс [1000000000:1000000000000) граничное значение и шаг за границу для класса [1000000000000:1000000000000000)",
		waterValue:  999999999999,
	},
	{
		testName:    "Water 1000000000000",
		description: "Класс [1000000000:1000000000000) шаг за границу и граничное значение для класса [1000000000000:1000000000000000)",
		waterValue:  1000000000000,
	},
	{
		testName:    "Water 1000000000001",
		description: "Шаг за границу для класса [1000000000000:1000000000000000)",
		waterValue:  1000000000001,
	},
	{
		testName:    "Water 550000000000000",
		description: "Класс [1000000000000:1000000000000000)типичный представитель",
		waterValue:  550000000000000,
	},
	// Проверки конвертации в миллиарды тонн
	{
		testName:    "Water 999999999999998",
		description: "Класс [1000000000000:1000000000000000) шаг в границу ",
		waterValue:  999999999999998,
	},
	{
		testName:    "Water 999999999999999",
		description: "Класс [1000000000000:1000000000000000) граничное значение и шаг за границу для класса [1000000000000000:1000000000000000000)",
		waterValue:  999999999999999,
	},
	{
		testName:    "Water 1000000000000000",
		description: "Класс [1000000000000:1000000000000000) шаг за границу и граничное значение для класса [1000000000000000:1000000000000000000)",
		waterValue:  1000000000000000,
	},
	{
		testName:    "Water 1000000000000001",
		description: "Шаг за границу для класса [1000000000000000:1000000000000000000)",
		waterValue:  1000000000000001,
	},
	// Тесты на другие классы эквивалентности
	{
		testName:    "Water число 576 строкой",
		description: "Проверка на класс эквивалентности число строкой",
		waterValue:  "576",
	},
	{
		testName:    "Water дробное число 11.6",
		description: "Проверка на класс эквивалентности число строкой",
		waterValue:  11.6,
	},
	{
		testName:    "Water null вместо числа",
		description: "Проверка на класс эквивалентности null",
		waterValue:  nil,
	},
	{
		testName:    "Water @ вместо числа",
		description: "Проверка на класс эквивалентности спецсимвол",
		waterValue:  "@",
	},
	{
		testName:    "Water строка abc вместо числа",
		description: "Проверка на класс эквивалентности буквы",
		waterValue:  "abc",
	},
}

func RunWaterTests(tcId int64) int64 {
	for _, td := range waterTests {
		if td.skip {
			fmt.Printf("Skipped test [%s].\n", td.testName)
			tcId++
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
		fmt.Printf("Running test [%s].\n", td.testName)
		err := RunTest(req, fmt.Sprintf("output/%d.png", tcId))
		if err != nil {
			fmt.Printf("Error on making screenshot in test [%s]. Error: %s", td.testName, err.Error())
		}
		tcId++
	}
	return tcId
}
