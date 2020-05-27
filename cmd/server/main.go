package main

import (
	contracts "iboks-contracts"
	"log"
	"os"
)

func main() {
	var path string
	var port string

	path, ok := os.LookupEnv("CONTRACTS_PATH")
	if !ok {
		log.Fatal("CONTRACTS_PATH not found in environment")
	}
	port, ok = os.LookupEnv("CONTRACTS_PORT")
	if !ok {
		port = "8083"
	}
	s := contracts.NewServer(port, path)

	log.Fatal(s.S.ListenAndServe())
}
