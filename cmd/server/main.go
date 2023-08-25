package main

import (
	"flag"
	"log"

	"github.com/pondparinya/go-template/config"
)

func main() {
	log.Println("starting server")
	// state := flag.String("state", "local", "product state: local/sit/uat/nft/prd")
	flag.Parse()

	if err := config.Loading("./config", "app-local", "yaml"); err != nil {
		log.Fatalln(err)
	}

	if err := config.Loading("./config", "error", "yaml"); err != nil {
		log.Fatalln(err)
	}

	log.Println(config.APP)
}
