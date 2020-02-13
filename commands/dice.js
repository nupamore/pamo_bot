/**
 * Send random number
 *
 * @param {Object} message
 */
function dice(message) {
    const max = message._.text * 1 || 100
    if (!Number.isInteger(max)) {
        message.channel.send(`What?`)
        return
    }

    const number = Math.floor(Math.random() * max) + 1
    message.channel.send(`${number} / ${max}`)
}

dice.comment = `!dice 6 - Get random number (default: 100)`
module.exports = dice
