package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zmcnellis/go-api-template/internal/models"
	"github.com/zmcnellis/go-api-template/internal/services"
)

// GetResources ...
func GetResources(svc *services.ResourceService, w http.ResponseWriter, r *http.Request) {
	resources := svc.GetResources()
	json.NewEncoder(w).Encode(&resources)
}

// GetResource ...
func GetResource(svc *services.ResourceService, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	resource := svc.GetResource(id)
	json.NewEncoder(w).Encode(&resource)
}

// CreateResource ...
func CreateResource(svc *services.ResourceService, w http.ResponseWriter, r *http.Request) {
	var resource models.Resource
	json.NewDecoder(r.Body).Decode(&resource)
	svc.CreateResource(&resource)
	json.NewEncoder(w).Encode(&resource)
}

// DeleteResource ...
func DeleteResource(svc *services.ResourceService, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	svc.DeleteResource(id)

	resources := svc.GetResources()
	json.NewEncoder(w).Encode(&resources)
}
