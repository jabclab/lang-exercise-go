package main

import (
	"log"
	"os"

	"github.com/garyburd/redigo/redis"
)

func redisPort() string {
	envVar := os.Getenv("REDIS_PORT")
	if envVar != "" {
		return envVar
	}

	return "6379"
}

func main() {
	c, err := redis.Dial("tcp", ":" + redisPort())

	if err != nil {
		log.Fatal(err)
	}

	// Defer comes after the error handling as if it
	// didn't work there will be no connection to
	// close.
	defer c.Close()
}
