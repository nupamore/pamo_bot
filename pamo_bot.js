
const Discord = require('discord.js')
const CONFIG = require('./config.json')
const client = new Discord.Client()


/**
 * Functions list
 * Create new functions and put in this array
 */
const f = {
    translate: require('./commands/translate'),
    dice: require('./commands/dice'),
    crawl: require('./commands/crawl'),
    image: require('./commands/image'),
    test: require('./commands/test')
}

/**
 * Default functions
 * @param {object} message
 */
function commandList(message) {
    const comments = Object.values(f).filter(_ => _.comment).map(_ => _.comment)
    message.channel.send(comments.reduce((p, n) => `${ p }\n${ n }`, ''))
}
async function noCommand(message) {
    const m = await message.channel.send(`ハワワ (ㆁᴗㆁ✿)?`)
    setTimeout(() => m.edit(`そんなことば知らない! ( ￣＾￣)`), 1500)
}

/**
 * Bot login
 */
client.on('ready', () => {
    console.log(`Bot has started, with ${client.users.size} users, in ${client.channels.size} channels of ${client.guilds.size} guilds.`)
    client.user.setActivity(CONFIG.discord.status)
})

/**
 * Request message
 */
client.on('message', message => {
    if (message.author.bot) return
    if (message.content.indexOf(CONFIG.discord.prefix) !== 0) return

    const args = message.content.slice(CONFIG.discord.prefix.length).trim().split(/\s+/g)
    const command = args.shift().toLowerCase()
    message._ = { text: args.join(' ') }

    // Connect functions to custom command
    const func = {
        command: () => commandList(message),
        crawl: () => f.crawl(message),
        dice: () => f.dice(message),
        image: () => f.image(message),
        test: () => f.test(message),
        kj: () => f.translate(message, 'smt', 'ko', 'ja'),
        jk: () => f.translate(message, 'smt', 'ja', 'ko'),
        je: () => f.translate(message, 'nmt', 'ja', 'en'),
        ej: () => f.translate(message, 'nmt', 'en', 'ja'),
        ke: () => f.translate(message, 'nmt', 'ko', 'en'),
        ek: () => f.translate(message, 'nmt', 'en', 'ko'),
        fe: () => f.translate(message, 'nmt', 'fr', 'en'),
        ef: () => f.translate(message, 'nmt', 'en', 'fr'),
        fk: () => f.translate(message, 'nmt', 'fr', 'ko'),
        kf: () => f.translate(message, 'nmt', 'ko', 'fr'),
        sk: () => f.translate(message, 'nmt', 'es', 'ko'),
        ks: () => f.translate(message, 'nmt', 'ko', 'es'),
    }[command] || (() => noCommand(message))
    func(command)
})

client.login(CONFIG.discord.token)