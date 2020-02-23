const mysql = require('mysql2/promise')
const CONFIG = require('./../../config.json')

const pool = mysql.createPool(CONFIG.db)
const QUERY = `
    DELETE FROM discord_images
    WHERE file_id = ?
`

/**
 * Delete a image
 *
 * @param {Object} req
 * @param {Object} res
 */
module.exports = async function deleteImage(req, res) {
    const { fileId, userId, serverId } = req.body
    const passport = req.session.passport
    if (!passport) {
        res.sendStatus(401)
    }
    const server = passport.user.guilds.find(guild => guild.id === serverId)
    const isMaster = server.permissions === 2147483647
    const isMine = userId === passport.user.id

    if (!isMaster && !isMine) {
        res.sendStatus(403)
    }
    const connection = await pool.getConnection(async conn => conn)
    try {
        const [rows] = await connection.query(QUERY, fileId)
        res.send(rows)
        connection.release()
    } catch (err) {
        res.sendStatus(400)
        connection.release()
    }
}
