package services

import (
	"os"

	"github.com/gofiber/session/v2"
	"golang.org/x/oauth2"
)

var authConfig *oauth2.Config

// Sessions : sessions
var Sessions *session.Session

// AuthSetup : auth init
func AuthSetup() {
	authConfig = &oauth2.Config{
		ClientID:     os.Getenv("OAUTH_KEY"),
		ClientSecret: os.Getenv("OAUTH_SECRET"),
		RedirectURL:  os.Getenv("OAUTH_CALLBACK"),
		Scopes:       []string{"identify", "guilds"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  os.Getenv("OAUTH_ENDPOINT") + "/authorize",
			TokenURL: os.Getenv("OAUTH_ENDPOINT") + "/token",
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
