const mysql = require('mysql2/promise')
const CONFIG = require('./../config.json')

const pool = mysql.createPool(CONFIG.db)
const QUERY = {
    CREATE: `
        INSERT INTO discord_images (
            file_id, file_name, owner_name, owner_id, owner_avatar, 
            guild_id, channel_id, width, height, reg_date, archive_date
        )
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, TIMESTAMP(?), NOW());
    `,
}

/**
 * Insert image to Databsse
 *
 * @param {Array} images
 * @param {String} groupId
 */
module.exports = async function insertImage(images, channelId, groupId) {
    const promises = images.map(async image => {
        const connection = await pool.getConnection(async conn => conn)
        try {
            const [rows] = await connection.query(QUERY.CREATE, [
                image.fileId,
                image.filename,
                image.userName,
                image.userId,
                image.userAvatar,
                groupId,
                channelId,
                image.fileWidth,
                image.fileHeight,
                new Date(image.timestamp),
            ])
            connection.release()
            return 1
        } catch (err) {
            // console.log(err)
            connection.release()
            return 0
        }
    })

    return Promise.all(promises)
}
