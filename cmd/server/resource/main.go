package main

import (
	"context"
	"flag"
	"log"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/yyyoichi/OhAuth0.1/internal/resource"
)

func main() {
	ctx := context.Background()
	l := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(l)

	envPath := flag.String("source", "", "env file")
	flag.Parse()
	if envPath != nil {
		if err := godotenv.Load(*envPath); err != nil {
			panic(err)
		}
		slog.Info("read env file", slog.String("path", *envPath))
	}

	var dbport string
	if dbport = os.Getenv("DATABASE_SERVER_PORT"); dbport == "" {
		panic("no required env found")
	}
	service, err := resource.NewService(ctx, resource.Config{
		DatabaseServerURL: "http://localhost:" + dbport,
	})
	if err != nil {
		log.Fatal(err)
	}

	var port string
	if port = os.Getenv("RESOURCE_SERVER_PORT"); port == "" {
		panic("no required env found")
	}
	router := resource.SetupRouter(service)
	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
