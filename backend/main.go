package main

import (
	"github.com/forest747/go-mvc-project/backend/src/rest"
	"log"
)

func main() {
	log.Println("Main log....")
	log.Fatal(rest.RunAPI("127.0.0.1:8080"))
}
