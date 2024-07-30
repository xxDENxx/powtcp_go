package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/xxDENxx/powtcp_go/server/internal/pkg/cache"
	"github.com/xxDENxx/powtcp_go/server/internal/pkg/config"
	"github.com/xxDENxx/powtcp_go/server/internal/pkg/server"
)

func main() {
	// future integrations with ELK
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	log.Println("Starting")
	ctx := context.Background()

	// Read config
	config := config.Config{}
	err := config.ReadFromEnv()
	if err != nil {
		log.Fatal("CONFIG: ", err)
	}

	// Init cache
	cache, err := cache.InitRedisCache(ctx, &config)
	if err != nil {
		log.Fatal("CONFIG: ", err)
	}
	// Fill context with data
	ctx = context.WithValue(ctx, "config", &config)
	ctx = context.WithValue(ctx, "cache", cache)

	// run server
	err = server.Run(ctx)
	if err != nil {
		log.Fatal("SERVER: ", err)
	}
}