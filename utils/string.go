package utils

import (
	"fmt"
	"regexp"
)

// IsDuplicate : check duplicate error
func IsDuplicate(err error) bool {
	if err == nil {
		return false
	}
	ok, _ := regexp.MatchString("Error 1062", err.Error())
	return ok
}

// DiscordImageCDN : discord attach params to url
func DiscordImageCDN(channelID string, fileID string, fileName string) string {
	return fmt.Sprintf(
		"https://cdn.discordapp.com/attachments/%s/%s/%s", channelID, fileID, fileName,
	)
}

// DiscordMediaServer : discord attach params to resize url
func DiscordMediaServer(channelID string, fileID string, fileName string, query string) string {
	return fmt.Sprintf(
		"https://media.discordapp.net/attachments/%s/%s/%s?%s", channelID, fileID, fileName, query,
	)
}
