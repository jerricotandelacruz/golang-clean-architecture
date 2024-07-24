package entities

type PetType int

const (
	Dog PetType = iota
	Cat
	Rabbit
	Hamster
)

type Pet struct {
	ID          int
	Name        string
	Breed       string
	Age         int
	Type        PetType
	IsAvailable bool
}
