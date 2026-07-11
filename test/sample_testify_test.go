package test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestFuncAssertion(t *testing.T) {
	// assert equality
	assert.Equal(t, 123, 123, "they should be equal")

	// assert inequality
	assert.NotEqual(t, 123, 456, "they should not be equal")

	// assert for nil (good for errors)
	assert.Nil(t, nil, "they should be nil")

	// assert for not nil (good when expect something)
	assert.NotNil(t, 1, "they should not be nil")

	// assert for error
	err := errors.New("something went wrong")
	assert.Error(t, err, "they should be error")

	// assert for not error
	assert.NoError(t, nil, "they should not be error")

	// assert boolean
	assert.True(t, true)
	assert.False(t, false)

	// assert empty
	users := [1]string{}
	assert.Empty(t, users, "they should be empty")

	// assert not empty
	users[0] = "Aaron"
	assert.NotEmpty(t, users, "they should not be empty")

	// assert contains
	assert.Contains(t, "Aaron", "ar")
	assert.NotContains(t, "Aaron", "zz")

	// assert length
	assert.Len(t, users, 1)

	// assert type
	assert.IsType(t, [1]string{}, users)

	var number = 0
	var text = ""
	var number1 = 1

	// assert zero value
	assert.Zero(t, number)
	assert.Zero(t, text)
	assert.NotZero(t, number1)
}

func TestRequireFunction(t *testing.T) {
	require.NoError(t, nil)
	require.Error(t, errors.New("empty"))
	require.Nil(t, nil)
	require.NotNil(t, 1)
	require.Equal(t, true, true)
	require.Len(t, []string{}, 0)
}

// UserFetcher represents a service that fetches user data.
type UserFetcher interface {
	FetchUser(id int) (string, error)
}

// GetUserGreeting returns a greeting message using the fetcher.
func GetUserGreeting(fetcher UserFetcher, id int) string {
	name, err := fetcher.FetchUser(id)
	if err != nil {
		return "Hello, Guest!"
	}
	return "Hello, " + name + "!"
}

// MockUserFetcher is a mock implementation of UserFetcher.
type MockUserFetcher struct {
	mock.Mock
}

func (m *MockUserFetcher) FetchUser(id int) (string, error) {
	args := m.Called(id)
	return args.String(0), args.Error(1)
}

func TestGetUserGreeting(t *testing.T) {
	// mock for success case — user found
	mockFetcher := new(MockUserFetcher)
	mockFetcher.On("FetchUser", 1).Return("Aaron", nil)

	result := GetUserGreeting(mockFetcher, 1)
	assert.Equal(t, "Hello, Aaron!", result)
	mockFetcher.AssertExpectations(t)

	// mock for error case — user not found
	mockFetcher2 := new(MockUserFetcher)
	mockFetcher2.On("FetchUser", 999).Return("", errors.New("not found"))

	result2 := GetUserGreeting(mockFetcher2, 999)
	assert.Equal(t, "Hello, Guest!", result2)
	mockFetcher2.AssertExpectations(t)
}
