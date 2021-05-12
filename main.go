package main

import (
	"fmt"
	"io"
	"os"

	"norm/capper"
	"norm/padding"
	"norm/timeshift"
	"norm/validator"

	"github.com/gocarina/gocsv"
)

type timeData struct {
	Timestamp     timeshift.DateTime `csv:"Timestamp"`
	Address       validator.Address  `csv:"Address"`
	Zipcode       padding.ZipCode    `csv:"ZIP"`
	FullName      capper.Uppercase   `csv:"FullName"`
	FooDuration   string             `csv:"FooDuration"`
	BarDuration   string             `csv:"BarDuration"`
	TotalDuration string             `csv:"TotalDuration"`
	Notes         validator.Notes    `csv:"Notes"`
}

func main() {

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		return gocsv.LazyCSVReader(in) // Allows use of quotes in CSV
	})

	dataFile := os.Stdin

	defer dataFile.Close()

	timeKeeps := []*timeData{}

	if err := gocsv.UnmarshalFile(dataFile, &timeKeeps); err != nil {
		panic(err)
	}

	csvContent, err := gocsv.MarshalString(&timeKeeps)

	if err != nil {
		panic(err)
	}

	fmt.Println(csvContent)
}
