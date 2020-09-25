package discord

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/diamondburned/arikawa/bot"
	"github.com/diamondburned/arikawa/gateway"
)

// Dice : return random number
func (cmd *Commands) Dice(_ *gateway.MessageCreateEvent, arg bot.RawArguments) (string, error) {
	max, err := strconv.Atoi(string(arg))

	if err != nil {
		max = 6
	}

	num := rand.Intn(max) + 1
	return fmt.Sprintf("%d / %d", num, max), nil
}
