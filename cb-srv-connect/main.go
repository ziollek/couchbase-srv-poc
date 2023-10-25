package main

import (
	"log"
	"time"
        "os"

	"github.com/couchbase/gocb/v2"
)


func main() {
	// Uncomment following line to enable logging
	// gocb.SetLogger(gocb.VerboseStdioLogger())
	commandArgs := os.Args[1:]
	// Update this to your cluster details
        if len(commandArgs) != 4 {
	   log.Printf("Expected (4 != %d)  arguments go run . connection-string bucket user pass\nExample:\n", len(commandArgs))
           log.Printf("go run . couchbase://10.0.0.1:11210 default default ***\nOr by srv domain:\n")
           log.Printf("go run . couchbase://mycluster.consul.suffix  dafault default ***\n")
           os.Exit(1)
        }
	connectionString := commandArgs[0]
	bucketName := commandArgs[1]
	username := commandArgs[2]
	password := commandArgs[3]
        log.Printf("connecting using following parameters connection=%s, bucket=%s, user=%s, pass=***", connectionString, bucketName, username)

	options := gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: username,
			Password: password,
		},
	}


	// Initialize the Connection
	cluster, err := gocb.Connect(connectionString, options)
	if err != nil {
		log.Fatal(err)
	}

	bucket := cluster.Bucket(bucketName)

	err = bucket.WaitUntilReady(60*time.Second, nil)
	if err != nil {
		log.Fatal(err)
	}
        log.Println("OK!")
}
