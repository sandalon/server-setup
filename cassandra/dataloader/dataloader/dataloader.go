package dataloader

import (
	"fmt"
	"github.com/gocql/gocql"
	"os"
)

var session *gocql.Session
var cluster *gocql.ClusterConfig
var errorCount int
var batch *gocql.Batch
var batchSize int

func Initialize(server string, keyspace string) {
	cluster = gocql.NewCluster(server)
	cluster.Keyspace = keyspace
	cluster.Timeout = 60000000000 // 60 seconds in ns

	var retry gocql.RetryPolicy
	retry.NumRetries = 2
	cluster.RetryPolicy = retry
	cluster.NumConns = 1
	errorCount = 0
	batchSize = 0
	batch = gocql.NewBatch(gocql.LoggedBatch)

	var err error
	session, err = cluster.CreateSession()
	if err != nil {
		fmt.Println("Error creating session")
		os.Exit(1)
	}
}

func ProcessBatch() {
	err := session.ExecuteBatch(batch)
	if err != nil {
		errorCount += 1
		fmt.Println("Error processing batch")
		fmt.Println(err)
		os.Exit(1)
	}
	batchSize = 0
	batch = gocql.NewBatch(gocql.LoggedBatch)
}

func CleanUp() {
	session.Close()
}

func GetErrorCount() int {
	return errorCount
}

func ProcessWord(headword string, content string) {
	//fmt.Println("Processing " + headword)
	batch.Query("INSERT INTO word (headword, content) VALUES (?, ?)",
		headword, content)

	return
}

func ProcessTitle(display string, title string) {
	batchSize += 1

	batch.Query("UPDATE lookup set title = ? where wordformdisplay = ?",
		title, display)

	if batchSize == 5000 {
		ProcessBatch()
	}
}

func ProcessMeta(display string, description string, keywords string, copyright string) {
	batchSize += 1

	batch.Query("UPDATE lookup set meta = ? where wordformdisplay = ?",
		description+keywords+copyright, display)

	if batchSize == 5000 {
		ProcessBatch()
	}
}

func ProcessLookup(display string, headword string) {
	fmt.Println("Processing Display " + headword)
	batchSize += 1

	var content string
	if err := session.Query(`SELECT content FROM word WHERE headword = ? LIMIT 1`,
		headword).Consistency(gocql.One).Scan(&content); err != nil {
		errorCount = errorCount + 1
		fmt.Println("Error fetching content for headword: " + headword)
		fmt.Println(err)
		errorCount += 1
		return
	}

	batch.Query("INSERT INTO lookup (wordformDisplay, headword, content) VALUES (?, ?, ?)",
		display, headword, content)

	if batchSize == 5000 {
		ProcessBatch()
	}

}
