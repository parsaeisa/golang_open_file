package cmd

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var openCsvCommand = &cobra.Command{
	Use:   "csv",
	Short: "This command is being used to open a csv file ",
	Run:   OpenCsv,
}

func OpenCsv(_ *cobra.Command, _ []string) {
	csvFile, err := os.Open("../addresses.csv")
	if err != nil {
		log.Fatalf("csv file could not be opened %v", err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		log.Fatalf("csv file lines could not be opened %v", err)
	}

	// this csvLines is a 2D array of strings
	for _, line := range csvLines {

	}
}
