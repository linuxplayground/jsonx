package main

import (
	//"fmt"
	"log"
	"net/http"
	
	"latham.nz/featly.common"
)

var err error

func main() {
	//Router
	router := common.NewRouter(routes)
	log.Fatal(http.ListenAndServe(":8080", router))
}
