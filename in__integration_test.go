package main

import (
	"net/http"
	"testing"

	"github.com/franela/goreq"
	"github.com/garyburd/redigo/redis"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type InSuite struct {
	suite.Suite
	RedisClient redis.Conn
}

func (suite *InSuite) SetupSuite() {
	var err error
	suite.RedisClient, err = redis.Dial("tcp", ":"+redisPort())
	if err != nil {
		suite.T().Fatal(err)
	}
}

func (suite *InSuite) SetupTest() {
	// Clear down Redis prior to each test running.
	_, err := suite.RedisClient.Do("FLUSHDB")
	if err != nil {
		suite.T().Fatal(err)
	}
}

func (suite *InSuite) TeardownSuite() {
	suite.RedisClient.Close()
}

func (suite *InSuite) TestReturns201IfCreatesMessage() {
	resp, err := goreq.Request{
		Method: "POST",
		Uri:    inUrl,
		Body:   "my test message",
	}.Do()

	if err != nil {
		suite.T().Fatal(err)
	}

	assert.Equal(suite.T(), http.StatusCreated, resp.StatusCode)
}

func (suite *InSuite) TestReturnsMessageIdInJsonIfCreated() {
	resp, err := goreq.Request{
		Method: "POST",
		Uri:    inUrl,
		Body:   "my test message",
	}.Do()

	if err != nil {
		suite.T().Fatal(err)
	}

	body, strErr := resp.Body.ToString()
	if strErr != nil {
		suite.T().Fatal(strErr)
	}

	assert.Equal(suite.T(), "{\"messageId\":1}", body)
}

func TestRouteIn(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping in short mode")
	}

	suite.Run(t, new(InSuite))
}
