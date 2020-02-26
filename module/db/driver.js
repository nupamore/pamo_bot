const mysql = require('mysql2/promise')
const CONFIG = require('./../../config.json')
const QUERY = require('./QUERY')

const pool = mysql.createPool(CONFIG.db)

module.exports = async function driver(query, params) {
    return await pool.query(QUERY[query], params)
}
