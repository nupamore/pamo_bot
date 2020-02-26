const db = require('./db/driver')

/**
 * Insert image to Databsse
 *
 * @param {Array} images
 * @param {String} groupId
 */
module.exports = async function insertImage(images, channelId, groupId) {
    const promises = images.map(async image => {
        try {
            await db('INSERT_IMAGES', [
                image.fileId,
                image.filename,
                image.userName,
                image.userId,
                image.userAvatar,
                groupId,
                channelId,
                image.fileWidth,
                image.fileHeight,
                new Date(image.timestamp),
            ])
            return 1
        } catch (err) {
            // console.log(err)
            return 0
        }
    })
    return Promise.all(promises)
}
