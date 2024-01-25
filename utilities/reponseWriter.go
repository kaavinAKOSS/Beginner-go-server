package utilities

import (
	"encoding/json"
	"log"
	"net/http"
)

func ResponseWriter(w http.ResponseWriter,payload interface{},statusCode int){
	if statusCode==400 {
		w.WriteHeader(400)
		return 
	}
	log.Print("payload ",payload)
	binaryData,err:=json.Marshal(payload)
	if err !=nil{
		log.Print("json marshal failed")
		w.WriteHeader(400)
		return ;
	}
	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(200)
	w.Write(binaryData)

}