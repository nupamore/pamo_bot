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
}
