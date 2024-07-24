package controller

import "net/http"

type AdopterController interface {
	RegisterAdopter(w http.ResponseWriter, r *http.Request)
	GetAdopters(w http.ResponseWriter, r *http.Request)
	GetAdopterById(w http.ResponseWriter, r *http.Request)
	EditAdopterById(w http.ResponseWriter, r *http.Request)
}

type adopterController struct{}

// EditAdopterById implements AdopterController.
func (a *adopterController) EditAdopterById(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// GetAdopterById implements AdopterController.
func (a *adopterController) GetAdopterById(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// GetAdopters implements AdopterController.
func (a *adopterController) GetAdopters(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// RegisterAdopter implements AdopterController.
func (a *adopterController) RegisterAdopter(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func NewAdopterController() AdopterController {
	return &adopterController{}
}
