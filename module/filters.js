const base62 = require('base62')
const CONFIG = require('./../config.json')
base62.setCharacterSet(CONFIG.web.base62)

module.exports = {
    imgUrl(channel_id, file_id, file_name) {
        return `https://cdn.discordapp.com/attachments/${channel_id}/${file_id}/${file_name}`
    },
    avatarUrl(owner_id, owner_avatar) {
        return `https://cdn.discordapp.com/avatars/${owner_id}/${owner_avatar}.jpg`
    },
    origin2Thumb(url) {
        const media = url.replace('cdn.discordapp.com', 'media.discordapp.net')
        return /(mp4)$/.test(url)
            ? media + '?format=jpeg&width=400&height=225'
            : media + '?width=400&height=225'
    },
    thumb2Origin(url) {
        return url.replace('media.discordapp.net', 'cdn.discordapp.com')
    },
    idEncode(id, index) {
        id = id.padStart(19, 0)
        const a = base62.encode(id.slice(0, 10) * 1)
        const b = base62.encode(id.slice(10, id.length) * 1)
        return a + '-' + b + index
    },
    idDecode(str) {
        const arr = str.split('')
        const index = arr.splice(-1, 1)[0]
        const [a, b] = arr.join('').split('-')
        const na = (base62.decode(a) + '').padStart(10, '0')
        const nb = (base62.decode(b) + '').padStart(9, '0')
        const decoded = (na + nb).replace(/^0+/, '')
        return [decoded, index]
    },
}
