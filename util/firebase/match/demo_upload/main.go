package main

import (
	"bytes"
	"context"
	"flag"
	"log"

	firebase "firebase.google.com/go"

	fsc "github.com/FlowingSPDG/Got5/examples/firebase"
)

var (
	projectID = ""
	bucket    = ""
)

func main() {
	// Parse -projectID flag
	flag.StringVar(&projectID, "projectID", "", "Firebase project ID")
	flag.StringVar(&bucket, "bucket", "", "Firebase Storage Bucket")
	flag.Parse()

	// Get Firebase service
	ctx := context.Background()
	conf := &firebase.Config{
		ProjectID:     projectID,
		StorageBucket: bucket,
	}
	fb, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}

	// Get Controller connected to firestore
	up, err := fsc.NewDemoUploader(ctx, fb)
	if err != nil {
		log.Fatalln(err)
	}

	if err := up.Upload(ctx, "UNJR7KfmIC1heVjHNyLD", "test_file_name.dem", bytes.NewBuffer([]byte("INVALID DATA"))); err != nil {
		log.Fatalln(err)
	}
}
