package main

import (
	"net/http"

	"./twitter"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LoginByTwitter(c *gin.Context) {
	// ref: https://christina04.hatenablog.com/entry/2016/07/11/193000

	client := twitter.Client()
	tempCredentials, err := client.RequestTemporaryCredentials(
		nil,
		"http://localhost:8080/auth/twitter/callback",
		nil,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		panic(err)
		return
	}

	// temp credentialsを保存
	session := sessions.Default(c)
	session.Set("request_token", tempCredentials.Token)
	session.Set("request_token_secret", tempCredentials.Secret)
	session.Save()

	// 認証URLを発行、リダイレクト
	authURL := client.AuthorizationURL(tempCredentials, nil)
	c.Redirect(http.StatusMovedPermanently, authURL)
	return
}
