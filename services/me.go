package services

import (
	"encoding/json"
	"errors"

	"github.com/monaco-io/request"
	"github.com/nupamore/pamo_bot/configs"
)

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

// GetUserInfo : get user info
func GetUserInfo(auth string) (*DiscordUser, error) {
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

// GetUsersGuilds : get users guilds
func GetUsersGuilds(auth string) ([]*DiscordGuild, error) {
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
