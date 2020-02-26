const db = require('./../../module/db/driver')

/**
 * Get image list
 *
 * @param {Object} req
 * @param {Object} res
 */
module.exports = async function images(req, res) {
    const { /*startDate, endDate, */ galleryId, owner, page } = req.query
    try {
        let total = undefined
        if (page == 1) {
            const [result] = await db('GET_IMAGES_COUNT', [
                /*startDate, endDate, */ galleryId,
                owner == 'All' ? '%' : owner,
            ])
            total = result[0].total
        }
        const [images] = await db('GET_IMAGES_INFO', [
            /*startDate, endDate, */ galleryId,
            owner == 'All' ? '%' : owner,
            (page - 1) * 12,
        ])
        res.send({ images, total })
    } catch (err) {
        console.log(err)
        res.sendStatus(400)
    }
}
