package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Entry struct {
	Tel        string
	Name       string
	SurName    string
	LastAccess string
}

var data []Entry
var index map[string]int

var CSVFILEPATH = "./csv.data"

func readCSVFile(path string) error {
	if _, err := os.Stat(path); err != nil {
		return err
	}

	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return err
	}

	for _, line := range lines {
		temp := Entry{
			Tel:        line[0],
			Name:       line[1],
			SurName:    line[2],
			LastAccess: line[3],
		}

		data = append(data, temp)

	}

	return nil

}

func saveCSV(filepath string) error {
	csvfile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer csvfile.Close()

	csvwriter := csv.NewWriter(csvfile)
	for _, row := range data {
		temp := []string{
			row.Tel, row.Name, row.SurName, row.LastAccess,
		}
		_ = csvwriter.Write(temp)
	}
	csvwriter.Flush()
	return nil
}

func createIndex() error {
	index = make(map[string]int)
	for i, k := range data {
		key := k.Tel
		index[key] = i
	}
	return nil
}

// If it returns nil, there was an error
func initS(T, N, S string) *Entry {
	if T == "" {
		return nil
	}

	LastAcces := strconv.FormatInt(time.Now().Unix(), 10)
	return &Entry{Tel: T, Name: N, SurName: S, LastAccess: LastAcces}
}

func insert(pS *Entry) error {
	if _, ok := index[(*pS).Tel]; ok {
		return fmt.Errorf("%s already exist!", pS.Tel)
	}

	data = append(data, *pS)
	_ = createIndex()

	if err := saveCSV(CSVFILEPATH); err != nil {
		return err
	}
	return nil
}

func deleteEntry(key string) error {
	i, ok := index[key]
	if !ok {
		return fmt.Errorf("%s cannot be found!", key)
	}
	data = append(data[:i], data[i+1:]...)

	delete(index, key)

	if err := saveCSV(CSVFILEPATH); err != nil {
		return err
	}
	return nil
}

// If it returns nil, Entry was`t found
func search(key string) *Entry {
	i, ok := index[key]
	if !ok {
		return nil
	}

	data[i].LastAccess = strconv.FormatInt(time.Now().Unix(), 10)
	_ = saveCSV(CSVFILEPATH)
	return &data[i]
}

func printList() {
	for _, v := range data {
		fmt.Println(v)
	}
}

func matchTel(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^\+[0-9]+?`)
	return re.Match(t)
}

func clearTel(s string) string {
	return strings.Map(func(r rune) rune {
		if r == '-' || r == '(' || r == ')' || r == ' ' {
			return -1
		}
		return r
	}, s)
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Usage: insert|delete|search|list <arguments>")
		return
	}

	if _, err := os.Stat(CSVFILEPATH); err != nil {
		fmt.Println("Creating", CSVFILEPATH)
		f, err := os.Create(CSVFILEPATH)
		if err != nil {
			fmt.Println(err)
			return
		}
		f.Close()

	}

	fileinfo, err := os.Stat(CSVFILEPATH)
	if err != nil {
		fmt.Println(err)
		return
	}

	mode := fileinfo.Mode()
	if !mode.IsRegular() {
		fmt.Println(CSVFILEPATH, "is not Reguarl file!")
		return
	}

	if err := readCSVFile(CSVFILEPATH); err != nil {
		fmt.Println(err)
		return
	}

	if err = createIndex(); err != nil {
		fmt.Println(err)
		return
	}

	switch os.Args[1] {
	case "insert":
		if len(os.Args) != 5 { //file insert Tel Name Surname
			fmt.Println("Usage: insert Telephone Name Surname")
			return
		}

		t := clearTel(os.Args[2])
		if !matchTel(t) {
			fmt.Println("Not a valid telephone number:", t)
			return
		}

		temp := initS(os.Args[2], os.Args[3], os.Args[4])
		if temp == nil {
			fmt.Println("Init error")
			return
		}
		if err := insert(temp); err != nil {
			fmt.Println(err)
			return
		}

		return
	case "delete":
		if len(os.Args) != 3 {
			fmt.Println("Usage: delete Telephone")
			return
		}

		t := clearTel(os.Args[2])
		if !matchTel(t) {
			fmt.Println("Not a valid telephone number:", t)
			return
		}

		if err := deleteEntry(t); err != nil {
			fmt.Println(err)
			return
		}

		return

	case "search":
		if len(os.Args) != 3 {
			fmt.Println("Usage: search Telephone")
			return
		}

		t := clearTel(os.Args[2])
		if !matchTel(t) {
			fmt.Println("Not a valid telephone number:", t)
			return
		}

		ent := search(t)
		if ent == nil {
			fmt.Println("Number not found!:", t)
			return
		}

		fmt.Println(*ent)

		return

	case "lsit":
		printList()
		return

	default:
		fmt.Println("Not a valid option")
		return
	}
}
