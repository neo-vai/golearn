package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Record struct {
	Name       string
	Surname    string
	Number     string
	LastAccess string
}

var MyData = []Record{}

func readCSVFile(filepath string) ([][]string, error) {
	if _, err := os.Stat(filepath); err != nil {
		return nil, err
	}

	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}

func saveCSVFile(filepath string) error {
	csvfile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer csvfile.Close()

	csvwriter := csv.NewWriter(csvfile)
	csvwriter.Comma = '\t'
	for _, row := range MyData {
		temp := []string{row.Name, row.Surname, row.Number, row.LastAccess}
		if err = csvwriter.Write(temp); err != nil {
			log.Println("WARNING:", err)
		}
	}
	csvwriter.Flush()
	return nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Used: csvData input output")
		return
	}

	input := os.Args[1]
	output := os.Args[2]

	lines, err := readCSVFile(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, line := range lines {
		temp := Record{
			Name:       line[0],
			Surname:    line[1],
			Number:     line[2],
			LastAccess: line[3],
		}
		MyData = append(MyData, temp)
		fmt.Println(temp)
	}
	if err = saveCSVFile(output); err != nil {
		fmt.Println(err)
		return
	}
}
