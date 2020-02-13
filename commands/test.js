/**
 * Test function
 *
 * @param {Object} message
 */
function test(message) {
    // nupamo only
    if (message.author.id != 314029849562054666) {
        message.channel.send(`何だお前`)
        return
    }
    message.channel.send(
        `groupId: ${message.channel.guild.id}\nchannelId: ${message.channel.id}`,
    )
}

module.exports = test
