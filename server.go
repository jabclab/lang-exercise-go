package main

import (
	"log"
	"os"

	"github.com/codegangsta/martini"
	"github.com/garyburd/redigo/redis"
	"github.com/martini-contrib/render"
)

var m *martini.ClassicMartini

func redisPort() string {
	envVar := os.Getenv("REDIS_PORT")
	if envVar != "" {
		return envVar
	}

	return "6379"
}

func init() {
	// Classic martini includes routing, logging,
	// panic recovery and static file serving.
	m = martini.Classic()

	// Additional middleware.
	m.Use(render.Renderer())

	// Routes
	m.Post("/in/", AddMessage)
	m.Get("/messages/:id/", GetMessage)
}

func main() {
	store, err := redis.Dial("tcp", ":"+redisPort())
	if err != nil {
		log.Fatal(err)
	}

	// Defer comes after the error handling as if it
	// didn't work there will be no connection to
	// close.
	defer store.Close()

	// Map the Redis instance
	m.Map(store)

	// Some examples seem to wrap this in an anonymous
	// go routine but not sure why at this point!
	m.RunOnAddr(":8888")
}
