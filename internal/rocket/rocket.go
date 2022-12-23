//go:generate mockgen -destination=rocket_mocks_test.go -package=rocket github.com/Raphael-Jin/Go-gRPC-Services/internal/rocket Store

package rocket

import (
	"context"
)

// Rocket should contain the def
// of our rocket
type Rocket struct {
	ID      string
	Name    string
	Type    string
	Flights int
}

// interface database should implement
type Store interface {
	GetRocketByID(id string) (Rocket, error)
	InsertRocket(rkt Rocket) (Rocket, error)
	DeleteRocket(id string) error
}

// update the rocket inventory
type Service struct {
	Store Store
}

// return a new instance of our rocket service
func New(store Store) Service {
	return Service{
		Store: store,
	}
}

// Get Rocket by id
func (s Service) GetRocketByID(ctx context.Context, id string) (Rocket, error) {
	rkt, err := s.Store.GetRocketByID(id)
	if err != nil {
		return Rocket{}, err
	}
	return rkt, nil
}

func (s Service) InsertRocket(ctx context.Context, rkt Rocket) (Rocket, error) {
	rkt, err := s.Store.InsertRocket(rkt)
	if err != nil {
		return Rocket{}, err
	}
	return rkt, nil
}

func (s Service) DeleteRocket(id string) error {
	err := s.Store.DeleteRocket(id)
	if err != nil {
		return err
	}
	return nil
}
