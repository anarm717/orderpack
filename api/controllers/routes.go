package controllers

import (
	_ "GoOrderPackProject/docs"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

func (s *Server) initializeRoutes() {

	var api1 = s.Router.PathPrefix("/api/v1").Subrouter()

	// Home Route
	api1.HandleFunc("/", s.Home).Methods("GET")

	api1.HandleFunc("/pack-sizes", s.GetPackSizes).Methods("GET")
	api1.HandleFunc("/pack-sizes", s.UpdatePackSizes).Methods("POST")
	api1.HandleFunc("/calculate-packs", s.CalculatePacksByQuantity).Methods("POST")

	// Swagger
	s.Router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)
}
