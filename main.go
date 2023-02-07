package main

import (
	constants "constants_project"
	"fmt"
	"router_functions"

	"github.com/joho/godotenv"
)

func main() {
	defer fmt.Printf("Server closed.\n")

	godotenv.Load()
	godotenv.Load(".env")

	fmt.Printf("Server open on %v.\n", constants.HOST_PORT)
	router_functions.StartServer(constants.HOST_PORT)
}
