package twitter

import (
	"../secret"
	"github.com/gomodule/oauth1/oauth"
)

const (
	RefreshTokenURL  = "https://api.twitter.com/oauth/request_token"
	AuthorizationURL = "https://api.twitter.com/oauth/authenticate"
	AccessTokenURL   = "https://api.twitter.com/oauth/access_token"
	AccountURL       = "https://api.twitter.com/1.1/account/verify_credentials.json"
)

func Client() *oauth.Client {
	newClient := &oauth.Client{
		TemporaryCredentialRequestURI: RefreshTokenURL,
		ResourceOwnerAuthorizationURI: AuthorizationURL,
		TokenRequestURI:               AccessTokenURL,
		Credentials: oauth.Credentials{
			Token:  secret.Credentials["twitter_key"],
			Secret: secret.Credentials["twitter_secret"],
		},
	}

	return newClient
}
