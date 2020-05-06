package main

import (
	"fmt"
	"net/http"
	"os"

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
	fmt.Println("[Info]", devices)
}
