const db = require('./../../module/db/driver')

/**
 * Get image list
 *
 * @param {Object} req
 * @param {Object} res
 */
module.exports = async function profile(req, res) {
    try {
        const [rows] = await db('EXIST_IMAGE_GUILDS_LIST')
        const list = rows.map(row => row.guild_id)
        const {
            id,
            username,
            discriminator,
            avatar,
            guilds,
        } = req.session.passport.user
        const hasBotGuilds = guilds.filter(guild => list.includes(guild.id))
        res.send({ id, username, discriminator, avatar, guilds: hasBotGuilds })
    } catch (err) {
        res.sendStatus(400)
    }
}
