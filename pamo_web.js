const express = require('express')
const session = require('express-session')
const passport = require('passport')
const DiscordStrategy = require('passport-discord').Strategy
const CONFIG = require('./config.json')

const images = require('./web/back/images')
const deleteImage = require('./web/back/deleteImage')
const uploaders = require('./web/back/uploaders')
const profile = require('./web/back/profile')
const randomImage = require('./commands/image')

/**
 * Server setup
 */
const app = express()
app.use('/', express.static(__dirname + '/web/dist'))
app.use(express.json())
app.use(express.urlencoded({ extended: false }))
app.use(
    session({
        secret: CONFIG.passport.clientID,
        resave: true,
        saveUninitialized: false,
    }),
)
app.use(passport.initialize())
app.use(passport.session())

/**
 * Routers
 */
app.use(function(req, res, next) {
    if (!req.session.passport && !req.path == '/auth/discord')
        res.redirect('/auth/discord')
    else next()
})
app.get('/', (req, res) => {
    res.redirect('/auth/discord')
})
app.get('/images', images)
app.delete('/image', deleteImage)
app.get('/uploaders', uploaders)
app.get('/profile', profile)
app.get('/randomImage/:serverId', (req, res) => {
    randomImage({
        guild: {
            id: req.params.serverId,
        },
        channel: {
            send(str, data) {
                if (data) res.redirect(data.files[0])
                else res.send(str)
            },
        },
        _: {},
    })
})

/**
 * Passport
 */
const User = {}
passport.serializeUser((user, done) => {
    var me = User[user.profile.accessToken]
    me = user.profile
    done(null, me)
})
passport.deserializeUser((user, done) => {
    done(null, user)
})
passport.use(
    new DiscordStrategy(CONFIG.passport, function(
        accessToken,
        refreshToken,
        profile,
        cb,
    ) {
        cb(null, { accessToken, refreshToken, profile })
    }),
)
app.get('/auth/discord', passport.authenticate('discord'))
app.get(
    '/auth/discord/callback',
    passport.authenticate('discord', {
        failureRedirect: '/fail',
    }),
    function(req, res) {
        res.redirect('/photo')
    },
)
app.get('/logout', (req, res) => {
    req.logout()
    res.redirect('/')
})

/**
 * Finally
 */
app.use((req, res, next) => {
    res.redirect('/')
})

app.listen(CONFIG.web.port, () => {
    console.log(`Server start ${CONFIG.web.port}`)
})
