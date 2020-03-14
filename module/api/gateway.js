const API = require('./API')
const CONFIG = require('../../config.json')

module.exports = async function gateway(URI, { path, params } = {}) {
    let [method, uri] = API[URI]
    const headers = { 'Content-Type': 'application/json' }
    const defaultParams = { credentials: 'include' }

    if (path) {
        Object.entries(path).forEach(([key, val]) => {
            uri = uri.replace(`{${key}}`, val)
        })
    }
    uri = CONFIG.server.url + uri

    let res = null
    if (method === 'GET') {
        const url = new URL(uri)
        url.search = new URLSearchParams(params).toString()
        res = await fetch(url, defaultParams)
    } else {
        const body = params && JSON.stringify(params)
        res = await fetch(uri, { ...defaultParams, method, headers, body })
    }
    const data = await res.json()
    return data
}
