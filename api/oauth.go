/*
 *    Copyright 2020 opensourceai
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opensourceai/go-api-service/pkg/app"
	"github.com/opensourceai/go-api-service/pkg/e"
	"github.com/unknwon/com"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"net/http"
)

var oauthStateString = com.ToStr(com.RandomCreateBytes(12))

func OauthApi(router *gin.Engine) {
	githubR := router.Group("/oauth/github")
	{
		githubR.GET("/login", githubLogin)
		githubR.GET("/callback", githubCallback)
	}
	githubG := router.Group("/oauth/google")
	{
		githubG.Any("/login", googleLogin)
		githubG.GET("/callback", googleCallback)
	}

}

var githubOauthConfig = &oauth2.Config{
	ClientID:     "41af06dd237b762a91cc",
	ClientSecret: "06044e6f11daef1f9e39e88d450cf2c32a4197d3",
	RedirectURL:  "http://api.tuboshu.io:8000/oauth/github/callback",
	Scopes:       []string{"user"},
	Endpoint:     github.Endpoint,
}

func githubLogin(c *gin.Context) {
	url := githubOauthConfig.AuthCodeURL(oauthStateString)
	fmt.Println(url)
	c.Redirect(http.StatusTemporaryRedirect, url)
}
func githubCallback(c *gin.Context) {
	appG := app.Gin{C: c}
	if state, b := c.GetQuery("state"); b {
		if state != oauthStateString {
			fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
			c.Redirect(http.StatusTemporaryRedirect, "/")
			return
		}
	}
	if code, b := c.GetQuery("code"); b {
		token, err := githubOauthConfig.Exchange(context.Background(), code)
		if err != nil {
			fmt.Printf("Code exchange failed with '%s'\n", err)
			appG.Response(http.StatusBadRequest, e.ERROR, nil)
			return
		}

		client := githubOauthConfig.Client(context.Background(), token)
		if response, err := client.Get("https://api.github.com/user"); err == nil && response != nil {
			defer response.Body.Close()
			contents, _ := ioutil.ReadAll(response.Body)
			appG.Response(http.StatusOK, e.SUCCESS, contents)
			return
		}

	} else {
		appG.Response(http.StatusBadRequest, e.ERROR, "")
	}

}

var googleOauthConfig = &oauth2.Config{
	ClientID:     "882682681914-ub6u8vac6o6fdr798l0skhau3tfj9hrf.apps.googleusercontent.com",
	ClientSecret: "PqGHL_LfX-lnIm7gSfNL77we",
	RedirectURL:  "http://api.tuboshu.io:8000/oauth/google/callback",
	Scopes: []string{"https://www.googleapis.com/auth/userinfo.profile",
		"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint: google.Endpoint,
}

func googleLogin(c *gin.Context) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	fmt.Println(url)
	c.Redirect(http.StatusTemporaryRedirect, url)
}
func googleCallback(c *gin.Context) {
	appG := app.Gin{C: c}
	if state, b := c.GetQuery("state"); b {
		if state != oauthStateString {
			fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
			c.Redirect(http.StatusTemporaryRedirect, "/")
			return
		}
	}
	if code, b := c.GetQuery("code"); b {
		token, err := googleOauthConfig.Exchange(context.Background(), code)
		if err != nil {
			fmt.Printf("Code exchange failed with '%s'\n", err)
			c.Redirect(http.StatusTemporaryRedirect, "/")
			return
		}
		appG.Response(http.StatusOK, e.SUCCESS, token)
	} else {
		appG.Response(http.StatusBadRequest, e.ERROR, nil)
	}

}
