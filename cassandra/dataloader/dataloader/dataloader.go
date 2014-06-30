package dataloader

import (
	"fmt"
	"github.com/gocql/gocql"
	"time"
)

var session *gocql.Session
var cluster *gocql.ClusterConfig
var errorCount int

func Initialize(keyspace string) {
	cluster = gocql.NewCluster("localhost")
	cluster.Keyspace = keyspace
	errorCount = 0

	var err error
	session, err = cluster.CreateSession()
	if err != nil {
		fmt.Println("Error creating session")
		errorCount = errorCount + 1
		return
	}
}

func CleanUp() {
	session.Close()
}

func GetErrorCount() int {
	return errorCount
}

func ProcessWord(headword string, content string) {
	fmt.Println("Processing " + headword)
	if err := session.Query("INSERT INTO word (headword, content) VALUES (?, ?)",
		headword, content).Exec(); err != nil {
		errorCount = errorCount + 1
		fmt.Println(err)
		return
	}
}

func ProcessLookup(display string, headword string) {
	fmt.Println("Processing Display " + display)

	var content string
	if err := session.Query(`SELECT content FROM word WHERE headword = ? LIMIT 1`,
        headword).Consistency(gocql.One).Scan(&content); err != nil {
				errorCount = errorCount + 1
        fmt.Println("Error fetching content for headword: " + headword)
				fmt.Println(err)
				return
    }

	time.Sleep(10 * time.Millisecond)

	if err := session.Query("INSERT INTO lookup (wordformDisplay, headword, content) VALUES (?, ?, ?)",
		display, headword, content).Exec(); err != nil {
		errorCount = errorCount + 1
		fmt.Println("Error saving lookup data")
		fmt.Println(err)
		return
	}
}
