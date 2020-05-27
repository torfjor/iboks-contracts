package main

import (
	contracts "github.com/torfjor/iboks-contracts"
	"log"
	"os"
)

func main() {
	var path string
	var port string

	path, ok := os.LookupEnv("CONTRACTS_PATH")
	if !ok {
		path = "C:\\SpaceMan32\\Contracts"
	}
	port, ok = os.LookupEnv("CONTRACTS_PORT")
	if !ok {
		port = "8083"
	}
	s := contracts.NewServer(port, path)

	log.Fatal(s.S.ListenAndServe())
}
