package main

import (
	"log"
	"os"

	"github.com/codegangsta/martini"
	"github.com/garyburd/redigo/redis"
	"github.com/martini-contrib/render"
)

var m *martini.ClassicMartini

func port(envVar, defaultValue string) string {
	envValue := os.Getenv(envVar)
	if envValue != "" {
		return envValue
	}

	return defaultValue
}

func redisPort() string {
	return port("REDIS_PORT", "6379")
}

func httpPort() string {
	return port("HTTP_PORT", "8888")
}

func init() {
	// Classic martini includes routing, logging,
	// panic recovery and static file serving.
	m = martini.Classic()

	// Additional middleware.
	m.Use(render.Renderer())

	// Routes.
	m.Post("/in/", AddMessage)
	m.Get("/messages/:id/", GetMessage)
}

func main() {
	store, err := redis.Dial("tcp", ":" + redisPort())
	if err != nil {
		log.Fatal(err)
	}

	// Defer comes after the error handling as if it
	// didn't work there will be no connection to
	// close.
	defer store.Close()

	// Map the Redis instance.
	m.Map(store)

	// Start the HTTP server.
	m.RunOnAddr(":" + httpPort())
}
