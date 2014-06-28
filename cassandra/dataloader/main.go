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

	if len(args) != 3 {
		fmt.Println("Usage: dataloader {keyspace} {entries file} {words2entries file}")
		return
	}

	keyspace := args[0]
	source := args[1]
	words2source := args[2]

	// check if the source file exists
	sourceFile, err := os.Open(source)
	if err != nil {
		fmt.Println("Error reading input file: " + source)
		return
	}
	defer sourceFile.Close()

	// check if the words2source file exists
	words2SourceFile, err := os.Open(words2source)
	if err != nil {
		fmt.Println("Error reading: " + words2source)
		return
	}

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

		dataloader.ProcessWord(headword, content)
		line, e = Readln(reader)
	}

	reader = bufio.NewReader(words2SourceFile)
	line, e = Readln(reader)
	for e == nil {

		entries := strings.Split(line, "\t")
		if len(entries) != 4 {
			fmt.Println("Bad line entry!")
			continue
		}

		display := entries[3]
		headword := entries[2]

		dataloader.ProcessLookup(display, headword)
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
