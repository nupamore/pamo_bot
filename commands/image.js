const CONFIG = require('./../config.json')
const db = require('./../module/db/driver')
const { imgUrl } = require('./../module/filters')

/**
 * Send random image
 *
 * @param {Object} message
 */
async function image(message) {
    const arg = message._.text
    const uploader = arg
        ? [...message.channel.guild.members.values()].find(
              member =>
                  member.user.username.toLowerCase() === arg.toLowerCase(),
          )
        : null
    if (arg && !uploader) {
        message.channel.send(`Couldn't find any image`)
        return
    }
    const uploaderId = uploader ? uploader.id : '%'

    try {
        const [rows] = await db('GET_RANDOM_IMAGE', [
            message.guild.id,
            uploaderId,
        ])
        if (!rows.length) {
            message.channel.send(`Couldn't find any image`)
        } else {
            const { channel_id, file_id, file_name } = rows[0]
            const url = imgUrl(channel_id, file_id, file_name)
            message.channel.send('', { files: [url] })
        }
    } catch (err) {
        message.channel.send(`DB error`)
    }
}

image.comment = `**${CONFIG.discord.prefix}image**` + ` - Show a random image\n`
image.comment +=
    `**${CONFIG.discord.prefix}image** ***username***` +
    ` - Show a random image uploaded by *username*`
module.exports = image
