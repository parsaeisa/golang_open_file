package cmd

import (
	"github.com/spf13/cobra"

	"github.com/parsaeisa/golang_open_file/internal/csv"
)

var openCsvCommand = &cobra.Command{
	Use:   "csv",
	Short: "This command is being used to open a csv file ",
	Run:   OpenCsv,
}

func OpenCsv(_ *cobra.Command, _ []string) {
	opener := csv.NewCsvopener("../addresses.csv")
	opener.Open()
}
