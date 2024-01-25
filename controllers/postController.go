package controller

import (
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
	apiConfig "github.com/kaavinAK/goTaskitServer/config"
	"github.com/kaavinAK/goTaskitServer/internal/database"
	"github.com/kaavinAK/goTaskitServer/models"
	"github.com/kaavinAK/goTaskitServer/utilities"
)

func CreatePostController(w http.ResponseWriter,r *http.Request,user database.User,apiConfig apiConfig.ApiConfig){
	postData:=models.PostModel{

	}
	err:=utilities.ExtractPostFromRequest(r,&postData)
	if err !=nil{
		log.Print("error occured in extracting post data in request body")
		utilities.ResponseWriter(w,nil,400)
		return
	}
	postInfo,err:=apiConfig.DB.CreatePost(r.Context(),database.CreatePostParams{
		ID: uuid.New(),
		Title: postData.Title,
		Description: postData.Description,
		Authorid: user.ID,
	})
	if err!=nil{
		log.Print("error happened in post data creation")
		utilities.ResponseWriter(w,nil,400)
		return 
	}
	utilities.ResponseWriter(w,postInfo,200)
 
}

func GetPostByIdController(w http.ResponseWriter,r *http.Request,user database.User,apiConfig apiConfig.ApiConfig){
	urlSplitter:=strings.Split(r.URL.Path, "/")
	postId:=urlSplitter[4]
	postIdUUID,err:=uuid.Parse(postId)
	if  err!=nil{
		log.Print("not a valid postID in Url")
		utilities.ResponseWriter(w,nil,400)
		return 
	}
	log.Print("info about postId url params ",postId)
	post,err:=apiConfig.DB.GetPostById(r.Context(),postIdUUID)
	if err!=nil{
		log.Print("error happened during post query")
		utilities.ResponseWriter(w,nil,400)
	}

	utilities.ResponseWriter(w,post,200)
}