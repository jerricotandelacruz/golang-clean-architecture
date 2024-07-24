package repositories

import "github.com/jerricotandelacruz/golang-clean-architecture/entities"

type AdoptionRepository interface {
	Create(*entities.Adoption) error
	UpdateStatus(int, *entities.Adoption) error
}
