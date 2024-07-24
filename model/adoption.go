package entities

import "time"

type AdoptionStatus int

const (
	Pending AdoptionStatus = iota
	Adopted
	Cancelled
	Returned
)

type Adoption struct {
	ID         int
	CustomerID int
	PetID      int
	Status     AdoptionStatus
	DateTime   time.Time
}
