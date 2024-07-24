package controller

import "net/http"

type PetController interface {
	NewPet(w http.ResponseWriter, r *http.Request)
	GetPets(w http.ResponseWriter, r *http.Request)
	GetPetById(w http.ResponseWriter, r *http.Request)
	EditPetById(w http.ResponseWriter, r *http.Request)
}

type petController struct{}

func NewPetController() PetController {

}
