package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/firestore"

	"github.com/yukin01/home-dashboard/worker/remo"
)

func handler(w http.ResponseWriter, r *http.Request) {
	token := os.Getenv("REMO_ACCESS_TOKEN")
	if token == "" {
		fmt.Println("[Error] token is missing")
		return
	}
	c := remo.NewClient(token)

	ctx := r.Context()
	devices, err := c.GetDevices(ctx)
	if err != nil {
		fmt.Println("[Error]", err)
		return
	}
	fmt.Printf("[Info] %#v\n", devices)

	if err = add(ctx, devices); err != nil {
		fmt.Println("[Error]", err)
	}
}

func add(ctx context.Context, devices []*remo.Device) error {
	client, err := firestore.NewClient(ctx, "yukin01-home-dashboard")
	if err != nil {
		return fmt.Errorf("failed to create firestore client: %s", err)
	}

	defer client.Close()

	for _, d := range devices {
		doc := client.Collection("devices").Doc(d.ID)
		event := &struct {
			remo.NewestEvents
			createdAt time.Time `firestore:"created_at"`
		}{d.NewestEvents, time.Now()}
		_, _, err = doc.Collection("events").Add(ctx, event)
		if err != nil {
			return fmt.Errorf("Failed adding event: %v", err)
		}
	}
	return nil
}
