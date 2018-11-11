package main

import (
	"encoding/json"
	"fmt"
	"github.com/alexrudd/cognito-srp"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/endpoints"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	cip "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/fatih/color"
	"github.com/levigross/grequests"
	"log"
	"time"
)

type User struct {
	Username     string `json:"username"`
	Password     string
	SessionToken string `json:"sessionToken"`
}

func (u *User) login() User {
	fmt.Println(u.Username)
	csrp, _ := cognitosrp.NewCognitoSRP(u.Username, u.Password, "us-east-1_IYJ3FvCKZ", "1jinmsd412vcs8pkhqg5u0gjd2", nil)
	cfg, _ := external.LoadDefaultAWSConfig()
	cfg.Region = endpoints.UsEast1RegionID
	cfg.Credentials = aws.AnonymousCredentials
	svc := cip.New(cfg)
	req := svc.InitiateAuthRequest(&cip.InitiateAuthInput{
		AuthFlow:       cip.AuthFlowTypeUserSrpAuth,
		ClientId:       aws.String(csrp.GetClientId()),
		AuthParameters: csrp.GetAuthParams(),
	})
	resp, _ := req.Send()
	fmt.Println("123456")
	fmt.Println(resp.ChallengeName)
	if resp.ChallengeName == cip.ChallengeNameTypePasswordVerifier {
		challengeInput, _ := csrp.PasswordVerifierChallenge(resp.ChallengeParameters, time.Now())
		chal := svc.RespondToAuthChallengeRequest(challengeInput)
		resp, _ := chal.Send()
		fmt.Println(resp.AuthenticationResult)
	} else {

	}
	return *u
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
