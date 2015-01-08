package main

import (
	"bytes"
	//"io/ioutil"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	testingRedisPort = 7777
	testingHttpPort  = 9999
)

var (
	inUrl  = fmt.Sprintf("http://localhost:%d/in/", testingHttpPort)
	outUrl = fmt.Sprintf("http://localhost:%d/messages/", testingHttpPort)
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

	// Not very nice, but give redis-server a chance to
	// start.
	time.Sleep(250 * time.Millisecond)
}

func TestMain(m *testing.M) {
	startRedis()

	// Start the HTTP server for our integration tests.
	os.Setenv("REDIS_PORT", strconv.Itoa(testingRedisPort))
	os.Setenv("HTTP_PORT", strconv.Itoa(testingHttpPort))
	go func() {
		main()
	}()

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

	resp, err := http.Post(
		inUrl, "text/plain", bytes.NewBufferString("my test message"),
	)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	//	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

//	if resp.Status != string(http.StatusCreated) {
//		t.Errorf("Incorrect status code")
//	}

	//if string(body[:]) != "{\"messageId\":1}" {
	//		t.Errorf("Message was not as expected")
	//	}
}

func BenchmarkMessageCreation(*testing.B) {

}
