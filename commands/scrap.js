const insertImage = require('../module/insertImage')

/**
 * Scrap images realtime
 *
 * @param {Object} message
 */
async function scrap(message) {
    const images = message.attachments.map(_ => ({
        userId: message.author.id,
        userName: message.author.username,
        userAvatar: message.author.avatar,
        fileId: _.id,
        filename: _.filename,
        fileWidth: _.width,
        fileHeight: _.height,
        timestamp: new Date(message.createdTimestamp),
    }))
    insertImage(
        images,
        message.channel.id,
        message.channel.guild.id,
    ).catch(err => console.log(err))
}

module.exports = scrap
