package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	apiConfig "github.com/kaavinAK/goTaskitServer/config"
	controller "github.com/kaavinAK/goTaskitServer/controllers"
)

func CreateAuthRouter(apiconfig apiConfig.ApiConfig) *chi.Mux{
	router:=chi.NewRouter()
	router.Post("/register",func(w http.ResponseWriter, r *http.Request) {
		controller.SignUpController(w,r,apiconfig)
	})
	router.Post("/login",func(w http.ResponseWriter, r *http.Request) {
		controller.LoginController(w,r,apiconfig)
	})
	return router
}


