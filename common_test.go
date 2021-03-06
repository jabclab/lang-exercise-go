package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"testing"
	"time"
)

const (
	localTestingRedisPort = 7777
	localTestingHttpPort  = 9999
)

var inUrl, outUrl string

func stopRedis() {
	// Make sure Redis server is running.
	killCmd := exec.Command(
		"killall",
		fmt.Sprintf("redis-server *:%d", localTestingRedisPort),
	)
	killCmd.Run()
}

func startRedis() {
	stopRedis()

	redisPath, err := exec.LookPath("redis-server")
	if err != nil {
		log.Fatal(err)
	}

	startCmd := exec.Command(
		redisPath,
		fmt.Sprintf("--port %d", localTestingRedisPort),
	)
	startErr := startCmd.Start()

	if startErr != nil {
		log.Fatal(startErr)
	}

	// Not very nice, but give redis-server a chance to
	// start.
	time.Sleep(250 * time.Millisecond)
}

func TestMain(m *testing.M) {
	redisRequired := os.Getenv("TRAVIS") != "true"

	// If we're running tests on Travis we do not need to
	// run Redis as this will be run for us (using stock
	// configuration).
	if redisRequired {
		startRedis()

		os.Setenv("REDIS_PORT", strconv.Itoa(localTestingRedisPort))
		os.Setenv("HTTP_PORT", strconv.Itoa(localTestingHttpPort))
	}

	// Use the implementation code which obtains HTTP port.
	http := httpPort()

	inUrl = fmt.Sprintf("http://localhost:%s/in/", http)
	outUrl = fmt.Sprintf("http://localhost:%s/messages/", http)

	// Start the HTTP server for our integration tests. Note
	// that this must be started in a go routine so that it
	// does not block execution of the tests. Another option
	// for this would be to use net/http/httptest with NewServer.
	// However I like this way as with using this we're driving
	// our web server without any intervention.
	go main()

	// Run the tests.
	result := m.Run()

	if redisRequired {
		stopRedis()
	}

	os.Exit(result)
}
