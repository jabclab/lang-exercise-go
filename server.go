package main

import (
	"log"

	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", ":6379")

	if err != nil {
		log.Fatal(err)
	}

	// Defer comes after the error handling as if it
	// didn't work there will be no connection to
	// close.
	defer c.Close()
}
