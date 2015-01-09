package main

import (
	"log"
	"testing"

	//"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// Define test suite sruct.
type AddMessageSuite struct {
	suite.Suite
}

func (suite *AddMessageSuite) TestExample() {
	log.Printf("hello")
}

// Run the test suite.
func TestAddMessage(t *testing.T) {
	suite.Run(t, new(AddMessageSuite))
}

type GetMessageSuite struct {
	suite.Suite
}

func (suite *GetMessageSuite) TestExample2() {
	log.Printf("hello 2")
}

func TestGetMessage(t *testing.T) {
	suite.Run(t, new(GetMessageSuite))
}
