package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"cloud.google.com/go/firestore"

	"github.com/yukin01/home-dashboard/worker/remo"
)

var (
	rc        *remo.Client
	projectID string
)

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	devices, err := rc.GetDevices(ctx)
	if err != nil {
		fmt.Println("[Error]", err)
		return
	}
	fmt.Printf("[Info] %#v\n", devices)

	if err = add(ctx, devices); err != nil {
		fmt.Println("[Error]", err)
	}
	fmt.Println("[Info] add events successfully")
}

func add(ctx context.Context, devices []*remo.Device) error {
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("failed to create firestore client: %s", err)
	}

	defer client.Close()

	for _, d := range devices {
		doc := client.Collection("devices").Doc(d.ID)
		events := NewEvents(d.NewestEvents)
		_, _, err = doc.Collection("events").Add(ctx, events)
		if err != nil {
			return fmt.Errorf("failed adding event: %s", err)
		}
	}
	return nil
}

func init() {
	token := os.Getenv("REMO_ACCESS_TOKEN")
	if token == "" {
		panic(errors.New("[Error] remo access token is missing"))
	}
	rc = remo.NewClient(token)

	projectID = os.Getenv("PROJECT_ID")
	if projectID == "" {
		panic(errors.New("[Error] project ID is missing"))
	}
	fmt.Println("[Info] get environment variables successfully")
}
