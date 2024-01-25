package controller

import (
	"log"
	"net/http"

	apiConfig "github.com/kaavinAK/goTaskitServer/config"
	"github.com/kaavinAK/goTaskitServer/internal/database"
	"github.com/kaavinAK/goTaskitServer/utilities"
)

type homeControllerReturnType struct{
	Posts []database.Post `json:"posts"` ;
	User database.User `json:"user"` ;

}
func HomeController(w http.ResponseWriter,r *http.Request,
	user database.User,
	apiConfig apiConfig.ApiConfig){
  postsData,err:=apiConfig.DB.GetPostByUser(r.Context(),user.ID)
  if err !=nil{
	log.Print("cannot  get user posts")
	utilities.ResponseWriter(w,nil,400)
	return
  }
  returnData:=homeControllerReturnType{
	Posts: postsData,
	User: user,
  }
  utilities.ResponseWriter(w,returnData,200)
}