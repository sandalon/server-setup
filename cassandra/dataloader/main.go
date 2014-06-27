package main

import (
	"./dataloader"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// command line args without the prog
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("Usage: dataloader {keyspace} {source file}")
		return
	}

	keyspace := args[0]
	source := args[1]

	// check if the options file exists
	sourceFile, err := os.Open(source)
	if err != nil {
		fmt.Println("Error reading input file: " + source)
		return
	}
	defer sourceFile.Close()

	dataloader.Initialize(keyspace)

	reader := bufio.NewReader(sourceFile)
	line, e := Readln(reader)
	for e == nil {

		entries := strings.Split(line, "\t")
		if len(entries) != 4 {
			fmt.Println("Bad line entry!")
			continue
		}

		headword := strings.Replace(entries[1], "\"", "", -1)
		content := entries[3]

		dataloader.Process(headword, content)
		line, e = Readln(reader)
	}

	dataloader.CleanUp()
}

// http://stackoverflow.com/questions/6141604/go-readline-string
func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}
