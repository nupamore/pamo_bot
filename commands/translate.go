package commands

import (
	"fmt"
	"strings"

	"github.com/diamondburned/arikawa/v2/bot"
	"github.com/diamondburned/arikawa/v2/gateway"
	"github.com/nupamore/pamo_bot/services"
)

// T : translate
func (cmd *Commands) T(m *gateway.MessageCreateEvent, arg bot.RawArguments) (string, error) {
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

	result, err := services.Translate.AWS("auto", target, text)
	if err != nil {
		return "Invalid target language", nil
	}

	return fmt.Sprintf(
		"%s(%s): %s",
		m.Author.Mention(),
		*result.SourceLanguageCode,
		*result.TranslatedText,
	), nil
}
