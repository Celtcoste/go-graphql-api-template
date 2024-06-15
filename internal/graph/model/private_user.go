package model

import (
	"github.com/celtcoste/go-graphql-api-template/internal/repository"
	"github.com/celtcoste/go-graphql-api-template/pkg/util/gqlutil"
)

// PrivateUser represents the logged in user
type PrivateUser struct {
	UserInfo *User `json:"userInfo"`
}

// NewUserPrivateFromDB create a new User from a repository.User struct
func NewUserPrivateFromDB(user repository.User) *PrivateUser {
	return &PrivateUser{
		UserInfo: &User{
			UUID: gqlutil.UUID(user.UUID),
			Name: user.Name,
		},
	}
}
