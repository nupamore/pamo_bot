const CONFIG = require('./../config.json')
const db = require('./../module/db/driver')
const { imgUrl } = require('./../module/filters')

/**
 * Send random image
 *
 * @param {Object} message
 */
async function image(message) {
    try {
        const [rows] = await db('GET_RANDOM_IMAGE', message.guild.id)
        if (!rows.length) {
            message.channel.send('Not supported this group')
        } else {
            const { channel_id, file_id, file_name } = rows[0]
            const url = imgUrl(channel_id, file_id, file_name)
            message.channel.send('', { files: [url] })
        }
    } catch (err) {
        message.channel.send(`DB error`)
    }
}

image.comment = `***${CONFIG.discord.prefix}image***` + ` - Show a random image`
module.exports = image
