package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type NameRecord struct {
	fname string
	lname string
}

func main() {
	sli := make([]NameRecord, 0, 20)
	var location string
	fmt.Println("Enter your file location: ")
	fmt.Scan(&location)
	file, err := os.Open(location)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fname := strings.Fields(scanner.Text())[0]
		lname := strings.Fields(scanner.Text())[1]
		if len(fname) > 20 {
			fname = fname[:20]
		}
		if len(lname) > 20 {
			lname = lname[:20]
		}
		sli = append(sli, NameRecord{fname: fname, lname: lname})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(sli); i++ {
		fmt.Printf("First Name: %s, Last Name: %s\n", sli[i].fname, sli[i].lname)
	}
}
