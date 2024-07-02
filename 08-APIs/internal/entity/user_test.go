package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John Doe", "j@j.com", "123456")
	assert.Nil(t, err)     // valida se o error é nil
	assert.NotNil(t, user) // valida que o usuario nao seja em branco
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Email)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "j@j.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("John Doe", "j@j.com", "123456")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123456")) // tem que retornar true com a senha correta
	assert.False(t, user.ValidatePassword(""))      // tem que retornar false com a senha incorreta
	assert.NotEqual(t, "123456", user.Password)     // garantir que a senha esteja incriptada, nao sendo igual a 123456
}
