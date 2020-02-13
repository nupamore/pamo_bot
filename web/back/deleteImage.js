const mysql = require('mysql2/promise')
const CONFIG = require('./../../config.json')

const pool = mysql.createPool(CONFIG.db)
const QUERY = `
    DELETE FROM images
    WHERE ORIGIN_URL = ?
`

/**
 * Delete a image
 *
 * @param {Object} req
 * @param {Object} res
 */
module.exports = async function deleteImage(req, res) {
    const { originUrl, serverId } = req.body
    const passport = req.session.passport
    if (!passport) {
        res.sendStatus(401)
    }
    const group = req.session.passport.user.guilds.find(
        guild => guild.id === serverId,
    )
    if (group.permissions !== 2147483647) {
        res.sendStatus(403)
    }
    const connection = await pool.getConnection(async conn => conn)
    try {
        const [rows] = await connection.query(QUERY, originUrl)
        res.send(rows)
        connection.release()
    } catch (err) {
        res.sendStatus(400)
        connection.release()
    }
}
