package controller

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	apiConfig "github.com/kaavinAK/goTaskitServer/config"
	"github.com/kaavinAK/goTaskitServer/internal/database"
	"github.com/kaavinAK/goTaskitServer/models"
	"github.com/kaavinAK/goTaskitServer/utilities"
)

func LoginController(w http.ResponseWriter,r *http.Request,apiConfig apiConfig.ApiConfig){
   userData:=models.UserModel{

   }
   utilities.ExtractUserFromRequest(r,&userData)	
   loggedInUser,err:=apiConfig.DB.GetExistingUser(r.Context(),database.GetExistingUserParams{
	Name: userData.Name,
	Password: userData.Password,
   })
   if err !=nil{
   log.Print("login error ",err)
	utilities.ResponseWriter(w,nil,400)
	return 
   }
   createdUserSession,sessionErr:=apiConfig.DB.CreateSession(r.Context(),database.CreateSessionParams{
	ID: uuid.New(),
	Sessionid: uuid.NewString(),
	Userid: loggedInUser.ID,
   })
   if sessionErr!=nil{
	log.Print("session err ",sessionErr)
	utilities.ResponseWriter(w,nil,400)
	
	return 
   }

   userCookie:=http.Cookie{
	Name: "token",
	Value: createdUserSession.Sessionid,
	Path: "/",
  }
  http.SetCookie(w,&userCookie)
  utilities.ResponseWriter(w,loggedInUser,200)
  
}