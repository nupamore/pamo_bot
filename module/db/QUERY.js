module.exports = {
    // image
    GET_RANDOM_IMAGE: `
        SELECT channel_id, file_id, file_name, owner_name, reg_date
        FROM discord_images
        WHERE guild_id = ?
            AND owner_id LIKE ?
        ORDER BY rand() limit 1;
    `,
    DELETE_IMAGE: `
        DELETE FROM discord_images
        WHERE file_id = ?
    `,
    INSERT_IMAGES: `
        INSERT IGNORE INTO discord_images (
            file_id, file_name, owner_name, owner_id, owner_avatar, 
            guild_id, channel_id, width, height, reg_date, archive_date
        )
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, TIMESTAMP(?), NOW());
    `,
    GET_IMAGES_INFO: `
        SELECT channel_id, file_id, file_name, owner_id, owner_name, owner_avatar, reg_date
        FROM discord_images
        WHERE guild_id = ?
            AND owner_id LIKE ?
        ORDER BY reg_date DESC
        LIMIT 12 OFFSET ?
    `,
    GET_IMAGES_COUNT: `
        SELECT COUNT(*) total
        FROM discord_images
        WHERE guild_id = ?
            AND owner_id LIKE ?
    `,
    // uploader
    GET_UPLOADERS_INFO: `
        SELECT owner_name, owner_id, owner_avatar, count(*) amount
        FROM discord_images
        WHERE guild_id = ?
        GROUP BY owner_id
        ORDER BY amount DESC
    `,
    // guilds
    EXIST_IMAGE_GUILDS_LIST: `
        SELECT DISTINCT guild_id
        FROM discord_images
    `,
    GET_GUILDS_LIST: `
        SELECT guild_id, scrap_channel_id, status
        FROM discord_guilds
    `,
    INSERT_GUILD: `
        INSERT INTO discord_guilds (
            guild_id, guild_name, scrap_channel_id, status, 
            reg_user, reg_date, mod_user, mod_date
        )
        VALUES (?, ?, ?, ?, ?, NOW(), ?, NOW());
    `,
    UPDATE_GUILD: `
        UPDATE discord_guilds SET 
            status = ?, 
            scrap_channel_id = ?,
            mod_user = ?,
            mod_date = NOW()
        WHERE guild_id = ?
    `,
    DELETE_GUILD: `
        DELETE FROM discord_guilds
        WHERE guild_id = ?
    `,
}
