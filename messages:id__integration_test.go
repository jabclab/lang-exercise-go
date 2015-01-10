package main

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/franela/goreq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MessageCreationResponse struct {
	MessageId float64 `json:"messageId"`
}

type MessageIdSuite struct {
	suite.Suite
}

func (suite *MessageIdSuite) TestReturnsMessageIfRequested() {
	msg := "testing message"

	inResponse, err := goreq.Request{
		Method: "POST",
		Uri:    inUrl,
		Body:   msg,
	}.Do()

	if err != nil {
		suite.T().Fatal(err)
	}

	var message MessageCreationResponse

	inResponse.Body.FromJsonTo(&message)

	outResponse, outErr := goreq.Request{
		Uri: outUrl + strconv.FormatFloat(message.MessageId, 'f', 0, 64) + "/",
	}.Do()

	if outErr != nil {
		suite.T().Fatal(outErr)
	}

	response, strErr := outResponse.Body.ToString()
	if strErr != nil {
		suite.T().Fatal(strErr)
	}

	assert.Equal(suite.T(), msg, response)
}

func (suite *MessageIdSuite) TestShouldReturn400IfMessageDoesNotExist() {
	invalidId := "12345"

	invalidResponse, err := goreq.Request{
		Uri: outUrl + invalidId + "/",
	}.Do()

	if err != nil {
		suite.T().Fatal(err)
	}

	assert.Equal(suite.T(), http.StatusBadRequest, invalidResponse.StatusCode)
}

func (suite *MessageIdSuite) TestShouldReturnErrorMessageIfMessageDoesNotExist() {
	invalidId := "12345"

	invalidResponse, err := goreq.Request{
		Uri: outUrl + invalidId + "/",
	}.Do()

	if err != nil {
		suite.T().Fatal(err)
	}

	body, strErr := invalidResponse.Body.ToString()
	if strErr != nil {
		suite.T().Fatal(strErr)
	}

	assert.Equal(suite.T(), "message with this ID does not exist", body)
}

func TestRouteMessageId(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping in short mode")
	}

	suite.Run(t, new(MessageIdSuite))
}
