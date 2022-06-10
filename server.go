package main

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"hyneo-backend/graph"
	"hyneo-backend/graph/generated"
	"hyneo-backend/internal/category/db"
	"hyneo-backend/internal/config"
	db2 "hyneo-backend/internal/item/db"
	db3 "hyneo-backend/internal/promo/db"
	"hyneo-backend/pkg/logging"
	"hyneo-backend/pkg/mysql"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "80"

func main() {
	logging.Init()
	log := logging.GetLogger()

	cfg := config.GetConfig()

	database := mysql.NewClient(context.TODO(), 5, cfg.MySQL)

	categoryRepository := db.NewRepository(*database, &log)
	itemRepository := db2.NewRepository(*database, &log)
	promoRepository := db3.NewRepository(*database, &log)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"https://hyneo.ru", "http://localhost:3000"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		ItemRepository:     itemRepository,
		CategoryRepository: categoryRepository,
		PromoRepository:    promoRepository,
	}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
