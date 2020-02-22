const Discord = require('discord.js')
const CONFIG = require('./config.json')
const client = new Discord.Client()

/**
 * Modules
 */
const guild = require('./module/guild')

/**
 * Functions list
 * Create new functions and put in this array
 */
const f = {
    translate: require('./commands/translate'),
    image: require('./commands/image'),
    dice: require('./commands/dice'),
    crawl: require('./commands/crawl'),
    scrap: require('./commands/scrap'),
}

/**
 * Default commands
 * @param {object} message
 */
function commandList(message) {
    const comments = Object.values(f)
        .filter(_ => _.comment)
        .map(_ => _.comment)
    message.channel.send(
        'Photo Archive: https://vrc.nupa.moe\n' +
            comments.reduce((p, n) => `${p}\n${n}`, ''),
    )
}
async function noCommand(message) {
    const m = await message.channel.send(`ハワワ (ㆁᴗㆁ✿)?`)
    setTimeout(() => m.edit(`そんなことば知らない! ( ￣＾￣)`), 1500)
}

/**
 * Realtime scraping
 */
function crawl(message, args) {
    // master only
    if (message.author.id != message.guild.ownerID) {
        message.channel.send(
            `You don't have permission. Contact the server master`,
        )
        return
    }
    const work =
        {
            on() {
                guild.addScrapChannel(message.guild.id, message.channel.id)
                message.channel.send(
                    `I'm watching the pictures coming up on this channel! ＾ｐ＾`,
                )
            },
            off() {
                guild.removeScrapChannel(message.guild.id, message.channel.id)
                message.channel.send(
                    `I'm gonna stop watching, but you have to erase it yourself. (^^;)`,
                )
            },
            past() {
                message.channel.send(`Yay! Past investigation! ≖‿≖`)
                f.crawl(message)
            },
        }[args[0]] || (() => noCommand(message))
    work()
}

/**
 * Bot login
 */
client.on('ready', () => {
    guild.addGuildInfo(client.guilds)
    client.user.setActivity(CONFIG.discord.status)
    console.log(
        `Bot has started, with ${client.users.size} users, in ${client.channels.size} channels of ${client.guilds.size} guilds.`,
    )
})

/**
 * Request message
 */
client.on('message', message => {
    if (message.author.bot) return
    if (guild.scrapChannels.has(message.channel.id) && message.attachments.size)
        f.scrap(message)
    if (message.content === CONFIG.discord.prefix) return
    if (message.content.indexOf(CONFIG.discord.prefix) !== 0) return

    const args = message.content
        .slice(CONFIG.discord.prefix.length)
        .trim()
        .toLowerCase()
        .split(/\s+/g)
    const command = args.shift()
    message._ = { text: args.join(' ') }

    // Connect functions to custom command
    const func =
        {
            help: () => commandList(message),
            command: () => commandList(message),
            crawl: () => crawl(message, args),
            dice: () => f.dice(message),
            image: () => f.image(message),
            test: () => f.test(message),
            kj: () => f.translate(message, 'kakao', 'kr', 'jp'),
            jk: () => f.translate(message, 'kakao', 'jp', 'kr'),
            ke: () => f.translate(message, 'kakao', 'kr', 'en'),
            ek: () => f.translate(message, 'kakao', 'en', 'kr'),
            kc: () => f.translate(message, 'kakao', 'kr', 'cn'),
            ck: () => f.translate(message, 'kakao', 'cn', 'kr'),
            je: () => f.translate(message, 'nmt', 'ja', 'en'),
            ej: () => f.translate(message, 'nmt', 'en', 'ja'),
            fe: () => f.translate(message, 'nmt', 'fr', 'en'),
            ef: () => f.translate(message, 'nmt', 'en', 'fr'),
            fk: () => f.translate(message, 'nmt', 'fr', 'ko'),
            kf: () => f.translate(message, 'nmt', 'ko', 'fr'),
            sk: () => f.translate(message, 'nmt', 'es', 'ko'),
            ks: () => f.translate(message, 'nmt', 'ko', 'es'),
            ke2: () => f.translate(message, 'nmt', 'ko', 'en'),
            ek2: () => f.translate(message, 'nmt', 'en', 'ko'),
            kj2: () => f.translate(message, 'nmt', 'ko', 'ja'),
            jk2: () => f.translate(message, 'nmt', 'ja', 'ko'),
        }[command] || (() => noCommand(message))
    func()
})

/**
 * Init
 */
;(async () => {
    await guild.init()
    client.login(CONFIG.discord.token)
})()
