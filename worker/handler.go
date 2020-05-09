package main

import (
	"context"
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
	rc = remo.NewClient(getEnv("REMO_ACCESS_TOKEN"))
	projectID = getEnv("PROJECT_ID")
	fmt.Println("[Info] get environment variables successfully")
}

func getEnv(key string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	panic(fmt.Errorf("%s is missing", key))
}
