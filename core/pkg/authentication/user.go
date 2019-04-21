// Copyright 2019 The CVPM Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package authentication

import (
	"fmt"
	"time"

	"github.com/alexrudd/cognito-srp"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/endpoints"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	cip "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

type User struct {
	Username     string `json:"username"`
	Password     string
}

// Login returns the authenticated user
func (u *User) Login() User {
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
	if resp.ChallengeName == cip.ChallengeNameTypePasswordVerifier {
		challengeInput, _ := csrp.PasswordVerifierChallenge(resp.ChallengeParameters, time.Now())
		chal := svc.RespondToAuthChallengeRequest(challengeInput)
		resp, _ := chal.Send()
		fmt.Println(resp.AuthenticationResult)
	} else {
	}
	return *u
}
