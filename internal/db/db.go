package db

import (
	"errors"
	"fmt"
	"log"
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

// retrive a rocket from the database by id
func (s Store) GetRocketByID(id string) (rocket.Rocket, error) {
	var rkt rocket.Rocket
	row := s.db.QueryRow(
		// "SELECT id FROM rockets where id=$1;",
		`SELECT id, type, name FROM rockets where id=$1;`,
		id,
	)

	err := row.Scan(&rkt.ID, &rkt.Name, &rkt.Type)
	if err != nil {
		log.Print(err.Error())
		return rocket.Rocket{}, err
	}

	return rkt, nil
}

func (s Store) InsertRocket(rkt rocket.Rocket) (rocket.Rocket, error) {
	_, err := s.db.NamedQuery(
		"INSERT INTO rockets (id, name, type) VALUES (:id, :name, :type)",
		rkt,
	)
	if err != nil {
		return rocket.Rocket{}, errors.New("failed to insert into the database")
	}
	return rocket.Rocket{
		ID:   rkt.ID,
		Type: rkt.Type,
		Name: rkt.Name,
	}, nil
}

func (s Store) DeleteRocket(id string) error {
	// uid, err := uuid.FromString(id)
	// if err != nil {
	// 	return err
	// }

	_, err := s.db.Exec("DELETE FROM rockets where id = $1", id)

	if err != nil {
		return err
	}
	return nil
}
