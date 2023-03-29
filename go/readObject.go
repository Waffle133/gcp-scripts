package main

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"log"
	"os"
)

const (
	bucketName = ""
	projectID  = ""
	objectName = ""
)

func readObject() {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	var o *storage.ObjectHandle
	w := os.Stdout

	sum := 1
	for sum < 10000 {
		sum += 1
		o = client.Bucket(bucketName).Object(objectName)
		rc, err := o.Attrs(ctx)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Fprintf(w, "%v.- Custom time %v\n", sum, rc.CustomTime)

	}

	// fmt.Println("Successfully created the bucket.")
}

func main() {
	readObject()
}
