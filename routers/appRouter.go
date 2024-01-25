package router

import (
//	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	apiConfig "github.com/kaavinAK/goTaskitServer/config"
	controller "github.com/kaavinAK/goTaskitServer/controllers"
	"github.com/kaavinAK/goTaskitServer/middlewares"
)


func CreateAppRouter(apiConfig apiConfig.ApiConfig) *chi.Mux{
router:=chi.NewRouter()
router.Get("/home",func(w http.ResponseWriter, r *http.Request) {
middlewares.AuthMiddleware(w,r,apiConfig,controller.HomeController)
})
router.Post("/post",func(w http.ResponseWriter, r *http.Request) {
	middlewares.AuthMiddleware(w,r,apiConfig,controller.CreatePostController)
})
router.Get("/post/{id}",func(w http.ResponseWriter, r *http.Request) {
	middlewares.AuthMiddleware(w,r,apiConfig,controller.GetPostByIdController)
})
return router
}