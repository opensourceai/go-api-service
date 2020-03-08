package router

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
	"net/http"
)

var oauthStateString = com.ToStr(com.RandomCreateBytes(12))

func Oauth2(router *gin.Engine) {
	githubR := router.Group("/oauth2/github")
	{
		githubR.Any("/login", githubLogin)
		githubR.GET("/callback", githubCallback)
	}
	githubG := router.Group("/oauth2/google")
	{
		githubG.Any("/login", googleLogin)
		githubG.Any("/callback", googleCallback)
	}

}

var githubOauthConfig = &oauth2.Config{
	ClientID:     "41af06dd237b762a91cc",
	ClientSecret: "06044e6f11daef1f9e39e88d450cf2c32a4197d3",
	RedirectURL:  "http://localhost:8000/oauth2/github/callback",
	Scopes:       []string{"user"},
	Endpoint:     github.Endpoint,
}

// @Summary 获取Oauth github认证信息
// @Tags Auth
// @Produce  json
// @Param user body auth true "user"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth/login [post]
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
			c.Redirect(http.StatusTemporaryRedirect, "/")
			appG.Response(http.StatusOK, e.SUCCESS, token)
			return
		}

		appG.Response(http.StatusBadRequest, e.ERROR, nil)

		//if response, err := http.Get("https://api.github.com/user?access_token=" + token.AccessToken); err != nil && response != nil {
		//	defer response.Body.Close()
		//	contents, _ := ioutil.ReadAll(response.Body)
		//	appG.Response(http.StatusOK, e.SUCCESS, contents)
		//}

	}
	appG.Response(http.StatusBadRequest, e.ERROR, "")

}

var googleOauthConfig = &oauth2.Config{
	ClientID:     "882682681914-ub6u8vac6o6fdr798l0skhau3tfj9hrf.apps.googleusercontent.com",
	ClientSecret: "PqGHL_LfX-lnIm7gSfNL77we",
	RedirectURL:  "http://localhost:8000/oauth2/google/callback",
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
		fmt.Println("token", token)
		appG.Response(http.StatusOK, e.SUCCESS, token)
	}
	appG.Response(http.StatusBadRequest, e.ERROR, nil)

}
