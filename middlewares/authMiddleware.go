package middlewares

import (
	"log"
	"net/http"

	apiConfig "github.com/kaavinAK/goTaskitServer/config"
	"github.com/kaavinAK/goTaskitServer/internal/database"
	"github.com/kaavinAK/goTaskitServer/utilities"
)


type authHandlerfunc func(w http.ResponseWriter,r* http.Request,user database.User,apiconfig apiConfig.ApiConfig)

func AuthMiddleware(w http.ResponseWriter,r* http.Request,apiConfig apiConfig.ApiConfig,handlerFunc authHandlerfunc){
 token:=utilities.ExtractTokenFromCookie(r)
 userSession,sessionErr:=apiConfig.DB.GetUserSession(r.Context(),token)
 if sessionErr!=nil{
	log.Print("session error in auth middleware")
	utilities.ResponseWriter(w,nil,400)
	return 
 }
 user,userErr:=apiConfig.DB.GetUserByID(r.Context(),userSession.Userid)
 if userErr !=nil{
	log.Print("user error in auth middleware")
	utilities.ResponseWriter(w,nil,400)
	return 
 }

 handlerFunc(w,r,user,apiConfig)
}