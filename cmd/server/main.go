package main

import (
	"log"

	"github.com/Raphael-Jin/Go-gRPC-Services/internal/db"
	"github.com/Raphael-Jin/Go-gRPC-Services/internal/rocket"
)

func Run() error {
	// initialzing grpc server
	log.Println("Starting up Rocket gRPC Service")
	rocketStore, err := db.New()
	if err != nil {
		return err
	}
	err = rocketStore.Migrate()
	if err != nil {
		log.Println("Failed to run migrations")
		return err
	}
	_ = rocket.New(rocketStore)
	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
