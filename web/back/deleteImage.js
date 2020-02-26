const db = require('./../../module/db/driver')

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
    try {
        const [rows] = await db('DELETE_IMAGE', fileId)
        res.send(rows)
    } catch (err) {
        res.sendStatus(400)
    }
}
