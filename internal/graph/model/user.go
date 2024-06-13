package model

import (
	"github.com/celtcoste/go-graphql-api-template/internal/repository"
	"github.com/celtcoste/go-graphql-api-template/pkg/util/gqlutil"
)

// User represents the User type in GraphQL
type User struct {
	UUID          gqlutil.UUID `json:"uuid"`
	Pseudo        string       `json:"pseudo"`
	Description   string       `json:"description"`
	ImageUUID     *gqlutil.UUID
	EmailVerified bool   `json:"emailVerified"`
	CGUAccepted   bool   `json:"cguAccepted"`
	Provider      string `json:"provider"`
}

// NewUserFromDB create a new User from a repository.User struct
func NewUserFromDB(user repository.User) *User {
	var uuid *gqlutil.UUID
	if user.ImageUUID != nil {
		tmp := gqlutil.UUID(*user.ImageUUID)
		uuid = &tmp
	} else {
		uuid = nil
	}
	return &User{
		UUID:          gqlutil.UUID(user.UUID),
		Pseudo:        user.Pseudo,
		Description:   user.Description,
		ImageUUID:     uuid,
		EmailVerified: user.EmailVerified,
		CGUAccepted:   user.CGUAccepted,
		Provider:      user.Provider,
	}
}

// UserEdge represents the UserEdge type in GraphQL
type UserEdge struct {
	Cursor string `json:"cursor"`
	Node   *User  `json:"node"`
}

// NewUserEdgesFromModel creates new []*UserEdge from a []*User
func NewUserEdgesFromModel(users []*User) []*UserEdge {
	edges := make([]*UserEdge, len(users))
	for i, user := range users {
		edges[i] = &UserEdge{
			Cursor: gqlutil.EncodeUUIDCursor(user.UUID.String()),
			Node:   user,
		}
	}
	return edges
}
