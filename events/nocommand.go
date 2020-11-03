package events

import (
	"fmt"

	"github.com/diamondburned/arikawa/gateway"
	"github.com/nupamore/pamo_bot/services"
)

// NoCommandHandler : if has not prefix
func NoCommandHandler(m *gateway.MessageCreateEvent) {
	scrapImage(m)
	autoTranslate(m)
}

func scrapImage(m *gateway.MessageCreateEvent) {
	hasImage := len(m.Attachments) > 0
	_, isScrapingChannel := services.Guild.ScrapingChannelIDs[m.ChannelID]

	if hasImage && isScrapingChannel {
		services.Image.Scrap(m.Message, m.GuildID)
	}
}

func autoTranslate(m *gateway.MessageCreateEvent) {
	_, isAutoTranslateChannel := services.Guild.AutoTranslateChannelIDs[m.ChannelID]
	if !isAutoTranslateChannel || len(m.Content) > 100 {
		return
	}

	detect, err := services.Translate.KakakoDetect(m.Content)
	if err != nil {
		return
	}
	var translatedText *string
	switch detect.LanguageInfo[0].Code {
	case "kr":
		translatedText, err = services.Translate.Papago("ko", "ja", m.Content)
	case "jp":
		translatedText, err = services.Translate.Papago("ja", "ko", m.Content)
	default:
		return
	}
	if err != nil {
		return
	}
	services.DiscordAPI.SendText(
		m.ChannelID,
		fmt.Sprintf("%s: %s", m.Author.Mention(), *translatedText),
	)
}
