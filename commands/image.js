
const mysql = require('mysql2/promise')
const CONFIG = require('./../config.json')

const pool = mysql.createPool(CONFIG.db)
const query = `
    SELECT ORIGIN_URL FROM images
    WHERE GROUP_ID = ?
    ORDER BY rand() limit 1;
`

/**
 * Send random image
 * 
 * @param {Object} message
 */
async function image(message) {
    try {
        const connection = await pool.getConnection(async conn => conn)
        const [rows] = await connection.query(query, message.channel.guild.id)
        if (!rows.length) {
            message.channel.send('Not supported this group')
        }
        else {
            message.channel.send('', { files: [rows[0].ORIGIN_URL] })
        }
        connection.release()
    }
    catch (err) {
        message.channel.send(`DB error`)
        connection.release()
    }
}


image.comment = `!image - Show a random image`
module.exports = image;