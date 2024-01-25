package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	apiConfig "github.com/kaavinAK/goTaskitServer/config"
	"github.com/kaavinAK/goTaskitServer/internal/database"
	router "github.com/kaavinAK/goTaskitServer/routers"
	_ "github.com/lib/pq"
)

func main(){
	goDotError:=godotenv.Load()
	if goDotError!=nil{
		log.Fatal("failed to load env variables")
	}
	port:=os.Getenv("PORT")
	db_url:=os.Getenv("DB_URL")
	dbConn,dberr:=sql.Open("postgres",db_url)
    if dberr!=nil{
		log.Fatal("db refused to connect...")
		return 
	}
	apiConfig:=apiConfig.ApiConfig{
		DB: database.New(dbConn),
	}
	ServerRouter:=chi.NewRouter()
	ServerRouter.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	  }))
	authRouter:=router.CreateAuthRouter(apiConfig)
	appRouter:=router.CreateAppRouter(apiConfig)
	ServerRouter.Mount("/api/auth",authRouter);
	ServerRouter.Mount("/api/app",appRouter)
	server:=&http.Server{
		Handler: ServerRouter,
		Addr: ":"+port,
	}
	log.Print("server starting ....")
	serverError:=server.ListenAndServe()
	if serverError!=nil{
		log.Fatal("server failed to start...")
	}
}