const API = require('./API')

module.exports = async function gateway(URI, params) {
    const [method, uri] = API[URI]
    const headers = { 'Content-Type': 'application/json' }

    let res = null
    if (method === 'GET') {
        const url = new URL(this.location.origin + uri)
        url.search = new URLSearchParams(params).toString()
        res = await fetch(url)
    } else {
        const body = params && JSON.stringify(params)
        res = await fetch(uri, { method, headers, body })
    }
    const data = await res.json()
    return data
}
