package services

import (
	"github.com/gofiber/session/v2"
	"github.com/nupamore/pamo_bot/configs"
	"golang.org/x/oauth2"
)

var authConfig *oauth2.Config

// Sessions : sessions
var Sessions *session.Session

// AuthSetup : auth init
func AuthSetup() {
	authConfig = &oauth2.Config{
		ClientID:     configs.Env["OAUTH_KEY"],
		ClientSecret: configs.Env["OAUTH_SECRET"],
		RedirectURL:  configs.Env["OAUTH_CALLBACK"],
		Scopes:       []string{"identify", "guilds"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  configs.Env["OAUTH_ENDPOINT"] + "/authorize",
			TokenURL: configs.Env["OAUTH_ENDPOINT"] + "/token",
		},
	}

	Sessions = session.New()
}

// GetLoginURL : get login url
func GetLoginURL(state string) string {
	return authConfig.AuthCodeURL(state)
}

// Authenticate : auth
func Authenticate(code string) (*oauth2.Token, error) {
	token, err := authConfig.Exchange(oauth2.NoContext, code)
	return token, err
}
