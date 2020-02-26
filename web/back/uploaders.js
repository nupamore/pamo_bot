const db = require('./../../module/db/driver')

/**
 * Get image list
 *
 * @param {Object} req
 * @param {Object} res
 */
module.exports = async function uploaders(req, res) {
    const galleryId = req.query.galleryId
    try {
        const [rows] = await db('GET_UPLOADERS_INFO', galleryId)
        res.send(rows)
    } catch (err) {
        res.sendStatus(400)
    }
}
