package model

import (
	chapter6 "chapter6/gen"
	"chapter6/pkg"
	"context"
)

type LoginInterface interface {
	ValidateUser(dbQuery *chapter6.Queries) bool
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse Response

//ValidateUser check whether username/password exist in database
func (login Login) ValidateUser(dbQuery *chapter6.Queries) bool {
	//query the data from database
	ctx := context.Background()
	u, _ := dbQuery.GetUserByName(ctx, login.Username)

	//username does not exist
	if u.UserName != login.Username {
		return false
	}

	return pkg.CheckPasswordHash(login.Password, u.PassWordHash)
}
