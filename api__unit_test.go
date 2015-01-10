package main

import (
	"testing"

	//"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

/* AddMessage */

type AddMessageSuite struct {
	suite.Suite
}

func (suite *AddMessageSuite) TestExample() {
}

func TestAddMessage(t *testing.T) {
	suite.Run(t, new(AddMessageSuite))
}

/* GetMessage */

type GetMessageSuite struct {
	suite.Suite
}

func (suite *GetMessageSuite) TestExample2() {
}

func TestGetMessage(t *testing.T) {
	suite.Run(t, new(GetMessageSuite))
}
