package database

import (
	"testing"

	"github.com/Scrowszinho/api-go/internal/entity"
	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("Gustavo", "gustavo@gmail.com", "123456")
	userDB := NewUser(db)

	err = userDB.Create(user)
	assert.Nil(t, err)

	var userFounder entity.User
	err = db.First(&userFounder, "id = ?", user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, userFounder.ID, user.ID)
	assert.Equal(t, userFounder.Email, user.Email)
	assert.Equal(t, userFounder.Name, user.Name)
	assert.NotNil(t, userFounder.Password)
}
