const mysql = require('mysql2/promise')
const CONFIG = require('./../config.json')
const { imgUrl } = require('./../module/filters')

const pool = mysql.createPool(CONFIG.db)
const query = `
    SELECT channel_id, file_id, file_name
    FROM discord_images
    WHERE guild_id = ?
    ORDER BY rand() limit 1;
`

/**
 * Send random image
 *
 * @param {Object} message
 */
async function image(message) {
    const connection = await pool.getConnection(async conn => conn)
    try {
        const [rows] = await connection.query(query, message.guild.id)
        if (!rows.length) {
            message.channel.send('Not supported this group')
        } else {
            const { channel_id, file_id, file_name } = rows[0]
            const url = imgUrl(channel_id, file_id, file_name)
            message.channel.send('', { files: [url] })
        }
        connection.release()
    } catch (err) {
        message.channel.send(`DB error`)
        connection.release()
    }
}

image.comment = `***${CONFIG.discord.prefix}image***` + ` - Show a random image`
module.exports = image
