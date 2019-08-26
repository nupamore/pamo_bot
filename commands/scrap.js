
const mysql = require('mysql2/promise')
const CONFIG = require('./../config.json')

const pool = mysql.createPool(CONFIG.db)
const query = `
    INSERT INTO images (ORIGIN_URL, OWNER, GROUP_ID, WIDTH, HEIGHT, REG_DATE, ARCHIVE_DATE)
    VALUES (?, ?, ?, ?, ?, DATE_ADD(TIMESTAMP(?), INTERVAL 9 HOUR), NOW());
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
 * Scrap images realtime
 * 
 * @param {Object} message 
 */
async function scrap(message) {
    const images = message.attachments.map(_ => ({
        user: message.author.username,
        timestamp: message.createdTimestamp,
        url: _.url,
        width: _.width,
        height: _.height
    }))
    insertDB(images, message.channel.guild.id)
    .catch(err => console.log(err))
}


module.exports = scrap