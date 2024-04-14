package main

import (
	"fmt"
	"os"
)

func main() {
	os.Mkdir("output", 0777)
	tcId := int64(1)
	fmt.Println("Выполняю тесты по CO2")
	tcId = RunCO2Tests(tcId)
	fmt.Println("Выполняю тесты по Water")
	tcId = RunWaterTests(tcId)
	fmt.Println("Выполняю тесты по Energy")
	tcId = RunEnergyTests(tcId)
	fmt.Println(tcId)
}
