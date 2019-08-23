
const express = require('express')
const CONFIG = require('./config.json')

const images = require('./web/back/images')
const uploaders = require('./web/back/uploaders')
const randomImage = require('./commands/image')


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
app.get('/uploaders', uploaders)
app.get('/randomImage', (req, res) => {
    randomImage({ channel: {
        guild: { id: 507169726473043968 },
        send(str, data) {
            if (data) res.redirect(data.files[0])
            else res.send(str)
        }
    }})
})


app.listen(CONFIG.web.port, () => {
    console.log(`Server start ${ CONFIG.web.port }`)
})