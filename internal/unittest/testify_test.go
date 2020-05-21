package unittest

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"testing"
)

func TestTestifyCase1(t *testing.T) {
	assert.Equal(t, 1, 1, "they should equal")
}

type MockObj struct {
	mock.Mock
}

func (m *MockObj) DoSomething(number int) (bool, error) {
	args := m.Called(number)
	return args.Bool(0), args.Error(1)
}

func TestSomething(t *testing.T) {
	obj := &MockObj{}
	obj.On("DoSomething", 123).Return(true, nil)
	obj.On("DoSomething", 456).Return(false, nil)

	{
		ok, err := obj.DoSomething(123)
		assert.True(t, ok, "it is true")
		assert.Nil(t, err, "it is nil")
	}
	{
		ok, err := obj.DoSomething(456)
		assert.False(t, ok, "it is false")
		assert.Nil(t, err, "it is nil")
	}
}

type ExampleTestSuite struct {
	suite.Suite
	VariableThatShouldStartAtFive int
}

func (suite *ExampleTestSuite) SetupTest() {
	suite.VariableThatShouldStartAtFive = 5
}

func (suite *ExampleTestSuite) TestExample() {
	assert.Equal(suite.T(), 5, suite.VariableThatShouldStartAtFive)
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, &ExampleTestSuite{})
}
