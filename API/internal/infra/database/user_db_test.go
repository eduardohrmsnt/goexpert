package database

import (
	"testing"

	"github.com/eduardohrmsnt/go-expert-api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
		return
	}

	db.AutoMigrate(&entity.User{})

	userDb := NewUser(db)

	user, err := entity.NewUser("Usuario", "email@email.com", "12345")

	if err != nil {
		t.Error(err)
		return
	}

	err = userDb.Create(user)

	if err != nil {
		t.Error(err)
		return
	}

	var userFound entity.User

	userDb.DB.Where("id = ?", user.ID.String()).First(&userFound)

	assert.Nil(t, err)
	assert.NotNil(t, userFound)
	assert.Equal(t, userFound.Name, "Usuario")
	assert.Equal(t, userFound.Email, "email@email.com")
	assert.True(t, user.ValidatePassword("12345"))
}

func TestFindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
		return
	}

	db.AutoMigrate(&entity.User{})

	userDb := NewUser(db)

	user, err := entity.NewUser("Usuario", "email@email.com", "12345")

	if err != nil {
		t.Error(err)
		return
	}

	err = userDb.Create(user)

	if err != nil {
		t.Error(err)
		return
	}

	userFound, err := userDb.FindByEmail("email@email.com")

	assert.Nil(t, err)
	assert.NotNil(t, userFound)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, userFound.Name, "Usuario")
	assert.Equal(t, userFound.Email, "email@email.com")
	assert.True(t, user.ValidatePassword("12345"))
}
