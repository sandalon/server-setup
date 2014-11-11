package main

import (
	"./dataloader"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// command line args without the prog
	args := os.Args[1:]

	if len(args) != 7 {
		fmt.Println("Usage: dataloader {server} {keyspace} {entries file} {words2entries file} {words2title file} {words2meta file} {all|word|lookup|title|meta}")
		return
	}

	server := args[0]
	keyspace := args[1]
	source := args[2]
	words2source := args[3]
	words2titleSource := args[4]
	words2metaSource := args[5]
	activity := args[6]

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

	// check if the words2titleSource file exists
	words2title, err := os.Open(words2titleSource)
	if err != nil {
		fmt.Println("Error reading: " + words2titleSource)
		return
	}

	words2meta, err := os.Open(words2metaSource)
	if err != nil {
		fmt.Println("Error reading: " + words2metaSource)
		return
	}

	dataloader.Initialize(server, keyspace)

	if activity == "all" || activity == "word" {
		fmt.Println("Starting Words...")
		reader := bufio.NewReader(sourceFile)
		line, e := Readln(reader)
		maxBatchSize := 5000
		currentBatchSize := 0
		for e == nil {
			entries := strings.Split(line, "\t")
			
			if len(entries) != 4 {
				fmt.Println("Bad line entry!")
				fmt.Println(len(entries))
				return
			}

			currentBatchSize += 1

			headword := strings.Replace(entries[1], "\"", "", -1)
			content := strings.Replace(entries[3], "\"", "", -1)

			dataloader.ProcessWord(headword, content)

			if currentBatchSize == maxBatchSize {
				dataloader.ProcessBatch()
				currentBatchSize = 0
			}

			line, e = Readln(reader)
		}

		if currentBatchSize > 0 {
			dataloader.ProcessBatch()
			currentBatchSize = 0
		}
	}

	if activity == "all" || activity == "lookup" {
		fmt.Println("Starting Lookups...")
		reader := bufio.NewReader(words2SourceFile)
		line, e := Readln(reader)
		for e == nil {

			entries := strings.Split(line, "\t")
			if len(entries) != 4 {
				fmt.Println("Bad line entry!")
				continue
			}

			display := strings.Replace(entries[3], "\"", "", -1)
			//headword := entries[2]
			headword := strings.Replace(entries[2], "\"", "", -1)

			dataloader.ProcessLookup(display, headword)
			line, e = Readln(reader)
		}

		dataloader.ProcessBatch()
	}

	if activity == "all" || activity == "title" {
		fmt.Println("Starting Titles...")
		reader := bufio.NewReader(words2title)
		line, e := Readln(reader)
		for e == nil {

			entries := strings.Split(line, "\t")
			if len(entries) != 2 {
				fmt.Println("Bad line entry!")
				continue
			}

			display := strings.Replace(entries[0], "\"", "", -1)
			title := strings.Replace(entries[1], "\"", "", -1)

			dataloader.ProcessTitle(display, title)
			line, e = Readln(reader)
		}

		dataloader.ProcessBatch()
	}

	if activity == "all" || activity == "meta" {
		fmt.Println("Starting Meta...")
		reader := bufio.NewReader(words2meta)
		line, e := Readln(reader)
		for e == nil {

			entries := strings.Split(line, "\t")
			if len(entries) != 4 {
				fmt.Println("Bad line entry!")
				continue
			}

			display := strings.Replace(entries[0], "\"", "", -1)
			description := strings.Replace(entries[1], "\"", "", -1)
			keywords := strings.Replace(entries[2], "\"", "", -1)
			copyright := strings.Replace(entries[3], "\"", "", -1)

			dataloader.ProcessMeta(display, description, keywords, copyright)
			line, e = Readln(reader)
		}

		dataloader.ProcessBatch()
	}

	fmt.Println("==========")
	fmt.Println("Error Count: " + strconv.Itoa(dataloader.GetErrorCount()))

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
