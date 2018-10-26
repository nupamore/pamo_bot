
const mysql = require('mysql2/promise')
const CONFIG = require('./../../config.json')

const pool = mysql.createPool(CONFIG.db)
const query = `
    SELECT ORIGIN_URL, OWNER, REG_DATE 
    FROM images
    WHERE GROUP_ID = 454681618943049728
        AND REG_DATE > ?
        AND REG_DATE < ?
    ORDER BY REG_DATE DESC
`

/**
 * Get image list
 * 
 * @param {Object} req 
 * @param {Object} res 
 */
module.exports = async function images(req, res) {
    const {startDate, endDate} = req.query
    const connection = await pool.getConnection(async conn => conn)
    try {
        const [rows] = await connection.query(query, [startDate, endDate])
        res.send(rows)
        connection.release()
    }
    catch (err) {
        res.sendStatus(400)
        connection.release()
    }
}