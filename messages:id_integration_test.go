package main

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/franela/goreq"
	"github.com/stretchr/testify/assert"
)

func TestReturnsMessageIfRequested(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping in short mode")
	}

	msg := "testing message"

	inResponse, err := goreq.Request{
		Method: "POST",
		Uri:    inUrl,
		Body:   msg,
	}.Do()

	if err != nil {
		t.Fatal(err)
	}

	var message MessageCreationResponse

	inResponse.Body.FromJsonTo(&message)

	outResponse, outErr := goreq.Request{
		Uri: outUrl + strconv.FormatFloat(message.MessageId, 'f', 0, 64) + "/",
	}.Do()

	if outErr != nil {
		t.Fatal(outErr)
	}

	response, strErr := outResponse.Body.ToString()
	if strErr != nil {
		t.Fatal(strErr)
	}

	assert.Equal(t, msg, response)
}

func TestShouldReturn400IfMessageDoesNotExist(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping in short mode")
	}

	invalidId := "12345"

	invalidResponse, err := goreq.Request{
		Uri: outUrl + invalidId + "/",
	}.Do()

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusBadRequest, invalidResponse.StatusCode)
}

func TestShouldReturnErrorMessageIfMessageDoesNotExist(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping in short mode")
	}

	invalidId := "12345"

	invalidResponse, err := goreq.Request{
		Uri: outUrl + invalidId + "/",
	}.Do()

	if err != nil {
		t.Fatal(err)
	}

	body, strErr := invalidResponse.Body.ToString()
	if strErr != nil {
		t.Fatal(strErr)
	}

	assert.Equal(t, "message with this ID does not exist", body)
}
