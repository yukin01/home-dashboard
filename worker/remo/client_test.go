package remo_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/yukin01/home-dashboard/worker/remo"
)

func TestGetDevices(t *testing.T) {
	token := os.Getenv("REMO_ACCESS_TOKEN")
	if token == "" {
		t.Fatal(errors.New("token is missing"))
	}
	c := remo.NewClient(token)
	ctx := context.Background()
	ds, err := c.GetDevices(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ds)
}
