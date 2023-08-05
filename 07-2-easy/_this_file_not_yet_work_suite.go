package main

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

// Define a test suite struct that embeds suite.Suite
type MySuite struct {
	suite.Suite
	myVar int
}

// Define a SetupTest method that will be called before each test
func (s *MySuite) SetupTest() {
	s.myVar = 42
}

// Define a TearDownTest method that will be called after each test
func (s *MySuite) TearDownTest() {
	// cleanup code goes here
}

// Define individual test functions that will be run by the suite
func (s *MySuite) TestSomething() {
	// use s.Assert() or s.Require() to make assertions
	s.Assert().Equal(42, s.myVar)
}

// Define another test function
func (s *MySuite) TestSomethingElse() {
	s.myVar = 0
	s.Require().NotEqual(42, s.myVar)
}

// In your Go test file
func TestMySuite(t *testing.T) {
	suite.Run(t, new(MySuite))
}
