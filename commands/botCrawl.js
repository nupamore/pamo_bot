const dayjs = require('dayjs')
const CONFIG = require('./../config.json')
const insertImage = require('../module/insertImage')

/**
 * Request logs
 *
 * @param {Array} MessageArray
 */
function parse(messages) {
    return messages
        .filter(_ => _.attachments.size && !_.author.bot)
        .map(_ => {
            const image = _.attachments.values().next().value
            return {
                userId: _.author.id,
                userName: _.author.username,
                userAvatar: _.author.avatar,
                fileId: image.id,
                filename: image.filename,
                fileWidth: image.width,
                fileHeight: image.height,
                timestamp: dayjs(_.createdTimestamp),
            }
        })
}

/**
 * Crawling images and save json file
 *
 * @param {Object} message
 */
async function crawl(message) {
    // master only
    if (message.author.id != message.guild.ownerID) {
        message.channel.send(
            `You don't have permission. Contact the server master`,
        )
        return
    }

    const PAGE = 100
    let index = 0
    let lastId = null
    let lastDate = null
    let foundImageAmount = 0
    const m = await message.channel.send(`ハワワ (ㆁᴗㆁ✿)?`)

    for (; index < PAGE; index++) {
        const messages = await message.channel.fetchMessages({
            limit: 100,
            before: lastId,
        })
        const messageArray = Array.from(messages.values())
        const last = messageArray[messages.size - 1]
        const images = parse(messageArray)
        if (!last) break

        lastId = last.id
        lastDate = dayjs(last.createdTimestamp).format('YYYY-MM-DD')
        const promises = await insertImage(
            images,
            message.channel.id,
            message.guild.id,
        )
        foundImageAmount += promises.length
            ? promises.reduce((p, n) => p + n)
            : 0
        m.edit(
            `***Past Crawl ~ ${lastDate} (${index} / ${PAGE})***\nFound new images: ${foundImageAmount} ...`,
        )
    }
    m.edit(
        `***Past Crawl ~ ${lastDate}***\nFound new images: ${foundImageAmount}\nDone!`,
    )
}

crawl.comment = [
    'Archive images',
    `**${CONFIG.discord.prefix}crawl** ***past***
    Crawling past images of this channel
    **${CONFIG.discord.prefix}crawl** ***on***
    The bot start real-time crawling`,
]
module.exports = crawl
