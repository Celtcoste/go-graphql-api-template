// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"

	"github.com/celtcoste/go-graphql-api-template/pkg/util/gqlutil"
)

type UserConnection struct {
	PageInfo *gqlutil.PageInfo `json:"pageInfo"`
	Edges    []*UserEdge       `json:"edges"`
}

type Languages string

const (
	LanguagesFr Languages = "FR"
	LanguagesEn Languages = "EN"
)

var AllLanguages = []Languages{
	LanguagesFr,
	LanguagesEn,
}

func (e Languages) IsValid() bool {
	switch e {
	case LanguagesFr, LanguagesEn:
		return true
	}
	return false
}

func (e Languages) String() string {
	return string(e)
}

func (e *Languages) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Languages(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Languages", str)
	}
	return nil
}

func (e Languages) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
