package discord

import (
	"strings"

	"github.com/diamondburned/arikawa/bot"
	"github.com/diamondburned/arikawa/gateway"
)

// T : translate
func (cmd *Commands) T(_ *gateway.MessageCreateEvent, arg bot.RawArguments) (string, error) {
	if arg == "" {
		return "No target language", nil
	}

	target := strings.Fields(string(arg))[0]
	if string(arg) == target {
		return "No senetences", nil
	}

	text := strings.TrimLeft(string(arg), target)
	text = strings.TrimSpace(text)
	if len(text) > 100 {
		return "Too many senetences", nil
	}

	result, err := Service.TranslateAWS("auto", target, text)
	if err != nil {
		return "Invalid target language", nil
	}

	return *result, nil
}
