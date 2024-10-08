package main

import (
	"base-convertor/api"
	"log"
)


func main(){
	server := api.NewServer()	

	log.Fatal(server.StartServer())
}