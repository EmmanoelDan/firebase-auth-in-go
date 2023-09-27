package config

import (
	"context"

	// "cloud.google.com/go/firestore"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var Client *firebase.App

func InitializeFirestore() {
	logger := GetLogger("firestore")
	opt := option.WithCredentialsFile("path/to/serviceAccountKey.json")
	client, err := firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		logger.Errf("Firestore initialization error: %v", err)
		return
	}

	Client = client
}
