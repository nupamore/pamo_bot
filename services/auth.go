package services

import (
	"encoding/json"
	"errors"

	"github.com/gofiber/session/v2"
	"github.com/monaco-io/request"
	"github.com/nupamore/pamo_bot/configs"
	"golang.org/x/oauth2"
)

// AuthService : auth service
type AuthService struct {
	Config   *oauth2.Config
	Sessions *session.Session
}

// Auth : auth service instance
var Auth = AuthService{}

// Setup : auth init
func (s *AuthService) Setup() {
	s.Config = &oauth2.Config{
		ClientID:     configs.Env["OAUTH_KEY"],
		ClientSecret: configs.Env["OAUTH_SECRET"],
		RedirectURL:  configs.Env["OAUTH_CALLBACK"],
		Scopes:       []string{"identify", "guilds"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  configs.Env["OAUTH_ENDPOINT"] + "/authorize",
			TokenURL: configs.Env["OAUTH_ENDPOINT"] + "/token",
		},
	}
	s.Sessions = session.New()
}

// LoginURL : get login url
func (s *AuthService) LoginURL(state string) string {
	return s.Config.AuthCodeURL(state)
}

// Authenticate : auth
func (s *AuthService) Authenticate(code string) (*oauth2.Token, error) {
	token, err := s.Config.Exchange(oauth2.NoContext, code)
	return token, err
}

type errRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// DiscordUser : discord user
type DiscordUser struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
	Avatar        string `json:"avatar"`
	Locale        string `json:"locale"`
}

// Info : get user info
func (s *AuthService) Info(auth string) (*DiscordUser, error) {
	client := request.Client{
		URL:    configs.Env["OAUTH_API"] + "/users/@me",
		Method: "GET",
		Header: map[string]string{
			"Authorization": auth,
		},
	}
	resp, err := client.Do()
	if err != nil {
		return nil, err
	}

	var res errRes
	json.Unmarshal(resp.Data, &res)
	if res.Message != "" {
		return nil, errors.New("OAuth api error")
	}

	var user DiscordUser
	json.Unmarshal(resp.Data, &user)
	return &user, nil
}

// DiscordGuild : discord guild
type DiscordGuild struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	Owner       bool   `json:"owner"`
	Permissions string `json:"permissions_new"`
}

// Guilds : get users guilds
func (s *AuthService) Guilds(auth string) ([]*DiscordGuild, error) {
	client := request.Client{
		URL:    configs.Env["OAUTH_API"] + "/users/@me/guilds",
		Method: "GET",
		Header: map[string]string{
			"Authorization": auth,
		},
	}
	resp, err := client.Do()
	if err != nil {
		return nil, err
	}

	var res errRes
	json.Unmarshal(resp.Data, &res)
	if res.Message != "" {
		return nil, errors.New("OAuth api error")
	}

	var guilds []*DiscordGuild
	json.Unmarshal(resp.Data, &guilds)
	return guilds, nil
}
