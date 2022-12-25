package db

import (
	"fmt"
	"os"

	"github.com/Raphael-Jin/Go-gRPC-Services/internal/rocket"
	"github.com/jmoiron/sqlx"
)

type Store struct {
	db *sqlx.DB
}

// return a new store or err
func New() (Store, error) {
	// get env varibale from docker
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")

	connectString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		dbHost,
		dbPort,
		dbUsername,
		dbTable,
		dbPassword,
	)

	db, err := sqlx.Connect("postgres", connectString)
	if err != nil {
		return Store{}, err
	}
	return Store{
		db: db,
	}, nil
}

func (s Store) GetRocketByID(id string) (rocket.Rocket, error) {
	return rocket.Rocket{}, nil
}

func (s Store) InsertRocket(rkt rocket.Rocket) (rocket.Rocket, error) {
	return rocket.Rocket{}, nil
}

func (s Store) DeleteRocket(id string) error {
	return nil
}
