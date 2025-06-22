package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John Doe", "j@j.com", "1234578")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "j@j.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("John Doe", "j@j.com", "1234578")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.True(t, user.ValidatePassword("1234578"))
}

func TestUser_PasswordInvalid(t *testing.T) {
	user, err := NewUser("John Doe", "j@j.com", "1234578")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.False(t, user.ValidatePassword("1234538"))
}
