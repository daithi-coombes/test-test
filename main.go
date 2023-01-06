package main

import (
	"encoding/csv"
	"log"
	"math/big"
	"os"
	"strconv"

	"github.com/davecgh/go-spew/spew"
)

const AREA_CODE = "061"

func main() {

	records, err := loadCSVFile("./census.csv")
	if err != nil {
		log.Fatal(err)
	}

	// find the avg income for area 061
	// calculate percentage of missing data (I presume missing income)

	var totalIncome int
	var missingIncome int
	var total int

	for i, record := range records {
		if i == 0 || record[0] != AREA_CODE {
			continue
		}

		income, err := strconv.Atoi(record[2])
		if err != nil {
			log.Fatalf("Invalid CSV data on row %d", i+1)
		}
		total++

		if income == int(-1) {
			missingIncome++

			continue
		}

		totalIncome += income
	}

	spew.Dump(records)

	totalF := new(big.Float).SetInt(big.NewInt(int64(total)))
	missingIncomeF := new(big.Float).SetInt(big.NewInt(int64(missingIncome)))
	spew.Dump(total)
	spew.Dump(totalIncome)
	spew.Dump(missingIncome)

	percentageMissing := new(big.Float).Quo(totalF, missingIncomeF)
	percent := new(big.Float).Mul(percentageMissing, big.NewFloat(100))
	spew.Dump(percentageMissing)
	spew.Dump(percent)
}

func loadCSVFile(path string) ([][]string, error) {
	var records [][]string

	file, err := os.Open(path)
	if err != nil {
		return records, err
	}

	reader := csv.NewReader(file)
	records, err = reader.ReadAll()
	if err != nil {
		return records, err
	}

	return records, nil
}
