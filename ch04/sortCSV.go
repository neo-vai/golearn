package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"sort"
)

type F1 struct {
	Name       string
	LastName   string
	Tel        string
	LastAccess string
}

type F2 struct {
	Name       string
	LastName   string
	CityNumber string
	Tel        string
	LastAccess string
}

func (f F1) GetName() string {
	return f.Name
}

func (f F1) GetLastName() string {
	return f.LastName
}

func (f F1) GetTel() string {
	return f.Tel
}

func (f F1) GetLastAccess() string {
	return f.LastAccess
}

func (f F2) GetName() string {
	return f.Name
}

func (f F2) GetLastName() string {
	return f.LastName
}

func (f F2) GetTel() string {
	return f.Tel
}

func (f F2) GetLastAccess() string {
	return f.LastAccess
}

type CsvField interface {
	GetName() string
	GetLastName() string
	GetTel() string
	GetLastAccess() string
}

type csvdata []CsvField

func (c csvdata) Len() int {
	return len(c)
}

func (c csvdata) Less(i, j int) bool {
	ie := c[i]
	je := c[j]

	return len(ie.GetName())+len(ie.GetLastName()) <= len(je.GetName())+len(je.GetLastName())
}

func (c csvdata) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

var Data csvdata = csvdata{}

func ReadCSVFile(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return err
	}
	if !info.Mode().IsRegular() {
		return errors.New("File is not regular")
	}

	csvReader := csv.NewReader(file)

	firstLine := true
	var format1 bool

	lines, err := csvReader.ReadAll()
	if err != nil {
		return err
	}

	for _, line := range lines {
		if firstLine {
			if len(line) == 4 {
				format1 = true
			} else if len(line) == 5 {
				format1 = false
			} else {
				return errors.New("Unknown csv format")
			}
			firstLine = false
		}

		if format1 {
			if len(line) == 4 {
				temp := F1{
					Name:       line[0],
					LastName:   line[1],
					Tel:        line[2],
					LastAccess: line[3],
				}
				Data = append(Data, temp)
			}
		} else {
			if len(line) == 5 {
				temp := F2{
					Name:       line[0],
					LastName:   line[1],
					CityNumber: line[2],
					Tel:        line[3],
					LastAccess: line[4],
				}
				Data = append(Data, temp)
			}
		}
	}
	return nil
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Not enough arguments")
		return
	}

	var cnt int
	for _, arg := range os.Args[1:] {
		if err := ReadCSVFile(arg); err != nil {
			fmt.Printf("Error to read file: %s : %s\n", arg, err)
		} else {
			cnt++
		}
	}

	if cnt <= 0 {
		fmt.Println("No one correct file path")
		return
	}

	sort.Sort(Data)

	for i, l := range Data {
		fmt.Printf("%d: %s, %s, %s\n", i, l.GetName(), l.GetLastName(), l.GetTel())
	}
}
