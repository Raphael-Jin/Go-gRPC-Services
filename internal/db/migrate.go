package db

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func (s *Store) Migrate() error {
	driver, err := postgres.WithInstance(s.db.DB, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"postgres",
		driver,
	)
	if err != nil {
		return err
	}

	m.Steps(2)

	s.db.Exec(
		"CREATE TABLE IF NOT EXISTS rockets(id varchar (10) NOT NULL PRIMARY KEY, type varchar (50), name varchar (50));",
	)

	return nil

	// driver, err := postgres.WithInstance(s.db.DB, &postgres.Config{})
	// if err != nil {
	// 	return err
	// }

	// m, err := migrate.NewWithDatabaseInstance(
	// 	"file:///migrations",
	// 	"postgres",
	// 	driver,
	// )
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return err
	// }

	// if err := m.Up(); err != nil {
	// 	if err.Error() == "no change" {
	// 		log.Println("no change made by migrations")
	// 	} else {
	// 		return err
	// 	}
	// }
	return nil
}
