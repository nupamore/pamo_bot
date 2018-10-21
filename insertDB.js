
const fs = require('fs')
const mysql = require('mysql2/promise')
const CONFIG = require('./config.json')
const images = JSON.parse(fs.readFileSync('./images.json'))

const pool = mysql.createPool(CONFIG.db)
const query = `
    INSERT INTO images (ORIGIN_URL, OWNER, GROUP_ID, CHANNEL_ID, WIDTH, HEIGHT, REG_DATE, ARCHIVE_DATE)
    VALUES (?, ?, ?, ?, ?, ?, DATE(?), NOW());
`
const groupId = '454681618943049728'
const channelId = '454681618943049730'


images.forEach(async image => {
    const connection = await pool.getConnection(async conn => conn)
    const { user, timestamp, url, width, height } = image
    try {
        const [rows] = await connection.query(query, [url, user, groupId, channelId, width, height, new Date(timestamp)])
        console.log(rows.insertId)
        connection.release()
    }
    catch (err) {
        console.log(err)
        connection.release()
    }
})