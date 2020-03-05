const CONFIG = require('./../config.json')

/**
 * Send random number
 *
 * @param {Object} message
 */
function dice(message) {
    const max = message._.text * 1 || 6
    if (!Number.isInteger(max)) {
        message.channel.send(`What?`)
        return
    }

    const number = Math.floor(Math.random() * max) + 1
    message.channel.send(`${number} / ${max}`)
}

dice.comment = [
    'Random number',
    `**${CONFIG.discord.prefix}dice** ***100***
    Get a random number (max 100)
    default: 6`,
]
module.exports = dice
