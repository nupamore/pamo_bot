const mysql = require('mysql2/promise')
const CONFIG = require('./../../config.json')

const pool = mysql.createPool(CONFIG.db)
const QUERY = {
    READ: `
        SELECT owner_name, owner_id, owner_avatar, count(*) amount
        FROM discord_images
        WHERE guild_id = ?
        GROUP BY owner_id
        ORDER BY amount DESC
    `,
}

/**
 * Get image list
 *
 * @param {Object} req
 * @param {Object} res
 */
module.exports = async function uploaders(req, res) {
    const galleryId = req.query.galleryId
    const connection = await pool.getConnection(async conn => conn)
    try {
        const [rows] = await connection.query(QUERY.READ, galleryId)
        res.send(rows)
        connection.release()
    } catch (err) {
        res.sendStatus(400)
        connection.release()
    }
}
