
const mysql = require('mysql2/promise')
const CONFIG = require('./../../config.json')

const pool = mysql.createPool(CONFIG.db)
const QUERY = `
    SELECT DISTINCT group_id
    FROM images
`

/**
 * Get image list
 * 
 * @param {Object} req 
 * @param {Object} res 
 */
module.exports = async function profile(req, res) {
    const connection = await pool.getConnection(async conn => conn)
    try {
        const [rows] = await connection.query(QUERY)
        const list = rows.map(row => row.group_id)
        const { id, username, discriminator, avatar, guilds } = req.session.passport.user
        guilds.forEach(guild => {
            if (list.includes(guild.id)) guild.hasBot = true
        })
        res.send({ id, username, discriminator, avatar, guilds })
        connection.release()
    }
    catch (err) {
        res.sendStatus(400)
        connection.release()
    }
}