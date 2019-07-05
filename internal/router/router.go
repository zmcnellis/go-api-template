package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zmcnellis/go-api-template/internal/controllers"
	"github.com/zmcnellis/go-api-template/internal/services"
)

type routeHandler func(w http.ResponseWriter, r *http.Request)
type controller func(svc *services.ResourceService, w http.ResponseWriter, r *http.Request)

// NewRouter ...
func NewRouter(resourceSvc *services.ResourceService) *mux.Router {
	router := mux.NewRouter()
	setupRoutes(router, resourceSvc)
	return router
}

func withService(f controller, s *services.ResourceService) routeHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		f(s, w, r)
	}
}

func setupRoutes(router *mux.Router, resourceSvc *services.ResourceService) {
	router.HandleFunc("/resources", withService(controllers.GetResources, resourceSvc)).Methods("GET")
	router.HandleFunc("/resources/{id}", withService(controllers.GetResource, resourceSvc)).Methods("GET")
	router.HandleFunc("/resources", withService(controllers.CreateResource, resourceSvc)).Methods("POST")
	router.HandleFunc("/resources/{id}", withService(controllers.DeleteResource, resourceSvc)).Methods("DELETE")
}
