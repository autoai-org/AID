// Copyright 2019 The CVPM Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package authentication

/*
import (
	"fmt"
	"github.com/fatih/color"
	"github.com/levigross/grequests"
)
*/

// User defines the basic structure of User
type User struct {
	Username string `json:"username"`
	Password string
}

func (u *User) login() User {
	var respUser User
	loginURL := apiURL + "user/login"
	loginRequestHeader := &grequests.RequestOptions{
		JSON:   map[string]string{"username": u.Username, "password": u.Password},
		IsAjax: true,
	}
	resp, err := grequests.Post(loginURL, loginRequestHeader)
	if err != nil {
		log.Fatal(err)
	}
	if resp.Ok != true {
		fmt.Println("\nLogin Failed!")
	} else {
		_ = json.Unmarshal(resp.Bytes(), &respUser)
		setCache("session-token", respUser.SessionToken)
		color.Green("\nLogin Successfully")
	}
	return respUser
}
