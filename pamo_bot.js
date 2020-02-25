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
    crawl: require('./commands/botCrawl'),
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
                guild.addScrapChannel(message)
                message.channel.send(
                    `I'm watching the pictures coming up on this channel! ＾ｐ＾`,
                )
            },
            off() {
                guild.removeScrapChannel(message)
                message.channel.send(
                    `I'm gonna stop watching, but you have to erase it yourself. (^^;)`,
                )
            },
            past() {
                f.crawl(message)
            },
        }[args[0]] || (() => noCommand(message))
    work()
}

/**
 * Bot login
 */
client.on('ready', () => {
    if (client.guilds.size !== guild.guildsList.size) {
        client.guilds.forEach(g => {
            if (guild.guildsList.get(g.id)) return
            guild.addGuildInfo(g)
        })
    }
    client.user.setActivity(CONFIG.discord.status)
    console.log(
        `Bot has started, with ${client.users.size} users, in ${client.channels.size} channels of ${client.guilds.size} guilds.`,
    )
})

/**
 * Join Event
 */
client.on('guildCreate', g => {
    guild.addGuildInfo(g)
})
client.on('guildDelete', g => {
    guild.removeGuildInfo(g.id)
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

            ke: () => f.translate(message, 'nmt', 'ko', 'en'),
            ek: () => f.translate(message, 'nmt', 'en', 'ko'),
            kj: () => f.translate(message, 'nmt', 'ko', 'ja'),
            jk: () => f.translate(message, 'nmt', 'ja', 'ko'),
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

            kj2: () => f.translate(message, 'kakao', 'kr', 'jp'),
            jk2: () => f.translate(message, 'kakao', 'jp', 'kr'),
            ke2: () => f.translate(message, 'kakao', 'kr', 'en'),
            ek2: () => f.translate(message, 'kakao', 'en', 'kr'),
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
