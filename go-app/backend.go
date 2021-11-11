package main

import (
	"context"
	"log"

	cloudstore "cloud.google.com/go/storage"
	"github.com/kramphub/kiya/backend"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/cloudkms/v1"
)

func newKMSBackend() (backend.Backend, error) {
	// Create the KMS client
	defaultClient, err := google.DefaultClient(context.Background(), cloudkms.CloudPlatformScope)
	if err != nil {
		return nil, err
	}
	kmsService, err := cloudkms.New(defaultClient)
	if err != nil {
		return nil, err
	}
	// Create the Bucket client
	storageService, err := cloudstore.NewClient(context.Background())
	if err != nil {
		log.Fatalf("failed to create client [%v]", err)
	}
	return backend.NewKMS(kmsService, storageService), nil
}
