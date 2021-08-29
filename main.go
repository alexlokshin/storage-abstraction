package main

import (
	"log"

	"github.com/alexlokshin/storage-abstraction/cloud"
)

func main() {
	client, err := cloud.CreateClient("aws")
	if err != nil {
		log.Fatalf("Unable to create client: %s\n", err.Error())
	}
	client.UploadFile("testbucket", "README.md")
}
