const mysql = require('mysql2/promise')
const CONFIG = require('./../config.json')

const pool = mysql.createPool(CONFIG.db)
const QUERY = {
    READ: `
        SELECT guild_id, scrap_channel_id, status
        FROM discord_guilds
    `,
    CREATE: `
        INSERT INTO discord_guilds (
            guild_id, guild_name, scrap_channel_id, status, 
            reg_user, reg_date, mod_user, mod_date
        )
        VALUES (?, ?, ?, ?, ?, NOW(), ?, NOW());
    `,
    UPDATE: `
        UPDATE discord_guilds SET 
            status = ?, 
            scrap_channel_id = ?,
            mod_user = ?,
            mod_date = NOW()
        WHERE guild_id = ?
    `,
}

const guildsList = new Map()
const scrapChannels = new Set()

/**
 * Start realtime scrap
 * @param {String} guildId
 * @param {String} channelId
 */
async function addScrapChannel(message) {
    scrapChannels.add(message.channel.id)
    const connection = await pool.getConnection(async conn => conn)
    try {
        const [rows] = await connection.query(QUERY.UPDATE, [
            'WATCH',
            message.channel.id,
            message.author.id,
            message.guild.id,
        ])
        connection.release()
        return true
    } catch (err) {
        connection.release()
    }
}

/**
 * Stop realtime scrap
 * @param {String} guildId
 * @param {String} channelId
 */
async function removeScrapChannel(message) {
    scrapChannels.delete(message.channel.id)
    const connection = await pool.getConnection(async conn => conn)
    try {
        const [rows] = await connection.query(QUERY.UPDATE, [
            'STOP',
            null,
            message.author.id,
            message.guild.id,
        ])
        connection.release()
        return true
    } catch (err) {
        connection.release()
    }
}

/**
 *
 * @param {MapIterator} guilds
 */
function addGuildInfo(guilds) {
    if (guilds.size === guildsList.size) return
    guilds.forEach(async guild => {
        if (guildsList.get(guild.id)) return

        const connection = await pool.getConnection(async conn => conn)
        try {
            await connection.query(QUERY.CREATE, [
                guild.id,
                guild.name,
                null,
                'CREATED',
                guild.ownerID,
                guild.ownerID,
            ])
            console.log(`New guild: ${guild.name}`)
            connection.release()
        } catch (err) {
            console.log(err)
            connection.release()
        }
    })
}

module.exports = {
    guildsList,
    scrapChannels,
    addScrapChannel,
    removeScrapChannel,
    addGuildInfo,
    /**
     * Initial function
     */
    async init() {
        // Get guilds list
        const connection = await pool.getConnection(async conn => conn)
        try {
            const [rows] = await connection.query(QUERY.READ)
            rows.forEach(row => {
                guildsList.set(row.guild_id, row)
                // Realtime scrap channels
                if (row.status === 'WATCH') {
                    scrapChannels.add(row.scrap_channel_id)
                }
            })
            connection.release()
            return true
        } catch (err) {
            connection.release()
        }
    },
}
