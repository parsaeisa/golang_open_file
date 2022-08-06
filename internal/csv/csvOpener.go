package csv

import (
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

type Address struct {
	Firstname string `csv:"firstname"`
	Lastname  string `csv:"lastname"`
	Address   string `csv:"address"`
}

type csvOpener struct {
	path string
}

type CsvOpener interface {
	Open()
}

func NewCsvopener(path string) CsvOpener {
	return &csvOpener{
		path: path,
	}
}

func (receiver *csvOpener) Open() {
	csvFile, err := os.Open("../addresses.csv")
	if err != nil {
		log.Fatalf("csv file could not be opened %v", err)
	}
	defer csvFile.Close()

	var Addresses []*Address

	err = gocsv.Unmarshal(csvFile, Addresses)
	if err != nil {
		log.Fatalf("Error happend while unmarshaling %v", err)
	}

	for _, address := range Addresses {
		print(address.Lastname)
	}
}
