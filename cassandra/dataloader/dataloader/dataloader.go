package dataloader

import (
	"fmt"
	"github.com/gocql/gocql"
)

var session *gocql.Session
var cluster *gocql.ClusterConfig

func Initialize(keyspace string) {
	cluster = gocql.NewCluster("localhost")
	cluster.Keyspace = keyspace
	var err error
	session, err = cluster.CreateSession()
	if err != nil {
		fmt.Println("Error creating session")
		return
	}
}

func CleanUp() {
	session.Close()
}

func Process(headword string, content string) {
	fmt.Println("Processing " + headword)
	if err := session.Query("INSERT INTO word (display, content) VALUES (?, ?)",
		headword, content).Exec(); err != nil {
		fmt.Println("Error!")
		return
	}
}
