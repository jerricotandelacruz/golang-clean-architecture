package repositories

import "github.com/jerricotandelacruz/golang-clean-architecture/entities"

type AdopterRepository interface {
	Create(*entities.Adopter) error
	Update(int, *entities.Adopter) error
	Select() ([]entities.Adopter, error)
	SelectById(int) (*entities.Adopter, error)
}
