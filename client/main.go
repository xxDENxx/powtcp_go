package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/xxDENxx/powtcp_go/client/internal/pkg/client"
)

func main() {
	// future integrations with ELK
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	log.Println("Starting")

	// Read config
	config := client.Config{}
	err := config.ReadFromEnv()
	if err != nil {
		log.Fatal("CONFIG: ", err)
	}

	// init context
	ctx := context.Background()
	ctx = context.WithValue(ctx, "config", &config)

	// run client
	err = client.Run(ctx)
	if err != nil {
		log.Fatal("CLIENT: ", err)
	}
}