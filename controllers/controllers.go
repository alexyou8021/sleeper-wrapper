package controllers

import (
	"github.com/alexyou8021/sleeper-wrapper.git/clients/sleeper"
	"github.com/alexyou8021/sleeper-wrapper.git/entities"
)

func Controller(username string) entities.User {
	result := sleeper.GetUserByUsername(username)
	return result
}
