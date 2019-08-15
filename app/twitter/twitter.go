package twitter

import (
	"../secret"
	"github.com/gomodule/oauth1/oauth"
)

const (
	refreshTokenURL  = "https://api.twitter.com/oauth/request_token"
	authorizationURL = "https://api.twitter.com/oauth/authenticate"
	accessTokenURL   = "https://api.twitter.com/oauth/access_token"
	accountURL       = "https://api.twitter.com/1.1/account/verify_credentials.json"
)

func Client() *oauth.Client {
	newClient := &oauth.Client{
		TemporaryCredentialRequestURI: refreshTokenURL,
		ResourceOwnerAuthorizationURI: authorizationURL,
		TokenRequestURI:               accessTokenURL,
		Credentials: oauth.Credentials{
			Token:  secret.Credentials["twitter_key"],
			Secret: secret.Credentials["twitter_secret"],
		},
	}

	return newClient
}
