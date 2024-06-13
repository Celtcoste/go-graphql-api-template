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
	var uuid *gqlutil.UUID
	if user.ImageUUID != nil {
		tmp := gqlutil.UUID(*user.ImageUUID)
		uuid = &tmp
	} else {
		uuid = nil
	}
	return &PrivateUser{
		UserInfo: &User{
			UUID:        gqlutil.UUID(user.UUID),
			Pseudo:      user.Pseudo,
			Description: user.Description,
			ImageUUID:   uuid,
		},
	}
}
