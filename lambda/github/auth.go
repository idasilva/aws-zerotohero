package github

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/oauth2"
	"net/http"
)

type Auth struct {
	accessToken string
	ctx         context.Context
}

func (a *Auth) authenticate(accessKey string) (*http.Client,error) {
	_, err := a.validateAccessKey(accessKey)
	if err != nil {
		return nil, err
	}

	token := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: accessKey,
		},
	)

	return oauth2.NewClient(a.ctx, token), nil
}

func (a *Auth) validateAccessKey(accessKey string) (bool, error) {

	if accessKey == "" {
		return false, errors.New(fmt.Sprintf("The accessKey %s is not valid: ", accessKey))
	}

	return true, nil
}

//NewAuthentication
func NewAuthentication() *Auth {

	return &Auth{
		ctx: context.Background(),
	}
}
