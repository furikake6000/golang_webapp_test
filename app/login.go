package main

import (
	"encoding/json"
	"net/http"

	"./twitter"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/oauth1/oauth"
)

func LoginByTwitter(c *gin.Context) {
	// ref: https://christina04.hatenablog.com/entry/2016/07/11/193000

	client := twitter.Client()

	// session確立、取得
	session := sessions.Default(c)

	// credentialsの取得
	tempCredentials, err := client.RequestTemporaryCredentials(
		nil,
		"http://localhost:8080/auth/twitter/callback",
		nil,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		panic(err)
	}

	// temp credentialsを保存
	session.Set("request_token", tempCredentials.Token)
	session.Set("request_token_secret", tempCredentials.Secret)
	session.Save()

	// 認証URLを発行、リダイレクト
	authURL := client.AuthorizationURL(tempCredentials, nil)
	c.Redirect(http.StatusFound, authURL)
	return
}

func TwitterCallback(c *gin.Context) {
	// 保存されていたtemp credentialsを取得
	session := sessions.Default(c)
	request_token := session.Get("request_token").(string)
	request_token_secret := session.Get("request_token_secret").(string)
	session.Save()
	// パラメータからoauth verifierを取得
	oauth_verifier := c.DefaultQuery("oauth_verifier", "")

	// AccessTokenを取得
	client := twitter.Client()
	accessToken, _, err := client.RequestToken(
		nil,
		&oauth.Credentials{
			Token:  request_token,
			Secret: request_token_secret,
		},
		oauth_verifier,
	)
	if err != nil {
		panic(err)
		return
	}

	// AccessTokenをcookieに保存
	session.Set("access_token", accessToken.Token)
	session.Set("access_token_secret", accessToken.Secret)

	response, err := client.Get(nil, accessToken, twitter.AccountURL, nil)
	if err != nil {
		panic(err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 500 {
		panic("500 error!")
	}
	if response.StatusCode >= 400 {
		panic("400 error!")
	}

	var user twitter.Account

	err = json.NewDecoder(response.Body).Decode(&user)
	if err != nil {
		panic(err)
	}

	c.HTML(200, "login_success.html", gin.H{
		"id":          user.Id,
		"screen_name": user.ScreenName,
		"image_url":   user.ImageURL,
	})
}
