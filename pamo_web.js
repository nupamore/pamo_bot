
const express = require('express')
const CONFIG = require('./config.json')

const images = require('./web/back/images')


/**
 * Server setup
 */
const app = express()
app.use('/', express.static(__dirname + '/web/front/dist'))

/**
 * Routers
 */
app.get('/', (req, res) => {
    res.redirect('/gallery.html')
})
app.get('/images', images)


app.listen(CONFIG.web.port, () => {
    console.log(`Server start ${ CONFIG.web.port }`)
})