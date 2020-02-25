const request = require('request')
const CONFIG = require('./../config.json')
const insertImage = require('../module/insertImage')

/**
 * Request logs
 *
 * @param {String} guild    Group ID
 * @param {String} channel  Channel ID
 * @param {Number} page
 */
function parse(guild, channel, page) {
    return new Promise((resolve, reject) => {
        const options = {
            url: `https://discordapp.com/api/v6/guilds/${guild}/messages/search?&channel_id=${channel}&offset=${page *
                25}&has=image&has=video&include_nsfw=true`,
            headers: {
                'user-agent':
                    'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) discord/0.0.300 Chrome/56.0.2924.87 Discord/1.6.15 Safari/537.36',
                authorization: CONFIG.discord.authorization,
            },
        }
        request.get(options, (error, response, body) => {
            if (!error && response.statusCode == 200) {
                const { total_results, messages } = JSON.parse(body)
                const images = messages
                    .reduce((p, n) => p.concat(n), [])
                    .filter(_ => _.attachments.length)
                    .map(_ => ({
                        userId: _.author.id,
                        userName: _.author.username,
                        userAvatar: _.author.avatar,
                        fileId: _.attachments[0].id,
                        filename: _.attachments[0].filename,
                        fileWidth: _.attachments[0].width,
                        fileHeight: _.attachments[0].height,
                        timestamp: new Date(_.timestamp),
                    }))
                resolve(images)
            } else {
                reject(body)
            }
        })
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

    const page = 30
    const promises = []
    for (let i = 0; i < page; i++) {
        promises.push(parse(message.channel.guild.id, message.channel.id, i))
    }

    Promise.all(promises)
        .then(outputs => {
            const images = outputs.reduce((p, n) => p.concat(n), [])
            return insertImage(
                images,
                message.channel.id,
                message.channel.guild.id,
            )
        })
        .then(success => {
            const count = success.reduce((p, n) => p + n)
            try {
                message.channel.send(`Crawled new images: ${count}`)
            } catch (err) {
                console.log(`Crawled new images: ${count}`)
            }
        })
        .catch(err => console.error(err))
}

/**
 * manual crawling
 */
// ;(() => {
//     crawl({
//         author: { id: 0 },
//         channel: {
//             id: '',
//             guild: { id: '' },
//         },
//     })
// })()

crawl.comment =
    `***${CONFIG.discord.prefix}crawl past***` +
    ` - Crawling past images of this channel\n` +
    `***${CONFIG.discord.prefix}crawl on***` +
    ` - The bot start real-time crawling`
module.exports = crawl
