const db = require('./db/driver')

const guildsList = new Map()
const scrapChannels = new Set()

/**
 * Start realtime scrap
 * @param {String} guildId
 * @param {String} channelId
 */
async function addScrapChannel(message) {
    scrapChannels.add(message.channel.id)
    if (guildsList.get(message.guild.id)) {
        try {
            const [rows] = await db('UPDATE_GUILD', [
                'WATCH',
                message.channel.id,
                message.author.id,
                message.guild.id,
            ])
            return true
        } catch (err) {}
    } else {
        // hotfix
        try {
            await db('INSERT_GUILD', [
                message.guild.id,
                message.guild.name,
                message.channel.id,
                'WATCH',
                message.guild.ownerID,
                message.guild.ownerID,
            ])
            console.log(`New guild: ${message.guild.name}`)
        } catch (err) {
            console.log(err)
        }
    }
}

/**
 * Stop realtime scrap
 * @param {String} guildId
 * @param {String} channelId
 */
async function removeScrapChannel(message) {
    scrapChannels.delete(message.channel.id)
    try {
        const [rows] = await db('UPDATE_GUILD', [
            'STOP',
            null,
            message.author.id,
            message.guild.id,
        ])
        return true
    } catch (err) {
        console.log(err)
    }
}

/**
 * Add guild info
 * @param {Object} guilds
 */
async function addGuildInfo(g) {
    try {
        await db('INSERT_GUILD', [
            g.id,
            g.name,
            null,
            'CREATED',
            g.ownerID,
            g.ownerID,
        ])
        console.log(`New guild: ${g.name}`)
    } catch (err) {
        console.log(err)
    }
}

/**
 * Remove guild info
 * @param {String} guildId
 */
async function removeGuildInfo(guildId) {
    try {
        await db('DELETE_GUILD', guildId)
    } catch (err) {
        console.log(err)
    }
}

module.exports = {
    guildsList,
    scrapChannels,
    addScrapChannel,
    removeScrapChannel,
    addGuildInfo,
    removeGuildInfo,
    /**
     * Initial function
     */
    async init() {
        // Get guilds list
        try {
            const [rows] = await db('GET_GUILDS_LIST')
            rows.forEach(row => {
                guildsList.set(row.guild_id, row)
                // Realtime scrap channels
                if (row.status === 'WATCH') {
                    scrapChannels.add(row.scrap_channel_id)
                }
            })
            return true
        } catch (err) {
            console.log(err)
        }
    },
}
