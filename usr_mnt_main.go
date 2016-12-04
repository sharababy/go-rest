package main


import (
	"net/http"
	"log"
	"fmt"
)




func main() {

	Server := NewServer()
	
	

	fmt.Println("Hosted on port :3000")
	log.Fatal(http.ListenAndServe(":3000",Server))
}