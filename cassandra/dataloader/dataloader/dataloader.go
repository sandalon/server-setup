package dataloader

import (
	"fmt"
	"github.com/gocql/gocql"
	//"time"
)

var session *gocql.Session
var cluster *gocql.ClusterConfig
var errorCount int
var batch *gocql.Batch
var batchSize int

func Initialize(server string, keyspace string) {
	cluster = gocql.NewCluster(server)
	cluster.Keyspace = keyspace
	errorCount = 0
	batchSize = 0

	var err error
	session, err = cluster.CreateSession()
	if err != nil {
		fmt.Println("Error creating session")
		errorCount = errorCount + 1
		panic(err)
		return
	}

	batch = gocql.NewBatch(gocql.LoggedBatch)
}

func ProcessBatch() {
	//err := session.ExecuteBatch(batch)
	//if err != nil {
	//	errorCount += 1
	//	fmt.Println("error processing batch: ")
	//	fmt.Println(err)
	//}
	//batch = gocql.NewBatch(gocql.LoggedBatch)
	//time.Sleep(10000)
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

	//fmt.Println("INSERT INTO word (headword, content) VALUES (" + headword + "," + content + ")")
	//fmt.Printf("%s\t%s\n", headword, content)

	return
}

func ProcessLookup(display string, headword string) {
	//fmt.Println("Processing Display " + display)
	batchSize += 1

	var content string
	if err := session.Query(`SELECT content FROM word WHERE headword = ? LIMIT 1`,
        headword).Consistency(gocql.One).Scan(&content); err != nil {
				errorCount = errorCount + 1
        fmt.Println("Error fetching content for headword: " + headword)
				fmt.Println(err)
				return
    }

	batch.Query("INSERT INTO lookup (wordformDisplay, headword, content) VALUES (?, ?, ?)",
		display, headword, content)

	fmt.Printf("%s\t%s\t%s\n", display, headword, content)

	if(batchSize == 100) {
		//fmt.Println("Processing display batch")
		//ProcessBatch()
	}

}
