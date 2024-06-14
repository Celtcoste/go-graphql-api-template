package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/celtcoste/go-graphql-api-template/internal"
	"github.com/celtcoste/go-graphql-api-template/internal/database"
	"github.com/celtcoste/go-graphql-api-template/internal/graph"
	"github.com/celtcoste/go-graphql-api-template/internal/graph/generated"
	"github.com/celtcoste/go-graphql-api-template/internal/repository"
	"github.com/celtcoste/go-graphql-api-template/internal/server"
	"github.com/celtcoste/go-graphql-api-template/pkg/cloud/cloudLogger"
	"github.com/celtcoste/go-graphql-api-template/pkg/util/translation"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func loadConfiguration() *internal.Configuration {
	directory := flag.String("configDir", "./", "massage.yaml file directory")
	flag.Parse()
	configuration, err := internal.NewConfiguration(*directory)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return configuration
}

func loadCloudLogger() *cloudLogger.Logger {
	logger, err := cloudLogger.NewLogger()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return logger
}

func loadDatabasePool(configuration *database.Configuration) *sqlx.DB {
	pool, err := database.NewDatabasePool(configuration)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return pool
}

func main() {
	ctx := context.Background()
	godotenv.Load()
	configuration := loadConfiguration()

	// NOTE: setup logging
	logger := loadCloudLogger()
	defer logger.Sync()

	// NOTE: setup database pool
	pool := loadDatabasePool(configuration.Database)
	defer pool.Close()

	// Note: setup repositories
	repositories := repository.NewContainer(pool)

	// Note: setup translator
	translator := translation.NewTranslator("locales", []string{"en_US", "fr_FR"})

	// Instantiate resolver
	resolver := graph.NewResolver(repositories, translator)
	config := generated.Config{Resolvers: resolver}
	gqlServer := server.NewGraphQLServer(config, logger, repositories)

	// NOTE: create server and attach handlers to internal router.
	api := server.NewApiTemplateServer(configuration.Server, logger)
	server.SetupGraphQLRoutes(
		repositories,
		configuration.Server.Playground,
		gqlServer,
		api.Router,
	)
	// NOTE: start server
	api.Start(ctx)
	os.Exit(0)
}
