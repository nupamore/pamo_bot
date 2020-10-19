package discord

import (
	"fmt"

	"github.com/diamondburned/arikawa/bot"
	"github.com/diamondburned/arikawa/gateway"
	"github.com/nupamore/pamo_bot/services"
)

// Commands : has prefix
type Commands struct {
	Ctx *bot.Context
}

// NoCommandHandler : if has not prefix
func NoCommandHandler(m *gateway.MessageCreateEvent) {
	// scrap image
	hasImage := len(m.Attachments) > 0
	_, isScrapingChannel := services.ScrapingChannelIDs[m.ChannelID]

	if hasImage && isScrapingChannel {
		services.ScrapImage(m.Message)
	}

	// auto transalte
	_, isAutoTranslateChannel := services.AutoTranslateChannelIDs[m.ChannelID]

	if hasImage || !isAutoTranslateChannel || len(m.Content) > 100 {
		return
	}
	detect, err := services.LanguageDetect(m.Content)
	if err != nil {
		return
	}
	var translatedText *string
	switch detect.LanguageInfo[0].Code {
	case "kr":
		translatedText, err = services.TranslatePapago("ko", "ja", m.Content)
	case "jp":
		translatedText, err = services.TranslatePapago("ja", "ko", m.Content)
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
