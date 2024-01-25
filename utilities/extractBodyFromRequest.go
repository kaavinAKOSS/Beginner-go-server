package utilities

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/kaavinAK/goTaskitServer/models"
)


func ExtractUserFromRequest(r* http.Request,dataModel* models.UserModel) error {
	decoder:=json.NewDecoder(r.Body)
	
	err:=decoder.Decode(&dataModel)
	log.Print("body ",r.Body,dataModel)
	if err!=nil{
		return errors.New("cant extract body from request")
	}
	return nil;

	
   
}
func ExtractPostFromRequest(r* http.Request,dataModel* models.PostModel) error{
	decoder:=json.NewDecoder(r.Body)
	err:=decoder.Decode(&dataModel)
	//log.Print("")
	if err !=nil{
		return errors.New("cant extract post from body")
	}
	return nil
}
