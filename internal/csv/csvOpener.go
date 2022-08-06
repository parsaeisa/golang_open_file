package csv

type Address struct {
	Firstname string `csv:"firstname"`
	Lastname  string `csv:"lastname"`
	Address   string `csv:"address"`
}
