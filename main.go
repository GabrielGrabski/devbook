package main

import (
	"api-go/src/config"
	"api-go/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.GetEnvVars()
	port := fmt.Sprintf(":%d", config.Port)
	fmt.Printf("Listening on port %s", port)
	r := router.GetRouter()
	log.Fatal(http.ListenAndServe(port, r))
}
