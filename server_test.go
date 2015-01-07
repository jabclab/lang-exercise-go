package main

import (
	"bytes"
	//"io/ioutil"
	"fmt"
	"net/http"
	"os"
	"log"
	"os/exec"
	"testing"
)

const (
	inUrl = "http://localhost:8080/in/"
	outUrl = "http://localhost:8080/messages/"
	testingRedisPort = 7777
)

func stopRedis() {
	// Make sure Redis server is running.
	killCmd := exec.Command(
		"killall",
		fmt.Sprintf("redis-server *:%d", testingRedisPort),
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
		fmt.Sprintf("--port %d", testingRedisPort),
	)
	startErr := startCmd.Start()

	if startErr != nil {
		log.Fatal(startErr)
	}
}

func TestMain(m *testing.M) {
	startRedis()

	// Start the HTTP server for our integration tests.
	os.Setenv("REDIS_PORT", "7777")
	main()

	// Run the tests.
	result := m.Run()

	stopRedis()

	os.Exit(result)
}

func TestServer(t *testing.T) {
}

func TestReturns201IfCreatesMessage(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping in short mode")
	}

	msg := "my testing message"

	resp, err := http.Post(inUrl, "text/plain", bytes.NewBufferString(msg))
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	//	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	if resp.Status != string(http.StatusCreated) {
		t.Errorf("Incorrect status code")
	}

	//if string(body[:]) != "{\"messageId\":1}" {
	//		t.Errorf("Message was not as expected")
	//	}
}

func BenchmarkMessageCreation(*testing.B) {

}
