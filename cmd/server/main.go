package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/pondparinya/go-template/config"
)

func main() {
	log.Println("starting server")
	state := flag.String("state", "local", "product state: local/sit/uat/nft/prd")
	flag.Parse()

	fileName := fmt.Sprintf("app-%s", *state)
	if err := config.Load("./config", fileName, "yaml"); err != nil {
		log.Fatalln(err)
	}

	if err := config.Load("./config", "error", "yaml"); err != nil {
		log.Fatalln(err)
	}

	log.Println(config.APP)
}
