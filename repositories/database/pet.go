package repositories

import "github.com/jerricotandelacruz/golang-clean-architecture/entities"

type PetRepository interface {
	Create(*entities.Pet) error
	Update(int, *entities.Pet) error
	Select() ([]entities.Pet, error)
	SelectById(int) (*entities.Pet, error)
}
