
const mysql = require('mysql2/promise')
const CONFIG = require('./../../config.json')

const pool = mysql.createPool(CONFIG.db)
const QUERY = {
    GET_IMAGES_COUNT: `
        SELECT COUNT(*) total
        FROM images
        WHERE GROUP_ID = 507169726473043968
            AND OWNER LIKE ?
    `,
    GET_IMAGES_INFO: `
        SELECT ORIGIN_URL, OWNER, REG_DATE 
        FROM images
        WHERE GROUP_ID = 507169726473043968
            AND OWNER LIKE ?
        ORDER BY REG_DATE DESC
        LIMIT 12 OFFSET ?
    `,
}

/**
 * Get image list
 * 
 * @param {Object} req 
 * @param {Object} res 
 */
module.exports = async function images(req, res) {
    const {/*startDate, endDate, */owner, page} = req.query
    const connection = await pool.getConnection(async conn => conn)
    try {
        let total = undefined
        if (page == 1) {
            const [result] =  await connection.query(QUERY.GET_IMAGES_COUNT, [/*startDate, endDate, */(owner == 'All') ? '%' : owner])
            total = result[0].total
        }
        const [images] = await connection.query(QUERY.GET_IMAGES_INFO, [/*startDate, endDate, */(owner == 'All') ? '%' : owner, (page - 1) * 12])
        res.send({ images, total })
        connection.release()
    }
    catch (err) {
        console.log(err)
        res.sendStatus(400)
        connection.release()
    }
}