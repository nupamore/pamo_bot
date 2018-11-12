
const request = require('request')
const mysql = require('mysql2/promise')
const CONFIG = require('./../config.json')

const pool = mysql.createPool(CONFIG.db)
const query = `
    INSERT INTO images (ORIGIN_URL, OWNER, GROUP_ID, WIDTH, HEIGHT, REG_DATE, ARCHIVE_DATE)
    VALUES (?, ?, ?, ?, ?, DATE(?), NOW());
`

/**
 * Insert images to Databsse
 * 
 * @param {Array} images 
 * @param {String} groupId 
 */
async function insertDB(images, groupId) {
    const promises = images.map(async image => {
        const connection = await pool.getConnection(async conn => conn)
        const { user, timestamp, url, width, height } = image
        try {
            const [rows] = await connection.query(query, [url, user, groupId, width, height, new Date(timestamp)])
            connection.release()
            return 1
        }
        catch (err) {
            // console.log(err)
            connection.release()
            return 0
        }
    })
    
    return Promise.all(promises)
}

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
            url: `https://discordapp.com/api/v6/guilds/${ guild }/messages/search?&channel_id=${ channel }&offset=${ page * 25 }&has=image&has=video&include_nsfw=true`,
            headers: {
                'user-agent': 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) discord/0.0.300 Chrome/56.0.2924.87 Discord/1.6.15 Safari/537.36', 
                'authorization': CONFIG.discord.authorization
            }
        }
        request.get(options, (error, response, body) => {
            if (!error && response.statusCode == 200) {
                const { total_results, messages } = JSON.parse(body)
                const images = messages.reduce((p, n) => p.concat(n), [])
                .filter(_ => _.attachments.length)
                .map(_ => ({
                    user: _.author.username,
                    timestamp: _.timestamp,
                    url: _.attachments[0].url,
                    width: _.attachments[0].width,
                    height: _.attachments[0].height
                }))
                resolve(images)
            }
            else {
                reject(response.statusCode)
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
    // nupamo only
    if (message.author.id != 314029849562054666) {
        message.channel.send(`何だお前`)
        return
    }

    const page = 100
    const promises = []
    for (let i=0; i<page; i++) {
        promises.push(parse(message.channel.guild.id, message.channel.id, i))
    }

    Promise.all(promises)
    .then(outputs => {
        const images = outputs.reduce((p, n) => p.concat(n), [])
        return insertDB(images, message.channel.guild.id)
    })
    .then(success => {
        const count = success.reduce((p, n) => p + n)
        message.channel.send(`Crawled new images: ${ count }`)
    })
    .catch(err => console.log(err))
}


module.exports = crawl