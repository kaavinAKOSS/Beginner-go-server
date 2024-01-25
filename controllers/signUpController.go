package controller

import (
	"log"
	"net/http"
	
	"time"

	"github.com/google/uuid"
	apiConfig "github.com/kaavinAK/goTaskitServer/config"
	"github.com/kaavinAK/goTaskitServer/internal/database"
	"github.com/kaavinAK/goTaskitServer/models"
	"github.com/kaavinAK/goTaskitServer/utilities"
)


func SignUpController(w http.ResponseWriter,r *http.Request,apiConfig apiConfig.ApiConfig)  {
	log.Print(r.Body)
	userData:=models.UserModel{

	}
	err:=utilities.ExtractUserFromRequest(r,&userData)
   if err!=nil{
	log.Print(err)
   }
   existingUser,_:=apiConfig.DB.GetExistingUser(r.Context(),database.GetExistingUserParams{
	Name: userData.Name,
	Password: userData.Password,
   })
  
   if existingUser.ID!=uuid.Nil{
	log.Print("user already exists")
     utilities.ResponseWriter(w,nil,400)
	 return 
   }
   
   log.Print("userData ",userData)
	dbUser,dberr:=apiConfig.DB.SignUpUser(r.Context(),database.SignUpUserParams{
		ID: uuid.New(),
		Name: userData.Name,
		Password: userData.Password,
		CreatedAt: time.Now(),
	})
	if dberr!=nil{
		utilities.ResponseWriter(w,nil,400)
		return 
	}
	log.Print("dbUser ",dbUser)
	responseUser:=struct{
		Name string `json:"name"`;
		Password string `json:"password"`
	}{
		Name: dbUser.Name,
		Password: dbUser.Password,
	}

	
       utilities.ResponseWriter(w,responseUser,200)
	
}