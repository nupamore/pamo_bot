package services

import (
	"github.com/diamondburned/arikawa/discord"
)

// SendDM : send direct message
func SendDM(userID discord.UserID, contents string) error {
	channel, err := DiscordAPI.CreatePrivateChannel(userID)
	_, err = DiscordAPI.SendText(channel.ID, contents)

	return err
}
