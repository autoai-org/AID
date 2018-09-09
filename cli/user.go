// Copyright 2019 The CVPM Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/levigross/grequests"
)

type User struct {
	Username     string `json:"username"`
	Password     string
	SessionToken string `json:"sessionToken"`
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

func (u *User) become() User {
	var respUser User
	becomeURL := apiURL + "user/me"
	becomeRequestHeader := &grequests.RequestOptions{
		JSON:   map[string]string{"sessionToken": u.SessionToken},
		IsAjax: true,
	}
	resp, err := grequests.Post(becomeURL, becomeRequestHeader)
	if err != nil {
		log.Fatal(err)
	}
	if resp.Ok != true {
		fmt.Println("Login Failed")
	} else {
		_ = json.Unmarshal(resp.Bytes(), &respUser)
		setCache("session-token", respUser.SessionToken)
		fmt.Printf("Hello, ")
		color.Cyan(respUser.Username)
	}
	return respUser
}

func (u *User) logOut() {
	err := deleteKey("session-token")
	if err == nil {
		fmt.Println("Logout Successfully")
	} else {
		panic(err)
	}
}
