package twitter

import (
	"my/secret"

	"github.com/gomodule/oauth1/oauth"
)

const (
	RefreshTokenURL  = "https://api.twitter.com/oauth/request_token"
	AuthorizationURL = "https://api.twitter.com/oauth/authenticate"
	AccessTokenURL   = "https://api.twitter.com/oauth/access_token"
	AccountURL       = "https://api.twitter.com/1.1/account/verify_credentials.json"
)

type Account struct {
	Id         string `json:"id_str"`
	ScreenName string `json:"screen_name"`
	ImageURL   string `json:"profile_image_url_https"`
}

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
