package main

import (
	"log"
	"sessionsdemo/initialize"
)

func main() {

	initialize.InitRedisClient()
	r := initialize.Routers()

	log.Fatal(r.Run(":8008"))
}
