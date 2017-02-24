package main


import (
	"net/http"
	"log"
	"fmt"
	"os"
)

func main() {

	Server := NewServer()

	fmt.Println("Hosted on port",GetPort())
	log.Fatal(http.ListenAndServe(GetPort(),Server))
}


func GetPort() string {
        var port = os.Getenv("PORT")
        // Set a default port if there is nothing in the environment
        if port == "" {
                port = "3000"
                fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
        }
        return ":" + port
}