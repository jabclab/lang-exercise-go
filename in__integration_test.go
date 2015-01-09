package main

import (
	"net/http"
	"testing"

	"github.com/franela/goreq"
	"github.com/stretchr/testify/assert"
)

func TestReturns201IfCreatesMessage(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping in short mode")
	}

	resp, err := goreq.Request{
		Method: "POST",
		Uri:    inUrl,
		Body:   "my test message",
	}.Do()

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

func TestReturnsMessageIdInJsonIfCreated(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping in short mode")
	}

	resp, err := goreq.Request{
		Method: "POST",
		Uri:    inUrl,
		Body:   "my test message",
	}.Do()

	if err != nil {
		t.Fatal(err)
	}

	body, strErr := resp.Body.ToString()
	if strErr != nil {
		t.Fatal(strErr)
	}

	// TODO: rather than hard coding 2 we should reset the
	//       Redis store before each test.
	assert.Equal(t, "{\"messageId\":2}", body)
}
