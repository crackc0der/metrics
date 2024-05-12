package main

import (
	"log"
	"log/slog"
	"net"
	"os"
	product_grpc "server/api/note_v1"
	"server/config"
	"server/internal/category"
	"server/internal/product"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func Run() {
	config, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

	dsn := config.GetDsn()

	dbConn, err := sqlx.Connect(config.DataBase.DBType, dsn)
	if err != nil {
		log.Fatal(err)
	}

	productRepository := product.NewRepository(dbConn)
	productService := product.NewService(productRepository)
	productEndpoint := product.NewEndpoint(productService, logger)

	categoryRepository := category.NewRepository(dbConn)
	categoryService := category.NewService(categoryRepository)
	categoryEndpoint := category.NewEndpoint(categoryService, logger)

	conn, err := net.Listen("tcp", config.Host.HostPort)
	if err != nil {
		log.Fatal("error in start grpc server: %w", err)
	}

	serv := grpc.NewServer()

	reflection.Register(serv)
	product_grpc.RegisterProductServer(serv, productEndpoint)
	product_grpc.RegisterCategoryServer(serv, categoryEndpoint)

	if err := serv.Serve(conn); err != nil {
		log.Fatal(err)
	}
}
